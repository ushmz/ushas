package models

import (
	"fmt"
	"ushas/database"
)

// Task : Struct for Task information.
type Task struct {
	// ID : The ID of task
	ID int `gorm:"unique;not null;column:id" json:"id"`

	// Query : Search query for this task.
	Query string `gorm:"not null;column:query" json:"query"`

	// Title : Title of this task.
	Title string `gorm:"not null;column:title" json:"title"`

	// Description : Description text of task.
	Description string `gorm:"not null;column:description" json:"description"`

	// SearchURL : Url used in this task.
	SearchURL string `gorm:"not null;column:search_url" json:"searchUrl"`
}

// TaskInfo : Struct for response of which task is assigned.
type TaskInfo struct {
	// GroupID : The ID assigned to the pair of "task IDs" and "condition ID"
	GroupID int

	// ConditionID : Assigned condition ID
	ConditionID int

	// TaskIds : Shows the IDs that user perform
	TaskIds []int
}

// CreateTask : Create new record into table.
func CreateTask(t *Task) error {
	db := database.GetDB()
	if err := db.Create(t).Error; err != nil {
		return RaiseInternalServerError(
			err,
			"Failed to create new `Task` resource",
		)
	}
	return nil
}

// GetTaskByID : Get single record from table by ID.
func GetTaskByID(id int) (*Task, error) {
	t := new(Task)
	db := database.GetDB()
	if err := db.Where("id = ?", id).First(t).Error; err != nil {
		return t, RaiseNotFoundError(
			err,
			fmt.Sprintf("Task for ID %d is not found", id),
		)
	}
	return t, nil
}

// ListTasks : Get all records from table.
func ListTasks() ([]Task, error) {
	tasks := []Task{}
	db := database.GetDB()
	if err := db.Find(&tasks).Error; err != nil {
		return tasks, RaiseInternalServerError(
			err,
			"Failed to fetch all Task resources",
		)
	}
	return tasks, nil
}

// UpdateTask : Update a record with given ID in table.
func UpdateTask(t *Task) error {
	db := database.GetDB()
	if err := db.Save(t).Error; err != nil {
		return RaiseInternalServerError(
			err,
			fmt.Sprintf("Failed to update Task resource of ID %d", t.ID),
			t,
		)
	}
	return nil
}

// DeleteTask : Delete a record with given ID from table.
func DeleteTask(id int) error {
	db := database.GetDB()
	if err := db.Delete(&Task{}, id).Error; err != nil {
		return RaiseInternalServerError(
			err,
			fmt.Sprintf("Failed to delete Task resource of ID %d", id),
		)
	}
	return nil
}
