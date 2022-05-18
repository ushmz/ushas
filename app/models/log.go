package models

import (
	"time"
	"ushas/database"

	"golang.org/x/xerrors"
	"gorm.io/gorm/clause"
)

// SerpDwellTimeLog : Which user does which task, and how many time did they spend on SERP.
type SerpDwellTimeLog struct {
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

// TableName : Overrides table name userd by `SerpDwellTimeLog` struct.
func (*SerpDwellTimeLog) TableName() string {
	return "logs_serp_dwell_time"
}

// UpsertSerpDwellTimeLog : Upserts dwell time log
// This inplicitly assumed that it will be called only once per second.
func UpsertSerpDwellTimeLog(l *SerpDwellTimeLog) error {
	db := database.GetDB()
	// If the key ("user_id" and "task_id") is duplicated,
	// update "time_on_page" and "ended_at" value, otherwise insert new record.
	// MySQL query is like following.
	// INSERT INTO `logs_serp_dwell_time` ... ON DUPLICATE KEY UPDATE ...;
	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "task_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"time_on_page", "updated_at"}),
	}).Create(l).Error
	if err != nil {
		return translateGormError(err, l)
	}
	return nil
}

// ListSerpDwellTimeLogs : Gets all records of dwell time logs.
func ListSerpDwellTimeLogs() ([]SerpDwellTimeLog, error) {
	logs := []SerpDwellTimeLog{}
	db := database.GetDB()
	if err := db.Find(logs).Error; err != nil {
		return logs, translateGormError(err, nil)
	}
	return logs, nil
}

// PageDwellTimeLog : Which user spend how many time on search result page.
type PageDwellTimeLog struct {
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

// TableName : Overrides table name userd by `PageViewingLog` struct.
func (*PageDwellTimeLog) TableName() string {
	return "logs_page_dwell_time"
}

// CreatePageDwellTimeLog : Creates new record into table.
func CreatePageDwellTimeLog(l *PageDwellTimeLog) error {
	db := database.GetDB()
	if err := db.Create(l).Error; err != nil {
		return translateGormError(err, l)
	}
	return nil
}

//ListPageDwellTimeLog : Gets all records from table.
func ListPageDwellTimeLog() ([]PageDwellTimeLog, error) {
	logs := []PageDwellTimeLog{}
	db := database.GetDB()
	if err := db.Find(logs).Error; err != nil {
		return logs, translateGormError(err, nil)
	}
	return logs, nil
}

//SerpEvent : Kinds of events on SERP.
type SerpEvent string

const (
	// CLICK : The user click and view page on SERP.
	CLICK = SerpEvent("click")

	// HOVER : The user put cursor on page link on SERP.
	HOVER = SerpEvent("hover")

	// PAGINATE : The user go next/previous page on SERP.
	PAGINATE = SerpEvent("paginate")
)

// Valid : Check given serp event value is valid or not.
func (e SerpEvent) Valid() error {
	switch e {
	case CLICK:
		return nil
	case HOVER:
		return nil
	case PAGINATE:
		return nil
	default:
		return xerrors.New("Invalid SERP event")
	}
}

// SerpEventLog : Behavior log such as click event, hover event and paginate event.
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
	// This field might be better to use Enum.
	Event SerpEvent `gorm:"not null;column:event"`

	// CreatedAt :
	CreatedAt time.Time `gorm:"not null;not null;column:created_at"`

	// UpdatedAt :
	UpdatedAt time.Time `gorm:"not null;column:updated_at"`
}

// TableName : Overrides table name userd by `SerpEventLog` struct.
func (*SerpEventLog) TableName() string {
	return "logs_event"
}

// CreateSerpEventLog : Creates new record into table.
func CreateSerpEventLog(l *SerpEventLog) error {
	db := database.GetDB()
	if err := db.Create(l).Error; err != nil {
		return translateGormError(err, l)
	}
	return nil
}

// ListSerpEventLog : Gets all records from table.
func ListSerpEventLog() ([]SerpEventLog, error) {
	logs := []SerpEventLog{}
	db := database.GetDB()
	if err := db.Find(logs).Error; err != nil {
		return logs, translateGormError(err, nil)
	}
	return logs, nil
}

// SearchSession : When the user starts and ends each task.
type SearchSession struct {
	// UserID : Assigned ID of user (worker)
	UserID int `gorm:"not null;column:user_id"`

	// TaskID : The ID of task that user working.
	TaskID int `gorm:"not null;column:task_id"`

	// ConditionID : User's condition ID that means group and task category.
	ConditionID int `gorm:"not null;column:condition_id"`

	// StartedAt : When user starts the task.
	StartedAt time.Time `gorm:"not null;column:started_at"`

	// EndedAt : When user ends the task.
	EndedAt time.Time `gorm:"not null;column:ended_at"`
}

// TableName : Overrides table name userd by `SearchSession` struct.
func (*SearchSession) TableName() string {
	// [TODO] `search_sessions` is better.
	return "search_session"
}

// UpsertSearchSession : Upserts search session log.
func UpsertSearchSession(l *SearchSession) error {
	db := database.GetDB()
	// If the key ("user_id" and "task_id") is duplicated, update "ended_at" value,
	// otherwise insert new record.
	// MySQL query is like following.
	// INSERT INTO `search_session` ... ON DUPLICATE KEY UPDATE `ended_at`=VALUES(ended_at);
	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "task_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"ended_at"}),
	}).Create(l).Error
	if err != nil {
		return translateGormError(err, l)
	}
	return nil
}

// ListSearchSession : Gets all records from table.
func ListSearchSession() ([]SearchSession, error) {
	logs := []SearchSession{}
	db := database.GetDB()
	if err := db.Find(logs).Error; err != nil {
		return logs, translateGormError(err, nil)
	}
	return logs, nil
}
