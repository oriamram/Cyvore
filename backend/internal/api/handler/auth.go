package handler

import (
	"net/http"
	"time"

	"backend/internal/auth"
	"backend/internal/model"
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
	var reg model.UserRegistration
	if err := c.ShouldBindJSON(&reg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := userService.Register(reg); err != nil {
		if err == model.ErrUserAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// SignIn handles user authentication
func SignIn(c *gin.Context) {
	var login model.UserLogin
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := userService.Login(login)
	if err != nil {
		if err == model.ErrUserNotFound || err == model.ErrInvalidPassword {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to authenticate user"})
		return
	}

	// Generate token pair
	tokens, err := auth.GenerateTokenPair(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
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

	c.JSON(http.StatusOK, gin.H{
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

	c.JSON(http.StatusOK, gin.H{"message": "Signed out successfully"})
}

// RefreshToken handles token refresh requests
func RefreshToken(c *gin.Context) {
	// Get refresh token from cookie
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
		return
	}

	// Validate refresh token
	claims, err := auth.ValidateToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Generate new token pair
	tokens, err := auth.GenerateTokenPair(claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
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

	c.JSON(http.StatusOK, gin.H{
		"access_token": tokens.AccessToken,
	})
} 