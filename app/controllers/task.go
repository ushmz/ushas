package controllers

import (
	"net/http"
	"strconv"
	"ushas/models"

	"github.com/labstack/echo/v4"
)

// TaskController : Struct for controll `Task` resource.
type TaskController struct{}

// NewTaskController : Return pointer to `LogController`.
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
		return newErrResponse(c, http.StatusInternalServerError, nil, nil)
	}

	p := CreateTaskRequest{}
	if err := c.Bind(&p); err != nil {
		return newErrResponse(c, http.StatusBadRequest, err, nil)
	}

	if err := c.Validate(p); err != nil {
		return newErrResponse(c, http.StatusBadRequest, err, p)
	}

	t := &models.Task{
		Query:       p.Query,
		Title:       p.Title,
		Description: p.Description,
		SearchURL:   p.SearchURL,
	}
	if err := models.CreateTask(t); err != nil {
		return newErrResponse(c, http.StatusInternalServerError, err, nil)
	}

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		t,
	))
}

// AllocateTask : Allocate tasks for new user.
func (tc *TaskController) AllocateTask(c echo.Context) error {
	if tc == nil {
		return newErrResponse(c, http.StatusInternalServerError, nil, nil)
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"",
	))
}

// GetTaskByID : Get a task by ID.
func (tc *TaskController) GetTaskByID(c echo.Context) error {
	if tc == nil {
		return newErrResponse(c, http.StatusInternalServerError, nil, nil)
	}

	p := c.Param("id")
	if len(p) > 0 {
		return newErrResponse(c, http.StatusBadRequest, nil, p)
	}

	id, err := strconv.Atoi(p)
	if err != nil {
		return newErrResponse(c, http.StatusBadRequest, err, id)
	}

	t, err := models.GetTaskByID(id)
	if err != nil {
		return newErrResponse(c, http.StatusInternalServerError, err, p)
	}

	return c.JSON(http.StatusOK, newResponse(http.StatusOK, http.StatusText(http.StatusOK), t))
}

// ListTask : Lists all tasks.
func (tc *TaskController) ListTask(c echo.Context) error {
	if tc == nil {
		return newErrResponse(c, http.StatusInternalServerError, nil, nil)
	}

	tasks, err := models.ListTasks()
	if err != nil {
		return newErrResponse(c, http.StatusInternalServerError, err, tasks)
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
	ID int `db:"id" json:"id"`

	// Query : Search query for this task.
	Query string `db:"query" json:"query"`

	// Title : Title of this task.
	Title string `db:"title" json:"title"`

	// Description : Description text of task.
	Description string `db:"description" json:"description"`

	// SearchURL : Url used in this task.
	SearchURL string `db:"search_url" json:"searchUrl"`
}

// UpdateTask : Update task.
func (tc *TaskController) UpdateTask(c echo.Context) error {
	if tc == nil {
		return newErrResponse(c, http.StatusInternalServerError, nil, nil)
	}

	p := UpdateTaskRequest{}
	if err := c.Bind(&p); err != nil {
		return newErrResponse(c, http.StatusBadRequest, err, p)
	}

	if err := c.Validate(p); err != nil {
		return newErrResponse(c, http.StatusBadRequest, err, p)
	}

	t := &models.Task{
		ID:          p.ID,
		Query:       p.Query,
		Title:       p.Title,
		Description: p.Description,
		SearchURL:   p.SearchURL,
	}
	if err := models.UpdateTask(t); err != nil {
		return newErrResponse(c, http.StatusInternalServerError, err, t)
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		t,
	))
}

// DeleteTask : Delete a single task.
func (tc *TaskController) DeleteTask(c echo.Context) error {
	if tc == nil {
		return newErrResponse(c, http.StatusInternalServerError, nil, nil)
	}

	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		return newErrResponse(c, http.StatusBadRequest, err, "Task ID must be number")
	}

	if err := models.DeleteTask(id); err != nil {
		return newErrResponse(c, http.StatusInternalServerError, err, id)
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"",
	))
}
