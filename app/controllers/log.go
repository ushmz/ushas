package controllers

import (
	"net/http"
	"ushas/models"

	"github.com/labstack/echo/v4"
)

type LogController struct{}

func NewLogController() *LogController {
	return new(LogController)
}

func (lc *LogController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"OK",
	))
}

// SerpViewingLogRequest : Struct for task viewing time log request body
type SerpViewingLogRequest struct {
	// UserID : The ID of user (worker)
	UserID int `json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `json:"task"`

	// ConditionId : User's condition ID that means group and task category.
	ConditionId int `json:"condition"`
}

func (lc *LogController) CreateSerpViewingLog(c echo.Context) error {
	if lc == nil {
		return newErrResponse(c, http.StatusInternalServerError, nil, nil)
	}

	p := SerpViewingLogRequest{}
	if err := c.Bind(&p); err != nil {
		return newErrResponse(c, http.StatusBadRequest, err, "Cannot bind request body")
	}

	if err := c.Validate(&p); err != nil {
		return newErrResponse(c, http.StatusBadRequest, err, p)
	}

	log := &models.SerpViewingLog{
		UserID: p.UserID,
		TaskID: p.TaskID,
		ConditionID: p.ConditionId,
	}
	if err := models.CreateSerpViewingLog(log); err != nil {
		return newErrResponse(c, http.StatusInternalServerError, err, p)
	}

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		p,
	))
}

// PageViewingLogRequest : Struct for each search result page viewing time log request body
type PageViewingLogRequest struct {
	// UserID : The ID of user (worker)
	UserID int `json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `json:"task"`

	// ConditionID : User's condition ID that means group and task category.
	ConditionID int `json:"condition"`

	// PageID : Page ID that user view.
	PageID int `json:"page"`
}

// SearchPageEventLogRequest : Struct for page click log request body.
type SearchPageEventLogRequest struct {
	// ID : The ID of each log record.
	ID string `json:"id"`

	// Uid : The ID of user (worker)
	User int `json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `json:"task"`

	// ConditionID : User's condition ID that means group and task category.
	ConditionID int `json:"condition"`

	// Time : User's page viewing time.
	Time int `json:"time"`

	// Page : The ID of page that user clicked.
	Page int `json:"page"`

	// Rank : Search result rank that user clicked.
	Rank int `json:"rank"`

	// IsVisible : Risk is visible or not.
	IsVisible bool `json:"visible"`

	// Event : It is expected to be "click", "hover" or "paginate"
	Event string `json:"event"`
}

// SearchSessionRequest : Struct fot search session request body.
type SearchSessionRequest struct {
	// UserID : Assigned ID of user (worker)
	UserID int `json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `json:"task"`

	// ConditionID : User's condition ID that means group and task category.
	ConditionID int `json:"condition"`
}
