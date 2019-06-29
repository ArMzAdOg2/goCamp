package Router

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/api/todos", getTodos)
	r.GET("/api/todos/:id", getTodosByIDHandler)
	r.POST("/api/todos", postTodoHandler)
	r.DELETE("/api/todos/:id", deleteTodosByIDHandler)
	r.PUT("/api/todos/:id", putTodosByIDHandler)
	return r
}
