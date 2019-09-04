package repository

import "crud-golang/domain"

type UserRepository interface {
	GetUser () ([]domain.User, error)
	AddUser(user domain.User) error
	UpdateUser(user domain.User) error

	Imprimir()
}