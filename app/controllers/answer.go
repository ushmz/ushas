package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"ushas/models"
	"ushas/views"

	"github.com/labstack/echo/v4"
)

// AnswerController : Struct for controll `Answer` resource.
type AnswerController struct{}

// NewAnswerController : Return pointer to `AnswerController`.
func NewAnswerController() *AnswerController {
	return new(AnswerController)
}

// Index : index
func (ac *AnswerController) Index(c echo.Context) error {
	if ac == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("AnswerController is nil"))
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"OK",
	))
}

// Get : Get single answer by ID.
func (ac *AnswerController) Get(c echo.Context) error {
	if ac == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("AnswerController is nil"))
	}

	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.NewAPIError(
			err,
			"Answer ID must be number",
			idstr,
		))
	}
	ans, err := models.GetAnswerByID(int(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	answer := views.NewAnswerView(ans)
	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		answer,
	))
}

// List : Lists all answers.
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
		return c.JSON(http.StatusInternalServerError, err)
	}

	answers := views.NewListAnswerView(ans)
	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		answers,
	))
}

// CreateAnswerRequest : Request parameters for createing new `Answer` resource.
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

// Create : Create new answer.
func (ac *AnswerController) Create(c echo.Context) error {
	if ac == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("AnswerController is nil"))
	}

	p := CreateAnswerRequest{}
	if err := c.Bind(&p); err != nil {
		// Failed to bind request body
		return echo.NewHTTPError(http.StatusBadRequest, models.NewAPIError(
			err,
			"Failed to bind request body.",
			p,
		))
	}
	if err := c.Validate(&p); err != nil {
		if e, ok := err.(*models.APIError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, e)
		}
		return echo.NewHTTPError(http.StatusBadRequest, models.NewAPIError(
			err,
			"Failed to validate.",
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
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		p,
	))
}

// UpdateAnswerRequest : Request parameters for update answer.
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

// Update : Update answer.
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
		return echo.NewHTTPError(http.StatusBadGateway, models.NewAPIError(err, "Failed to bind requested body", p))
	}

	if err := c.Validate(&p); err != nil {
		if e, ok := err.(*models.APIError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, e)
		}
		return echo.NewHTTPError(http.StatusBadRequest, models.NewAPIError(err, "Invalid request body.", p))
	}

	model := &models.Answer{
		UserID:      p.UserID,
		TaskID:      p.TaskID,
		ConditionID: p.ConditionID,
		Answer:      p.Answer,
		Reason:      p.Reason,
	}
	if err := models.UpdateAnswer(model); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		p,
	))

}

// Delete : Delete a single answer.
func (ac *AnswerController) Delete(c echo.Context) error {
	if ac == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, models.NewAPIError(
			errors.New("AnswerController is nil"),
			"Something wrong with server.",
			nil,
		))
	}

	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.NewAPIError(
			err,
			"Answer ID must be number",
			idstr,
		))
	}

	if err := models.DeleteAnswer(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		id,
	))
}
