package main 

import (
	"github.com/gin-gonic/gin"
)

func main() {
	memoryStorage := NewMemoryStorage()
	handler := NewHandler(memoryStorage)

	router := gin.Default()

	router.POST("/tasks", handler.CreateTodoTask)
	router.GET("/tasks/:id", handler.GetTask)
	router.PUT("/tasks/:id", handler.UpdateTask)
	router.DELETE("/tasks/:id", handler.DeleteTask)

	router.Run()
}