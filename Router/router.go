package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		fmt.Println("hello")
		c.Next()
		fmt.Println("Goodbye!")
	})
	r.GET("/api/todos", getTodos)
	r.GET("/api/todos/:id", getTodosByIDHandler)
	r.POST("/api/todos", postTodoHandler)
	r.DELETE("/api/todos/:id", deleteTodosByIDHandler)
	r.PUT("/api/todos/:id", putTodosByIDHandler)
	return r
}
