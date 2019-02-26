package database

import (
	"github.com/applepine1125/learning-cleanarchitecture-in-go/src/pkg/domain"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u domain.User) error {
	err := repo.Exec("INSERT INTO users (first_name, last_name) VALUES (?,?)", u.FirstName, u.LastName)
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) FindById(id int) (user domain.User) {
	var firstname string
	var lastname string

	row := repo.QueryRow("SELECT * FROM users WHERE id = ?", id)
	row.Scan(&firstname, &lastname)

	user.ID = id
	user.FirstName = firstname
	user.LastName = lastname
	return user
}

func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	rows, err := repo.Query("SELECT * FROM users")
	defer rows.Close()
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var id int
		var firstname string
		var lastname string
		if err := rows.Scan(&id, &firstname, &lastname); err != nil {
			continue
		}
		user := domain.User{
			ID:        id,
			FirstName: firstname,
			LastName:  lastname,
		}
		users = append(users, user)
	}
	return users, nil
}
