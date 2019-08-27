package main

import (
	"crud-golang/data/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
	"time"
)


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


	//http.HandleFunc("/", serveStatic)
	//http.ListenAndServe(":8080", nil)

    var err error
	r := gin.Default()
	r.GET("/get", func(context *gin.Context) {
		result, err := repository.GetUser()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"status": "internal error uno :" + err.Error()})
			return
		}
		context.JSON(http.StatusOK, result)

	})

	r.POST("/post", func(context *gin.Context) {
		var b repository.User
		if context.Bind(&b) == nil {
			b.CreatedAt = time.Now()
			if err := repository.AddUser(b); err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"status": "internal error post" + err.Error()})
				return
			}
			context.JSON(http.StatusOK, gin.H{"status": "200"})
			return
		}
		context.JSON(http.StatusUnprocessableEntity, gin.H{"status": "500"})

	})

	r.POST("/update", func(context *gin.Context) {
		var b repository.User
		if context.Bind(&b) == nil {
			b.CreatedAt = time.Now()
			if err := repository.UpdateUser(b); err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"status": "internal error post" + err.Error()})
				return
			}
			context.JSON(http.StatusOK, gin.H{"status": "200"})
			return
		}
		context.JSON(http.StatusUnprocessableEntity, gin.H{"status": "500"})
	})


	repository.OpenDB(err,r)

}
