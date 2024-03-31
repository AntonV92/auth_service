package models

type User struct {
	Id        int
	Username  string
	Email     string
	Password  string
	Confirmed bool
	CreatedAt string
	UpdatedAt string
}
