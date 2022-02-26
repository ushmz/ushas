package models

// Task : Struct for Task information.
type Task struct {
	// ID : The ID of task
	ID int `db:"id" json:"id"`

	// Query : Search query for this task.
	Query string `db:"query" json:"query"`

	// Title : Title of this task.
	Title string `db:"title" json:"title"`

	// Description : Description text of task.
	Description string `db:"description" json:"description"`

	// SearchUrl : Url used in this task.
	SearchUrl string `db:"search_url" json:"searchUrl"`
}

// TaskInfo : Struct for response of which task is assigned.
type TaskInfo struct {
	// GroupId : The ID assigned to the pair of "task IDs" and "condition ID"
	GroupId int

	// ConditionId : Assigned condition ID
	ConditionId int

	// TaskIds : Shows the IDs that user perform
	TaskIds []int
}
