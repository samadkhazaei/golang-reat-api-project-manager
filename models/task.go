package models

type Task struct {
	ID          int    `json:"id"`
	ProjectID   int    `json:"project_id"`
	Status      string `json:"status"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
