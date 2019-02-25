package domain

type Tasks []Task

type Task struct {
	ID           int
	UserFullName string
	Title        string
	Description  string
	Date         string
}
