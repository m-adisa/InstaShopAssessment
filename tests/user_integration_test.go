package tests

import (
	"bytes"
	"fmt"
	"instashop/config"
	"instashop/models"
	"instashop/routes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Setup router
func SetupRouter() *gin.Engine {
	r := gin.Default()
	routes.UserRoutes(r)
	return r
}

// Setup test database
func SetupTestDatabase() {
	config.ConnectDatabase()

	// Migrate the database
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database. Error: %v", err)
	}

	fmt.Println("Database migrated successfully!")
}

// TestCreateUser tests the create user route
func TestCreateUser(t *testing.T) {
	SetupTestDatabase()
	router := SetupRouter()

	// Prepare the request payload
	payload := []byte(`{"name": "John Doe", "email": "jVb0W@example.com", "password": "password123", "role": "regular"}`)

	// Create a request
	req, _ := http.NewRequest("POST", "/users/register", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// Check the response
	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, "User created successfully", recorder.Body.String())

	// Check if the user was created in the database
	var user models.User
	result := config.DB.Where("email = ?", "jVb0W@example.com").First(&user)
	assert.Equal(t, nil, result.Error)
}
