package domain

import "fmt"

type Users []User

type User struct {
	ID        int64
	FirstName string
	LastName  string
}

func (u *User) GetFullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func (u *User) IsEmpty() bool {
	return u.ID == 0 || u.FirstName == "" || u.LastName == ""
}

func (us *Users) FindById(id int64) User {
	for _, u := range *us {
		if u.ID == id {
			return u
		}
	}
	return User{}
}
