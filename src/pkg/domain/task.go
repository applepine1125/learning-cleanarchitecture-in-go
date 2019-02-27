package domain

type Tasks []Task

type Task struct {
	ID           int
	UserID       int
	UserFullName string
	Title        string
	Description  string
}

func (t *Task) IsEmpty() bool {
	return t.ID == 0 || t.UserID == 0 || t.UserFullName == "" || t.Title == "" || t.Description == ""
}
