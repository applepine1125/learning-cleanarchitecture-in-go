package domain

type Tasks []Task

type Task struct {
	ID           int
	UserID       int
	UserFullName string
	Title        string
	Description  string
}
