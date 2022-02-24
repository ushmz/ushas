package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"ushas/models"
	"ushas/views"

	"github.com/labstack/echo/v4"
)

type AnswerController struct{}

func NewAnswerController() *AnswerController {
	return new(AnswerController)
}

func (ac *AnswerController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"OK",
	))
}

func (ac *AnswerController) Get(c echo.Context) error {
	idstr := c.QueryParam("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, newResponse(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			"Answer ID must be number",
		))
	}
	ans, err := models.GetAnswerByID(int(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newResponse(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			"Failed to get answer",
		))
	}

	answer := views.NewAnswerView(ans)
	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		answer,
	))
}

type CreateAnswerRequest struct {
	// UserID : Means external ID.
	UserID int `json:"user" validate:"required,numeric"`

	// TaskID : The identity of task.
	TaskID int `json:"task" validate:"required,numeric"`

	// ConditionID : This point out which kind of task did user take.
	ConditionID int `json:"condition" validate:"required,numeric"`

	// Answer : The Url of evidence of users' decision.
	Answer string `json:"answer" validate:"required"`

	// Reason : The reason of users' decision.
	Reason string `json:"reason" validate:"required"`
}

func (ac *AnswerController) Create(c echo.Context) error {
	p := CreateAnswerRequest{}
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, newResponse(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			p,
		))
	}
	if err := c.Validate(p); err != nil {
		if e, ok := err.(*models.APIError); ok {
			return c.JSON(http.StatusBadRequest, newResponse(
				http.StatusBadRequest,
				err.Error(),
				e.Result,
			))
		}
		return c.JSON(http.StatusBadRequest, newResponse(
			http.StatusBadRequest,
			err.Error(),
			p,
		))
	}

	model := &models.Answer{
		UserID:      p.UserID,
		TaskID:      p.TaskID,
		ConditionID: p.ConditionID,
		Answer:      p.Answer,
		Reason:      p.Reason,
	}

	if err := models.CreateAnswer(model); err != nil {
		if errors.Is(err, &models.APIError{}) {
			return c.JSON(http.StatusInternalServerError, newResponse(
				http.StatusInternalServerError,
				err.Error(),
				p,
			))
		}
		return c.JSON(http.StatusInternalServerError, newResponse(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			p,
		))
	}

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		p,
	))
}
