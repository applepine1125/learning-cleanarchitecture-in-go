package domain

type User struct {
	ID        int64
	FirstName string
	LastName  string
}

func (u *User) getFullName() string {
	return u.FirstName + u.LastName
}
