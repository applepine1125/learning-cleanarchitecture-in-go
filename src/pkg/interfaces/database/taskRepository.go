package database

import (
	"github.com/applepine1125/learning-cleanarchitecture-in-go/src/pkg/usecase"
)

type TaskRepository struct {
	SqlHandler
}

func (repo *TaskRepository) Store(t usecase.Task) error {
	err := repo.Exec("INSERT INTO tasks (user_id, title, description) VALUES (?,?,?)", t.UserID, t.Title, t.Description)
	if err != nil {
		return err
	}

	return nil
}

func (repo *TaskRepository) FindById(id int) (task usecase.Task) {
	var userid int
	var title string
	var description string

	row := repo.QueryRow("SELECT * FROM users WHERE id = ?", id)
	row.Scan(&userid, &title, &description)

	task.ID = id
	task.UserID = userid
	task.Title = title
	task.Description = description
	return task
}

func (repo *TaskRepository) FindAll() (tasks usecase.Tasks, err error) {
	rows, err := repo.Query("SELECT * FROM tasks")
	defer rows.Close()
	if err != nil {
		return tasks, err
	}

	for rows.Next() {
		var id int
		var userid int
		var title string
		var description string
		if err := rows.Scan(&id, &userid, &title, &description); err != nil {
			continue
		}
		task := usecase.Task{
			ID:          id,
			UserID:      userid,
			Title:       title,
			Description: description,
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
