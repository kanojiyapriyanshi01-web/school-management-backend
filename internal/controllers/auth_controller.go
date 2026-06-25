package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"school_backend/internal/config"
	"school_backend/internal/models"
	"school_backend/internal/auth"
)

func Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ? AND role = ?", req.Email, req.Role).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := auth.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}

func Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
		Role     string `json:"role" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
		Status:   "active",
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

func GetMe(c *gin.Context) {
	userID := c.GetUint("user_id")
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func SeedAdmin(c *gin.Context) {
	users := []struct {
		Name     string
		Email    string
		Password string
		Role     string
	}{
		{"Admin User", "admin@school.com", "Admin@2025", "admin"},
		{"Ravi Sharma", "staff@school.com", "Staff@2025", "staff"},
		{"Rahul Kumar", "student@school.com", "Student@2025", "student"},
		{"Suresh Kumar", "parent@school.com", "Parent@2025", "parent"},
	}

	for _, u := range users {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
		config.DB.Model(&models.User{}).Where("email = ?", u.Email).Update("password", string(hashed))
		user := models.User{Name: u.Name, Email: u.Email, Password: string(hashed), Role: u.Role, Status: "active"}
		config.DB.FirstOrCreate(&user, models.User{Email: u.Email})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Users seeded with new passwords",
		"users": []gin.H{
			{"role": "admin", "email": "admin@school.com", "password": "Admin@2025"},
			{"role": "staff", "email": "staff@school.com", "password": "Staff@2025"},
			{"role": "student", "email": "student@school.com", "password": "Student@2025"},
			{"role": "parent", "email": "parent@school.com", "password": "Parent@2025"},
		},
	})
}