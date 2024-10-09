package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Nomorehope/homework/handlers"
	"github.com/Nomorehope/homework/middleware" // middleware для авторизации
	"github.com/Nomorehope/homework/models"
	"github.com/gin-gonic/gin"
)

var DB *gorm.DB

// Инициализация базы данных
func initDataBase() {
	dsn := "host=localhost user=postgres password=q1w2e3r4 dbname=study port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	// Миграция моделей
	DB.AutoMigrate(&models.Task{}, &models.User{})
}

func main() {
	initDataBase() // Подключение и миграция базы данных

	r := gin.Default()

	// Передача базы данных в хендлеры через контекст
	r.Use(func(c *gin.Context) {
		c.Set("db", DB)
		c.Next()
	})

	// Публичные маршруты (не требуют авторизации)
	public := r.Group("/")
	{
		public.POST("/login", handlers.Login)   // Логин
		public.POST("/users", handlers.NewUser) // Регистрация нового пользователя
	}

	// Защищённые маршруты (требуют JWT авторизации)
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware()) // Подключение middleware для проверки JWT
	{
		// Маршруты для работы с задачами
		protected.GET("/tasks", handlers.TasksList)
		protected.GET("/tasks/:id", handlers.GetTask)
		protected.POST("/tasks", handlers.CreateTask)
		protected.PUT("/tasks/:id", handlers.UpdateTask)
		protected.DELETE("/tasks/:id", handlers.DeleteTask)

		// Маршруты для управления пользователями
		protected.GET("/users", handlers.ListUsers)
		protected.GET("/users/:id", handlers.GetUser)
		protected.PUT("/users/:id", handlers.UpdateUser)
		protected.DELETE("/users/:id", handlers.DeleteUser)
	}

	r.Run(":8080")
}
