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
	idParam := ctx.Param("t_id")
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
