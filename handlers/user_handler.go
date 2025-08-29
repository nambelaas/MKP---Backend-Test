package handlers

import (
	"net/http"
	"os"
	"tiket-bioskop-mkp/config"
	"tiket-bioskop-mkp/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler() *UserHandler {
	return &UserHandler{DB: config.DB}
}

func (h *UserHandler) Register(c *gin.Context) {
	var input models.UserRegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := models.Users{
		FullName:     input.FullName,
		Email:        input.Email,
		PasswordHash: string(hashedPass),
		Role:         models.UserRole(input.Role),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create new user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var input models.UserLoginInput

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.Users
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Email not found",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Incorrect Password",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserId": user.ID,
		"Role":   user.Role,
		"exp":    time.Now().Add(time.Hour * 3).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	tokenString, _ := token.SignedString([]byte(secret))

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
