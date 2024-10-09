package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("your_secret_key")

type Claims struct {
	UID string `json:"uid"`
	jwt.RegisteredClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			ctx.Abort()
			return
		}

		// Проверяем наличие префикса Bearer
		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token required"})
			ctx.Abort()
			return
		}

		// Получаем токен из заголовка
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Валидация токена
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		// Проверяем, нет ли ошибок при разборе токена и его валидности
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		// Извлекаем claims и проверяем истечение срока действия токена
		claims, ok := token.Claims.(*Claims)
		if !ok || claims == nil || claims.ExpiresAt == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token structure"})
			ctx.Abort()
			return
		}

		if claims.ExpiresAt.Time.Before(time.Now()) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			ctx.Abort()
			return
		}

		ctx.Set("userID", claims.UID) // Устанавливаем UID пользователя в контекст
		ctx.Next()
	}
}
