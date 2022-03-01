package controllers

import (
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
	if ac == nil {
		return c.JSON(http.StatusInternalServerError, newResponse(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			"Controller is nil",
		))
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"OK",
	))
}

func (ac *AnswerController) Get(c echo.Context) error {
	if ac == nil {
		return c.JSON(http.StatusInternalServerError, newResponse(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			"Controller is nil",
		))
	}

	idstr := c.Param("id")
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
		if e, ok := err.(*models.APIError); ok {
			return c.JSON(e.Code, newResponse(e.Code, e.Message, ""))
		}

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

func (ac *AnswerController) List(c echo.Context) error {
	if ac == nil {
		return c.JSON(http.StatusInternalServerError, newResponse(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			"Controller is nil",
		))
	}

	ans, err := models.ListAnswers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newResponse(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			"Failed to get answers",
		))
	}

	answers := views.NewListAnswerView(ans)
	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		answers,
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
	if ac == nil {
		return c.JSON(http.StatusInternalServerError, newResponse(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			"Controller is nil",
		))
	}

	p := CreateAnswerRequest{}
	if err := c.Bind(&p); err != nil {
		return newErrResponse(c, http.StatusBadRequest, err, "Failed to bind request body")
	}
	if err := c.Validate(p); err != nil {
		return newErrResponse(c, http.StatusBadRequest, err, p)
	}

	model := &models.Answer{
		UserID:      p.UserID,
		TaskID:      p.TaskID,
		ConditionID: p.ConditionID,
		Answer:      p.Answer,
		Reason:      p.Reason,
	}

	if err := models.CreateAnswer(model); err != nil {
		return newErrResponse(c, http.StatusInternalServerError, err, p)
	}

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		p,
	))
}

type UpdateAnswerRequest struct {
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

func (ac *AnswerController) Update(c echo.Context) error {
	if ac == nil {
		return c.JSON(http.StatusInternalServerError, newResponse(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			"Controller is nil",
		))
	}

	p := UpdateAnswerRequest{}
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, newResponse(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			"Failed to parse request body",
		))
	}

	model := &models.Answer{
		UserID:      p.UserID,
		TaskID:      p.TaskID,
		ConditionID: p.ConditionID,
		Answer:      p.Answer,
		Reason:      p.Reason,
	}
	if err := models.UpdateAnswer(model); err != nil {
		if e, ok := err.(*models.APIError); ok {
			return c.JSON(e.Code, newResponse(e.Code, e.Message, e.Result))
		}
		return c.JSON(http.StatusInternalServerError, newResponse(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			"Failed to update",
		))
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		p,
	))

}

func (ac *AnswerController) Delete(c echo.Context) error {
	if ac == nil {
		return c.JSON(http.StatusInternalServerError, newResponse(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			"Controller is nil",
		))
	}

	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, newResponse(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			"Answer ID must be number",
		))
	}
	if err := models.DeleteAnswer(int(id)); err != nil {
		if e, ok := err.(*models.APIError); ok {
			return c.JSON(e.Code, newResponse(e.Code, e.Message, ""))
		}

		return c.JSON(http.StatusInternalServerError, newResponse(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			"Failed to delete answer",
		))
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		id,
	))
}
