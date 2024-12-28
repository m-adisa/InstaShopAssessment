package controllers

import (
	"instashop/auth"
	"instashop/config"
	"instashop/models"
	"instashop/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignUp

// SignUp godoc
// @Summary Sign up a new user
// @Description Sign up a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param input body models.User true "User details"
// @Success 201 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 409 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /users/register [post]
func SignUp(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.APIResponse{Error: err.Error()})
		return
	}

	// Check if the user already exists
	var existingUser models.User
	if err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, utils.APIResponse{Error: "User already exists"})
		return
	}

	// Validate the input
	if err := utils.Validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, utils.APIResponse{Error: err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.APIResponse{Error: err.Error()})
		return
	}
	user.Password = hashedPassword

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.APIResponse{Error: err.Error()})
		return
	}

	// Generate JWT token
	token, err := auth.GenerateToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.APIResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, utils.APIResponse{
		Message: "User created successfully",
		Data: gin.H{
			"email": user.Email,
			"token": token,
		},
	})
}

// LoginUser

// LoginUser godoc
// @Summary Login a user
// @Description Login a user
// @Tags Users
// @Accept json
// @Produce json
// @Param input body models.LoginInput true "Login details"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /users/login [post]
func LoginUser(c *gin.Context) {
	var credentials models.LoginInput

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, utils.APIResponse{Error: err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", credentials.Email).First((&user)).Error; err != nil {
		c.JSON(http.StatusBadRequest, utils.APIResponse{Error: "Invalid email or password"})
		return
	}

	// Compare Password
	if !auth.ComparePassword(credentials.Password, user.Password) {
		c.JSON(http.StatusBadRequest, utils.APIResponse{Error: "Invalid email or password"})
		return
	}

	token, err := auth.GenerateToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.APIResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.APIResponse{
		Message: "User logged in successfully",
		Data: gin.H{
			"email": user.Email,
			"token": token,
		},
	})
}
