package auth

import (
	"fmt"
	"instashop/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type Claims struct {
	Email    string `json:"email"`
	Role     string `json:"role"`
	UserID   uint   `json:"user_id"`
	UserName string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateToken(user *models.User) (string, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file. Error: %v", err)
	}

	JwtSecret := []byte(os.Getenv("JwtSecret"))

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		Email:    user.Email,
		Role:     user.Role,
		UserID:   user.ID,
		UserName: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "Instashop",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		log.Fatal("Error in creating token")
	}

	return tokenString, nil
}

// Middleware to validate token
func ValidateToken() gin.HandlerFunc {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file. Error: %v", err)
	}

	JwtSecret := []byte(os.Getenv("JwtSecret"))

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			c.Abort()
			return
		}

		const bearerPrefix = "Bearer "
		if len(authHeader) <= len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		tokenStr := authHeader[len(bearerPrefix):]

		token, err := jwt.ParseWithClaims(tokenStr, &Claims{},
			func(token *jwt.Token) (interface{}, error) {
				// Validate signing method
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return JwtSecret, nil
			})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Set("user_email", claims.Email)
		c.Set("user_role", claims.Role)
		c.Set("user_id", claims.UserID)
		c.Set("user_name", claims.UserName)

		c.Next()
	}
}

// Middleware to protect admin routes
func AdminOnly(c *gin.Context) {
	role := c.GetString("user_role")

	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Contact Admin staff to perform this operation."})
		c.Abort()
		return
	}

	c.Next()
}
