package controllers

import (
	"net/http"
	"ushas/models"

	"github.com/labstack/echo/v4"
)

// TaskController : Struct for controll `Task` resource.
type TaskController struct{}

// NewTaskController : Return pointer to `TaskController`.
func NewTaskController() *TaskController {
	return new(TaskController)
}

// Index : index
func (tc *TaskController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"OK",
	))
}

// CreateTaskRequest : Struct for task create request body.
type CreateTaskRequest struct {
	// Query : Search query for this task.
	Query string `json:"query"`

	// Title : Title of this task.
	Title string `json:"title"`

	// Description : Description text of task.
	Description string `json:"description"`

	// SearchURL : Url used in this task.
	SearchURL string `json:"searchUrl"`
}

// Create : Create new task.
func (tc *TaskController) Create(c echo.Context) error {
	if tc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrNilReceiver)
	}

	p := CreateTaskRequest{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.NewAppError(err, http.StatusBadRequest, "Failed bind request body.", p),
		)
	}

	if err := c.Validate(p); err != nil {
		if e, ok := err.(*models.AppError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, e)
		}
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.NewAppError(err, http.StatusBadRequest, "Failed bind request body.", p),
		)
	}

	t := &models.Task{
		Query:       p.Query,
		Title:       p.Title,
		Description: p.Description,
		SearchURL:   p.SearchURL,
	}
	if err := models.CreateTask(t); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		t,
	))
}

// Allocate : Allocates tasks for new user.
func (tc *TaskController) Allocate(c echo.Context) error {
	if tc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrNilReceiver)
	}

	return echo.NewHTTPError(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

// GetByIDRequest : Request body.
type GetByIDRequest struct {
	ID int `json:"id" param:"id" validate:"required,numeric"`
}

// GetByID : Get a task by ID.
func (tc *TaskController) GetByID(c echo.Context) error {
	if tc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrNilReceiver)
	}

	p := GetByIDRequest{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.NewAppError(err, http.StatusBadRequest, "Failed bind request body.", p),
		)
	}

	if err := c.Validate(&p); err != nil {
		if e, ok := err.(*models.AppError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, e)
		}
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.NewAppError(err, http.StatusBadRequest, "Failed bind request body.", p),
		)
	}

	t, err := models.GetTaskByID(p.ID)
	if err != nil {
		if e, ok := err.(*models.AppError); ok {
			return echo.NewHTTPError(http.StatusNotFound, e)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, newResponse(http.StatusOK, http.StatusText(http.StatusOK), t))
}

// List : Lists all tasks.
func (tc *TaskController) List(c echo.Context) error {
	if tc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrNilReceiver)
	}

	tasks, err := models.ListTasks()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		tasks,
	))
}

// UpdateTaskRequest : Request parameters for update task.
type UpdateTaskRequest struct {
	// ID : The ID of task
	ID int `json:"id" param:"id" validate:"required,numeric"`

	// Query : Search query for this task.
	Query string `json:"query" validate:"required"`

	// Title : Title of this task.
	Title string `json:"title" validate:"required"`

	// Description : Description text of task.
	Description string `json:"description" validate:"required"`

	// SearchURL : Url used in this task.
	SearchURL string `json:"search_url" validate:"required"`
}

// Update : Updates task.
func (tc *TaskController) Update(c echo.Context) error {
	if tc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrNilReceiver)
	}

	p := UpdateTaskRequest{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.NewAppError(err, http.StatusBadRequest, "Failed bind request body.", p),
		)
	}

	if err := c.Validate(p); err != nil {
		if e, ok := err.(*models.AppError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, e)
		}
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.NewAppError(err, http.StatusBadRequest, "Failed bind request body.", p),
		)
	}

	t := &models.Task{
		ID:          p.ID,
		Query:       p.Query,
		Title:       p.Title,
		Description: p.Description,
		SearchURL:   p.SearchURL,
	}
	if err := models.UpdateTask(t); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		t,
	))
}

// DeleteTaskRequest : Parameters to delete task.
type DeleteTaskRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}

// Delete : Deletes a single task.
func (tc *TaskController) Delete(c echo.Context) error {
	if tc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrNilReceiver)
	}

	p := DeleteTaskRequest{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.NewAppError(err, http.StatusBadRequest, "Failed bind request body.", p),
		)
	}

	if err := c.Validate(&p); err != nil {
		if e, ok := err.(*models.AppError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, e)
		}
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.NewAppError(err, http.StatusBadRequest, "Failed bind request body.", p),
		)
	}

	if err := models.DeleteTask(p.ID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"",
	))
}
