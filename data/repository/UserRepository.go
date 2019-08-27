/*package main

import (
	"database/sql"
	"time"
)




type userRepository struct {
	db *sql.DB
}


const (
	DbHost     = "crudgolang"
	DbUser     = "postgress-crud"
	DbPassword = "crudpassword"
	DbName     = "crud_name"
	Migration  = `CREATE TABLE IF NOT EXISTS users (
id serial PRIMARY KEY,
nameUser text NOT NULL,
lastname text NOT NULL,
email text NOT NULL,
created_at timestamp with time zone DEFAULT current_timestamp)`
)


type User struct {
	NameUser  string    `json:"nameUser" binding:"required"`
	LastName  string    `json:"lastname" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}



func  (db *userRepository)GetUser() ([]User, error) {
	const query = `SELECT nameUser, lastname ,email, created_at FROM users ORDER BY created_at DESC LIMIT 100`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	results := make([]User, 0)

	for rows.Next() {
		var nameUser string
		var lastname string
		var email string
		var created_at time.Time

		err = rows.Scan(&nameUser, &lastname, &email, &created_at)
		if err != nil {
			return nil, err
		}

		results = append(results, User{nameUser, lastname, email, created_at})
	}
	return results, nil

}


func (db *userRepository) AddUser(user User) error {
	const query = `INSERT INTO users(nameUser,lastname,email,created_at) VALUES ($1,$2,$3,$4)`
	_, err := db.Exec(query, user.NameUser, user.LastName, user.Email, user.CreatedAt)
	return err
}

func (db *userRepository) UpdateUser(user User) error {
	const query = `UPDATE  users SET nameUser = $1, lastname = $2 WHERE email = $3`
	_, err := db.Exec(query, user.NameUser, user.LastName, user.Email)
	return err
}
*/