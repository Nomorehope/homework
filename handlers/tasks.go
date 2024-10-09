package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/Nomorehope/homework/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtSecret = []byte("your_secret_key") // Секретный ключ для JWT

type Claims struct {
	UID                  string `json:"uid"`
	jwt.RegisteredClaims        // Используем этот заголовок
}

func TasksList(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var tasks []models.Task
	db.Find(&tasks)
	ctx.JSON(http.StatusOK, tasks)
}

func GetTask(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	idParam := ctx.Param("id")
	t_id, error := uuid.Parse(idParam)
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
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.Task_id = uuid.New()
	db := ctx.MustGet("db").(*gorm.DB)

	db.Create(&task) // Сохраняем задачу в базе данных
	ctx.JSON(http.StatusCreated, task)
}

func UpdateTask(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	idParam := ctx.Param("id")
	t_id, err := uuid.Parse(idParam)
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
	id, err := uuid.Parse(idParam)
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
	u_id, error := uuid.Parse(idUser)
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
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Хэширование пароля перед сохранением
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	user.Password = string(hashedPassword)

	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func UpdateUser(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	idParam := ctx.Param("id")       // idUser
	u_id, err := uuid.Parse(idParam) // Изменено с strconv.Atoi на uuid.Parse
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.First(&user, u_id).Error; err != nil { // Убедитесь, что используете .Error
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.UID = u_id // Убедитесь, что обновляете UID
	db.Save(&user)
	ctx.JSON(http.StatusOK, user)
}

func DeleteUser(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	idParam := ctx.Param("id")
	u_id, err := uuid.Parse(idParam)
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

func GenerateJWT(userID string) (string, error) {
	// Определяем стандартные claims (утверждения) для JWT
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Токен действителен 24 часа
		"iat":     time.Now().Unix(),                     // Время выпуска токена
	}

	// Создание токена с использованием алгоритма HMAC и ваших claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подпись токена с использованием секретного ключа
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func Login(ctx *gin.Context) {
	var loginData struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	db := ctx.MustGet("db").(*gorm.DB)
	var user models.User
	if err := db.Where("login = ?", loginData.Login).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Проверка хэшированного пароля
	log.Printf("Hashed password in DB: %s", user.Password)
	log.Printf("Input password: %s", loginData.Password)
	hashedinputpass, _ := bcrypt.GenerateFromPassword([]byte(loginData.Password), bcrypt.DefaultCost)
	log.Printf("generatefrompass: %s", string(hashedinputpass))

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"}) // Пароль не совпал
		return
	}

	// Генерация JWT токена
	token, err := GenerateJWT(user.UID.String()) // UID как строка для JWT
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
