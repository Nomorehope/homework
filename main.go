package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm" // TODO: добавить в зависимости

	"github.com/Nomorehope/homework/handlers"
	"github.com/Nomorehope/homework/models"
	"github.com/gin-gonic/gin"
)

var DB *gorm.DB

func initDataBase() {
	dsn := "host=localhost user=postgres password=q1w2e3r4 dbname=study port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // TODO: добавить в зависимости
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	// Миграция моделей
	DB.AutoMigrate(&models.Task{}, &models.User{})
}

func main() {
	initDataBase() // Подключение и миграция базы данных

	r := gin.Default()

	// Передача базы данных в хендлеры
	r.Use(func(c *gin.Context) {
		c.Set("db", DB)
		c.Next()
	})

	// Определение маршрутов
	r.GET("/tasks", handlers.TasksList)
	r.GET("/tasks/:id", handlers.GetTask)
	r.POST("/tasks", handlers.CreateTask)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.DELETE("/tasks/:id", handlers.DeleteTask)

	r.GET("/users", handlers.ListUsers)
	r.GET("/users/:id", handlers.GetUser)
	r.POST("/users", handlers.NewUser)
	r.PUT("/users/:id", handlers.UpdateUser) // TODO: сделать обновление
	r.DELETE("users/:id", handlers.DeleteUser)

	r.Run(":8080")
}
