package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"time")

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

var db *sql.DB

func GetUser() ([]User, error) {
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

func AddUser(user User) error {
	const query = `INSERT INTO users(nameUser,lastname,email,created_at) VALUES ($1,$2,$3,$4)`
	_, err := db.Exec(query, user.NameUser, user.LastName, user.Email, user.CreatedAt)
	return err
}

func UpdateUser(user User) error {
	const query = `UPDATE  users SET nameUser = $1, lastname = $2 WHERE email = $3`
	_, err := db.Exec(query, user.NameUser, user.LastName, user.Email)
	return err
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("test.html")
	if err != nil {
		fmt.Println(err)
	}
	items := struct {
		Version string
	}{
		Version: "0.0.1",
	}
	t.Execute(w, items)
}

func main() {

	http.HandleFunc("/", serveStatic)
	http.ListenAndServe(":8080", nil)

	var err error
	r := gin.Default()
	r.GET("/get", func(context *gin.Context) {
		result, err := GetUser()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"status": "internal error uno :" + err.Error()})
			return
		}
		context.JSON(http.StatusOK, result)

	})

	r.POST("/post", func(context *gin.Context) {
		var b User
		if context.Bind(&b) == nil {
			b.CreatedAt = time.Now()
			if err := AddUser(b); err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"status": "internal error post" + err.Error()})
				return
			}
			context.JSON(http.StatusOK, gin.H{"status": "200"})
			return
		}
		context.JSON(http.StatusUnprocessableEntity, gin.H{"status": "500"})

	})

	r.POST("/update", func(context *gin.Context) {
		var b User
		if context.Bind(&b) == nil {
			b.CreatedAt = time.Now()
			if err := UpdateUser(b); err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"status": "internal error post" + err.Error()})
				return
			}
			context.JSON(http.StatusOK, gin.H{"status": "200"})
			return
		}
		context.JSON(http.StatusUnprocessableEntity, gin.H{"status": "500"})
	})

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
