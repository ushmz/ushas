package views

import "ushas/models"

type TaskView struct {
	// ID : The ID of task
	ID int `json:"id"`

	// Query : Search query for this task.
	Query string `json:"query"`

	// Title : Title of this task.
	Title string `json:"title"`

	// Description : Description text of task.
	Description string `json:"description"`

	// SearchUrl : Url used in this task.
	SearchUrl string `json:"searchUrl"`
}

func NewTaskView(t *models.Task) *TaskView {
	v := &TaskView{
		ID:          t.ID,
		Query:       t.Query,
		Title:       t.Title,
		Description: t.Description,
		SearchUrl:   t.SearchUrl,
	}
	return v
}
