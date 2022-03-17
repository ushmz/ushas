package views

import "ushas/models"

// TaskView : Response data of "/task" endpoint.
type TaskView struct {
	// ID : The ID of task
	ID int `json:"id"`

	// Query : Search query for this task.
	Query string `json:"query"`

	// Title : Title of this task.
	Title string `json:"title"`

	// Description : Description text of task.
	Description string `json:"description"`

	// SearchURL : Url used in this task.
	SearchURL string `json:"searchUrl"`
}

// NewTaskView : Returns single task response data.
func NewTaskView(t *models.Task) *TaskView {
	v := &TaskView{
		ID:          t.ID,
		Query:       t.Query,
		Title:       t.Title,
		Description: t.Description,
		SearchURL:   t.SearchURL,
	}
	return v
}
