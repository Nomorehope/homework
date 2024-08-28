package main

import (
	"github.com/Nomorehope/homework/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Определение маршрутов
	r.GET("/tasks", handlers.TasksList)
	r.GET("/tasks/:id", handlers.GetTask)
	r.POST("/tasks", handlers.CreateTask)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.DELETE("/tasks/:id", handlers.DeleteTask)

	r.GET("/users", handlers.ListUsers)
	r.GET("/users/:id", handlers.GetUser)
	r.POST("/users", handlers.NewUser)

	r.Run(":8080")
}
