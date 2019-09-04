package repository

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"

	_ "github.com/lib/pq"
)

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



var db *sql.DB

func OpenDB(err error, r *gin.Engine) {
	// open a connection to the database
	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DbHost, DbUser, DbPassword, DbName)
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	// do not forget to close the connection
	defer db.Close()
	// ensuring the table is created
	_, err = db.Query(Migration)
	if err != nil {
		log.Println("failed to run migrations", err.Error())
		return
	}
	// running the http server
	log.Println("running..")
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}

func (user User) GetUser() ([]User, error) {
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


func (user User) AddUser(newUser User)  error {
	const query = `INSERT INTO users(nameUser,lastname,email,created_at) VALUES ($1,$2,$3,$4)`
	_, err := db.Exec(query, newUser.NameUser, newUser.LastName, newUser.Email, newUser.CreatedAt)
	return err
}


func (user User) UpdateUser (updateUser User) error{
	const query = `UPDATE  users SET nameUser = $1, lastname = $2 WHERE email = $3`
	_, err := db.Exec(query, updateUser.NameUser, updateUser.LastName, updateUser.Email)
	return err
}

func (user User ) Imprimir1(interfacePrueba UserRepository) {
	interfacePrueba.Imprimir()
}

func (user User )Imprimir()   {
	println(user.NameUser)
}
