package handlers

import (
	"net/http"
	"strconv"

	"github.com/Nomorehope/homework/models"
	"github.com/gin-gonic/gin"
)

func TasksList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.Tasks)
}

func GetTask(ctx *gin.Context) {
	idParam := ctx.Param("id")
	t_id, error := strconv.Atoi(idParam)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	for _, task := range models.Tasks {
		if task.Task_id == t_id {
			ctx.JSON(http.StatusOK, task)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func CreateTask(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.Task_id = len(models.Tasks) + 1
	models.Tasks = append(models.Tasks, task)
	ctx.JSON(http.StatusCreated, task)
}

func UpdateTask(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask models.Task
	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range models.Tasks {
		if task.Task_id == id {
			updatedTask.Task_id = id
			models.Tasks[i] = updatedTask
			ctx.JSON(http.StatusOK, updatedTask)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})

}

func DeleteTask(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	for i, task := range models.Tasks {
		if task.Task_id == id {
			models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func ListUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.Users)
}

func GetUser(ctx *gin.Context) {
	idUser := ctx.Param("id")
	id, error := strconv.Atoi(idUser)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}
	for _, user := range models.Users {
		if user.UID == id {
			ctx.JSON(http.StatusOK, user)
			return
		}
	}
}

func NewUser(ctx *gin.Context) {
	var user models.User
	if error := ctx.ShouldBindJSON(&user); error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	user.UID = len(models.Users) + 1
	models.Users = append(models.Users, user)
	ctx.JSON(http.StatusCreated, user)
}

func UpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("id") // idUser
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}
	var updatedUser models.User
	if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // err.Error()
		return
	}
	for i, user := range models.Users {
		if user.UID == id {
			updatedUser.UID = id
			models.Users[i] = updatedUser
			ctx.JSON(http.StatusOK, updatedUser)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id") // idUser
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
	}
	for i, user := range models.Users {
		if user.UID == id {
			models.Users = append(models.Users[:i], models.Users[i+1:]...) // remove user
			ctx.JSON(http.StatusOK, gin.H{"message": "User deleted"})      // return message
			return
		}
	}
}
