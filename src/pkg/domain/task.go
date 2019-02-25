package domain

type Tasks []Task
type Task struct {
	ID          int64
	UserID      int64
	Title       string
	Description string
}
