package usecase

type Tasks []Task

type Task struct {
	ID          int
	UserID      int
	Title       string
	Description string
}

func (t *Task) IsEmpty() bool {
	return t.ID == 0 || t.UserID == 0 || t.Title == "" || t.Description == ""
}
