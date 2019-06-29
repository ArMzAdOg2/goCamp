package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"github.com/ArMzAdOg2/goCamp/router"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Student struct {
	ID   int
	Name string
}

type todo struct {
	ID     int    "json:id"
	Title  string "json:title"
	Status string "json:status"
}

var todos = []todo{}

func getStudentHandler(c *gin.Context) {
	r := router:SetupRouter()
	c.JSON(200, "OK")
}

func postStudentHandler(c *gin.Context) {
	c.JSON(200, "OK")
}

func getDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getTodos(c *gin.Context) {
	db, err := getDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	stmt, err := db.Prepare("select id,title,status from todos;")
	rows, err := stmt.Query()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	for rows.Next() {
		t := todo{}
		err = rows.Scan(&t.ID, &t.Title, &t.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		todos = append(todos, t)
	}
	c.JSON(http.StatusOK, todos)
	todos = []todo{}
}

func getTodosByIDHandler(c *gin.Context) {
	db, err := getDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	id := c.Param("id")
	stmt, err := db.Prepare("select id,title,status from todos where id=$1;")
	row := stmt.QueryRow(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	todo := todo{}
	err = row.Scan(&todo.ID, &todo.Title, &todo.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"id":     todo.ID,
		"title":  todo.Title,
		"status": todo.Status,
	})

	defer db.Close()
}

func postTodoHandler(c *gin.Context) {
	db, err := getDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	todo := todo{}
	c.ShouldBindJSON(&todo)
	stmt, err := db.Prepare("insert into todos (title,status) values ($1,$2) returning id;")
	row := stmt.QueryRow(todo.Title, todo.Status)
	var id int
	err = row.Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	stmt, err = db.Prepare("select id,title,status from todos where id=$1;")
	row = stmt.QueryRow(id)
	err = row.Scan(&todo.ID, &todo.Title, &todo.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(201, gin.H{
		"id":     todo.ID,
		"title":  todo.Title,
		"status": todo.Status,
	})
	defer db.Close()
}

func deleteTodosByIDHandler(c *gin.Context) {
	db, err := getDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	id := c.Param("id")
	stmt, err := db.Prepare("delete from todos where id=$1;")
	_, err = stmt.Query(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
	defer db.Close()
}

func putTodosByIDHandler(c *gin.Context) {
	db, err := getDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	id := c.Param("id")
	todo := todo{}
	c.ShouldBindJSON(&todo)
	stmt, err := db.Prepare("update todos set title = '$2', status = '$3' where id=$1;")
	row := stmt.QueryRow(id, todo.Title, todo.Status)
	err = row.Scan(&todo.ID, &todo.Title, &todo.Status)
	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"title":  todo.Title,
		"status": todo.Status,
	})
	defer db.Close()
}

func main() {
	r := setupRouter()
	r.Run(":1234")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		fmt.Println("hello")
		token := c.GetHeader("Authorization")
		fmt.Println(token)
		if token != "Bearer 123" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Wrong token",
			})
			return
		}
		c.Next()
		fmt.Println("token :", token)
		fmt.Println("Goodbye!")
	})
	r.GET("/api/todos", getTodos)
	r.GET("/api/todos/:id", getTodosByIDHandler)
	r.POST("/api/todos", postTodoHandler)
	r.DELETE("/api/todos/:id", deleteTodosByIDHandler)
	r.PUT("/api/todos/:id", putTodosByIDHandler)
	return r
}
