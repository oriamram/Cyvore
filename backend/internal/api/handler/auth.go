package handler

import (
	"net/http"
	"time"

	"backend/internal/api/response"
	"backend/internal/auth"
	"backend/internal/models"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

var userService *service.UserService

func init() {
	var err error
	userService, err = service.NewUserService()
	if err != nil {
		panic(err)
	}
}

// Register handles user registration
func Register(c *gin.Context) {
	var reg models.UserRegistration
	if err := c.ShouldBindJSON(&reg); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := userService.Register(reg); err != nil {
		if err == models.ErrUserAlreadyExists {
			response.Error(c, http.StatusConflict, "Username already exists")
			return
		}
		response.Error(c, http.StatusInternalServerError, "Failed to register user")
		return
	}

	response.Success(c, gin.H{"message": "User registered successfully"})
}

// SignIn handles user authentication
func SignIn(c *gin.Context) {
	var login models.UserLogin
	if err := c.ShouldBindJSON(&login); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	user, err := userService.Login(login)
	if err != nil {
		if err == models.ErrUserNotFound || err == models.ErrInvalidPassword {
			response.Error(c, http.StatusUnauthorized, "Invalid username or password")
			return
		}
		response.Error(c, http.StatusInternalServerError, "Failed to authenticate user")
		return
	}

	// Generate token pair
	tokens, err := auth.GenerateTokenPair(user.ID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate tokens")
		return
	}

	// Set refresh token in HTTP-only cookie
	c.SetCookie(
		"refresh_token",
		tokens.RefreshToken,
		int(7*24*time.Hour.Seconds()), // 7 days
		"/",
		"",
		true,  // secure
		true,  // httpOnly
	)

	response.Success(c, gin.H{
		"access_token": tokens.AccessToken,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

// SignOut handles user sign out
func SignOut(c *gin.Context) {
	// Clear refresh token cookie
	c.SetCookie(
		"refresh_token",
		"",
		-1,
		"/",
		"",
		true,  // secure
		true,  // httpOnly
	)

	response.Success(c, gin.H{"message": "Signed out successfully"})
}

// RefreshToken handles token refresh requests
func RefreshToken(c *gin.Context) {
	// Get refresh token from cookie
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Refresh token not found")
		return
	}

	// Validate refresh token
	claims, err := auth.ValidateToken(refreshToken)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Invalid refresh token")
		return
	}

	// Generate new token pair
	tokens, err := auth.GenerateTokenPair(claims.UserID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate tokens")
		return
	}

	// Set new refresh token in HTTP-only cookie
	c.SetCookie(
		"refresh_token",
		tokens.RefreshToken,
		int(7*24*time.Hour.Seconds()), // 7 days
		"/",
		"",
		true,  // secure
		true,  // httpOnly
	)

	response.Success(c, gin.H{
		"access_token": tokens.AccessToken,
	})
} 