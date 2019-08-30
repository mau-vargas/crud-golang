package repository

type UserRepository interface {
	GetUser () ([]User, error)
	AddUser(user User) error
	UpdateUser(user User) error
}