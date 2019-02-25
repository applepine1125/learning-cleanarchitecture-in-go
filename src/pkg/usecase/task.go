package usecase

type Tasks []Task

type Task struct {
	ID          int
	UserID      int
	Title       string
	Description string
}
