package domain

import "time"

type User struct {
	NameUser  string    `json:"nameUser" binding:"required"`
	LastName  string    `json:"lastname" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}