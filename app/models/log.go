package models

import (
	"time"
	"ushas/database"
)

type SerpViewingLog struct {
	// UserID : The ID of user (worker)
	UserID int `gorm:"not null;column:user_id"`

	// TaskID : The ID of task that user working.
	TaskID int `gorm:"not null;column:task_id"`

	// ConditionId : User's condition ID that means group and task category.
	ConditionID int `gorm:"not null;column:condition_id"`

	// DwellTime : Time(sec.) that the user spend in SERP
	DwellTime int `gorm:"not null;column:time_on_page"`

	// CreatedAt :
	CreatedAt time.Time `gorm:"not null;column:created_at"`

	// UpdatedAt :
	UpdatedAt time.Time `gorm:"not null;column:updated_at"`
}

func CreateSerpViewingLog(l *SerpViewingLog) error {
	db := database.GetDB()
	if err := db.Create(l).Error; err != nil {
		return RaiseInternalServerError(
			err,
			"Failed to create new Log resource",
		)
	}
	return nil
}

func ListSerpViewingLogs() ([]SerpViewingLog, error) {
	logs := []SerpViewingLog{}
	db := database.GetDB()
	if err := db.Find(logs).Error; err != nil {
		return logs, RaiseInternalServerError(
			err,
			"Failed to fetch all Logs",
		)
	}
	return logs, nil
}

type PageViewingLog struct {
	// UserID : The ID of user (worker)
	UserID int `gorm:"not null;column:user_id"`

	// TaskID : The ID of task that user working.
	TaskID int `gorm:"not null;column:task_id"`

	// PageID : Page ID that user view.
	PageID int `gorm:"not null;column:page_id"`

	// ConditionID : User's condition ID that means group and task category.
	ConditionID int `gorm:"not null;column:condition_id"`

	// DwellTime : Time(sec.) that the user spend in SERP
	DwellTime int `gorm:"not null;column:time_on_page"`

	// CreatedAt :
	CreatedAt time.Time `gorm:"not null;column:created_at"`

	// UpdatedAt :
	UpdatedAt time.Time `gorm:"not null;column:updated_at"`
}

func CreatePageViewingLog(l *PageViewingLog) error {
	db := database.GetDB()
	if err := db.Create(l).Error; err != nil {
		return RaiseInternalServerError(
			err,
			"Failed to create new Log resource",
		)
	}
	return nil
}

func ListPageViewingLog() ([]PageViewingLog, error) {
	logs := []PageViewingLog{}
	db := database.GetDB()
	if err := db.Find(logs).Error; err != nil {
		return logs, RaiseInternalServerError(
			err,
			"Failed to fetch all Logs",
		)
	}
	return logs, nil
}

type SerpEventLog struct {
	// ID : The ID of each log record.
	ID string `gorm:"not null;column:id"`

	// UserId : The ID of user (worker)
	UserID int `gorm:"not null;column:user_id"`

	// TaskId : The ID of task that user working.
	TaskID int `gorm:"not null;column:task_id"`

	// ConditionId : User's condition ID that means group and task category.
	ConditionID int `gorm:"not null;column:condition_id"`

	// Time : User's page viewing time.
	Time int `gorm:"not null;column:time_on_page"`

	// Page : The ID of page that user clicked.
	Page int `gorm:"not null;column:serp_page"`

	// Rank : Search result rank that user clicked.
	Rank int `gorm:"not null;column:serp_rank"`

	// IsVisible : Risk is visible or not.
	IsVisible bool `gorm:"not null;column:is_visible"`

	// Event : It is expected to be "click", "hover" or "paginate"
	Event string `gorm:"not null;column:event"`

	// CreatedAt :
	CreatedAt time.Time `gorm:"not null;not null;column:created_at"`

	// UpdatedAt :
	UpdatedAt time.Time `gorm:"not null;column:updated_at"`
}

func CreateSerpEventLog(l *PageViewingLog) error {
	db := database.GetDB()
	if err := db.Create(l).Error; err != nil {
		return RaiseInternalServerError(
			err,
			"Failed to create new Log resource",
		)
	}
	return nil
}

func ListSerpEventLog() ([]SerpEventLog, error) {
	logs := []SerpEventLog{}
	db := database.GetDB()
	if err := db.Find(logs).Error; err != nil {
		return logs, RaiseInternalServerError(
			err,
			"Failed to fetch all Logs",
		)
	}
	return logs, nil
}

type SearchSession struct {
	// UserID : Assigned ID of user (worker)
	UserID int `gorm:"not null;column:user_id"`

	// TaskID : The ID of task that user working.
	TaskID int `gorm:"not null;column:task_id"`

	// ConditionID : User's condition ID that means group and task category.
	ConditionID int `gorm:"not null;column:condition_id"`

	// StartedAt :
	StartedAt time.Time `gorm:"not null;column:started_at"`

	// EndedAt :
	EndedAt time.Time `gorm:"not null;column:ended_at"`
}

func CreateSearchSession(l SearchSession) error {
	db := database.GetDB()
	if err := db.Create(l).Error; err != nil {
		return RaiseInternalServerError(
			err,
			"Failed to create new Log resource",
		)
	}
	return nil
}

func ListSearchSession() ([]SearchSession, error) {
	logs := []SearchSession{}
	db := database.GetDB()
	if err := db.Find(logs).Error; err != nil {
		return logs, RaiseInternalServerError(
			err,
			"Failed to fetch all Logs",
		)
	}
	return logs, nil
}
