package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/manab-pr/nebulo/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID      string `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
	jwt.RegisteredClaims
}

const (
	TokenKey             = "user_id"
	TokenExpirationHours = 24
)

var cfg *config.Config

func loadConfig() *config.Config {
	if cfg == nil {
		cfg = config.LoadConfig()
	}
	return cfg
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(_ *jwt.Token) (interface{}, error) {
			return []byte(loadConfig().JWT.Secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Check if token is expired
		if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}

		// Set user information in context
		c.Set(TokenKey, claims.UserID)
		c.Set("phone_number", claims.PhoneNumber)
		c.Next()
	}
}

func GenerateToken(userID, phoneNumber string) (tokenString string, expiresAt int64, err error) {
	expirationTime := time.Now().Add(TokenExpirationHours * time.Hour)
	claims := &Claims{
		UserID:      userID,
		PhoneNumber: phoneNumber,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(loadConfig().JWT.Secret))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expirationTime.Unix(), nil
}

func GetUserIDFromContext(c *gin.Context) (string, bool) {
	userID, exists := c.Get(TokenKey)
	if !exists {
		return "", false
	}
	return userID.(string), true
}
