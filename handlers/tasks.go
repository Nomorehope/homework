package handlers

import (
	"net/http"
	"strconv"

	"github.com/Nomorehope/homework/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TasksList(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var tasks []models.Task
	db.Find(&tasks)
	ctx.JSON(http.StatusOK, tasks)
}

func GetTask(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	idParam := ctx.Param("id")
	t_id, error := strconv.Atoi(idParam)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var task models.Task
	if result := db.First(&task, t_id); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	ctx.JSON(http.StatusOK, task)

}

func CreateTask(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&task) // Сохраняем задачу в базе данных
	ctx.JSON(http.StatusCreated, task)
}

func UpdateTask(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	idParam := ctx.Param("id")
	t_id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var task models.Task
	if err := db.First(&task, t_id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Task not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&task)
	ctx.JSON(http.StatusOK, task)
}

func DeleteTask(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var task models.Task

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	if err := db.First(&task, id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	db.Delete(&task)
	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted"})

}

func ListUsers(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var users []models.User
	db.Find(&users)
	ctx.JSON(http.StatusOK, users)
}

func GetUser(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	idUser := ctx.Param("id")
	u_id, error := strconv.Atoi(idUser)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}
	var user models.User
	if result := db.First(&user, u_id); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func NewUser(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var user models.User
	if error := ctx.ShouldBindJSON(&user); error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	db.Create(&user)
	ctx.JSON(http.StatusCreated, user)
}

func UpdateUser(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	idParam := ctx.Param("id") // idUser
	u_id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // err.Error()
		return
	}
	if err := db.First(&user, u_id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	db.Save(&user)
	ctx.JSON(http.StatusOK, user)
}

func DeleteUser(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	idParam := ctx.Param("id")
	u_id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if result := db.First(&user, u_id); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	db.Delete(&user) // Удаляем пользователя
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
