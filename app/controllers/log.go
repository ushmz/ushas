package controllers

import (
	"errors"
	"net/http"
	"ushas/models"

	"github.com/labstack/echo/v4"
)

// LogController : Struct for controll `Log` resource.
type LogController struct{}

// NewLogController : Return pointer to `LogController`.
func NewLogController() *LogController {
	return new(LogController)
}

// Index : index
func (lc *LogController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"OK",
	))
}

// SerpDwellTimeLogRequest : Struct for task viewing time log request body.
type SerpDwellTimeLogRequest struct {
	// UserID : The ID of user (worker).
	UserID int `json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `json:"task"`

	// ConditionID : User's condition ID that means group and task category.
	ConditionID int `json:"condition"`
}

// UpsertSerpDwellTimeLog : Create dwell time log.
// This inplicitly assumed that it will be called only once per second.
func (lc *LogController) UpsertSerpDwellTimeLog(c echo.Context) error {
	if lc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("LogController is nil"))
	}

	p := SerpDwellTimeLogRequest{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Failed bind request body.", p))
	}

	if err := c.Validate(&p); err != nil {
		if e, ok := err.(*models.InternalError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, e)
		}
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Failed to validate.", p))
	}

	log := &models.SerpDwellTimeLog{
		UserID:      p.UserID,
		TaskID:      p.TaskID,
		ConditionID: p.ConditionID,
	}
	if err := models.UpsertSerpDwellTimeLog(log); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		p,
	))
}

// ListSerpDwellTimeLogs : Returns all logs.
func (lc *LogController) ListSerpDwellTimeLogs(c echo.Context) error {
	if lc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("LogController is nil"))
	}

	logs, err := models.ListSerpDwellTimeLogs()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		logs,
	))
}

// PageDwellTimeLogRequest : Struct for each search result page viewing time log request body.
type PageDwellTimeLogRequest struct {
	// UserID : The ID of user (worker).
	UserID int `json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `json:"task"`

	// ConditionID : User's condition ID that means group and task category.
	ConditionID int `json:"condition"`

	// PageID : Page ID that user view.
	PageID int `json:"page"`

	// DwellTime : Time(sec.) that the user spend in search result page.
	DwellTime int `gorm:"not null;column:time_on_page"`
}

// CreatePageDwellTimeLog : Create new page dwell time log.
func (lc *LogController) CreatePageDwellTimeLog(c echo.Context) error {
	if lc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("LogController is nil"))
	}

	p := PageDwellTimeLogRequest{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Failed bind request body.", p))
	}

	if err := c.Validate(&p); err != nil {
		if e, ok := err.(*models.InternalError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, e)
		}
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Failed to validate.", p))
	}

	l := &models.PageDwellTimeLog{
		UserID:      p.UserID,
		TaskID:      p.TaskID,
		PageID:      p.PageID,
		ConditionID: p.ConditionID,
		DwellTime:   p.DwellTime,
	}

	if err := models.CreatePageDwellTimeLog(l); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		p,
	))
}

// ListPageDwellTimeLog : Return all logs.
func (lc *LogController) ListPageDwellTimeLog(c echo.Context) error {
	if lc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("LogController is nil"))
	}

	logs, err := models.ListPageDwellTimeLog()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		logs,
	))
}

// SearchPageEventLogRequest : Struct for page click log request body.
type SearchPageEventLogRequest struct {
	// ID : The ID of each log record.
	// ID string `json:"id"`

	// Uid : The ID of user (worker)
	UserID int `json:"user" validate:"required,numeric"`

	// TaskID : The ID of task that user working.
	TaskID int `json:"task" validate:"required,numeric"`

	// ConditionID : User's condition ID that means group and task category.
	ConditionID int `json:"condition" validate:"required,numeric"`

	// Time : User's page viewing time.
	Time int `json:"time" validate:"required,numeric"`

	// Page : The ID of page that user clicked.
	Page int `json:"page" validate:"required,numeric"`

	// Rank : Search result rank that user clicked.
	Rank int `json:"rank" validate:"required,numeric"`

	// IsVisible : Risk is visible or not.
	// `required` tag checks if the value of the field is the default value of the datatype.
	// So, to validate that the JSON set the bool, datatype is `*bool`.
	// See also : https://github.com/go-playground/validator/issues/319
	IsVisible *bool `json:"visible" validate:"required"`

	// Event : It is expected to be "click", "hover" or "paginate"
	Event string `json:"event" validate:"required,oneof=click hover paginate"`
}

// CreateSerpEventLog : Create new serp event log.
func (lc *LogController) CreateSerpEventLog(c echo.Context) error {
	if lc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("LogController is nil"))
	}

	p := SearchPageEventLogRequest{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Failed bind request body.", p))
	}

	if err := c.Validate(&p); err != nil {
		if e, ok := err.(*models.InternalError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, e)
		}
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Failed to validate.", p))
	}

	// if err := models.SerpEvent(p.Event).Valid(); err != nil {
	// 	return newErrResponse(c, http.StatusBadRequest, err, p)
	// }

	l := &models.SerpEventLog{
		UserID:      p.UserID,
		TaskID:      p.TaskID,
		ConditionID: p.ConditionID,
		Time:        p.Time,
		Page:        p.Page,
		Rank:        p.Rank,
		IsVisible:   *p.IsVisible,
		Event:       models.SerpEvent(p.Event),
	}
	if err := models.CreateSerpEventLog(l); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK), ""),
	)
}

// ListSerpEventLog : Return all logs.
func (lc *LogController) ListSerpEventLog(c echo.Context) error {
	if lc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("LogController is nil"))
	}

	logs, err := models.ListSerpEventLog()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		logs,
	))
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

// UpsertSearchSession : Create search session log.
func (lc *LogController) UpsertSearchSession(c echo.Context) error {
	if lc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("LogController is nil"))
	}

	p := SearchSessionRequest{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Failed bind request body.", p))
	}

	if err := c.Validate(&p); err != nil {
		if e, ok := err.(*models.InternalError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, e)
		}
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Failed to validate.", p))
	}

	log := &models.SearchSession{
		UserID:      p.UserID,
		TaskID:      p.TaskID,
		ConditionID: p.ConditionID,
	}
	if err := models.UpsertSearchSession(log); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		p,
	))
}

// ListSearchSession : Return all search session logs.
func (lc *LogController) ListSearchSession(c echo.Context) error {
	if lc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("LogController is nil"))
	}

	logs, err := models.ListSearchSession()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		logs,
	))
}
