package handlers

import (
	"net/http"
	"tiket-bioskop-mkp/config"
	"tiket-bioskop-mkp/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TheaterHandler struct {
	DB *gorm.DB
}

func NewTheaterHandler() *TheaterHandler {
	return &TheaterHandler{DB: config.DB}
}

func (h *TheaterHandler) GetAllTheaters(c *gin.Context) {
	var theaters []models.Theaters

	if err := config.DB.Find(&theaters).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to fetch theaters",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": theaters,
	})
}

func (h *TheaterHandler) GetTheaterById(c *gin.Context) {
	var theater models.Theaters
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).Find(&theater).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to fetch theaters",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": theater,
	})
}

func (h *TheaterHandler) CreateTheater(c *gin.Context) {
	var input models.Theaters
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Check duplicate title
	var existing models.Theaters
	if err := config.DB.Where("name = ?", input.Name).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Theater with this name already exists",
		})
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create add theater",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Add Theater success",
		"data":    input,
	})
}

func (h *TheaterHandler) UpdateTheater(c *gin.Context) {
	id := c.Param("id")

	var input models.Theaters
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var theater models.Theaters
	if err := config.DB.First(&theater, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Theater not found"})
		return
	}

	if err := config.DB.Model(&theater).Updates(input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update theater"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update theater success",
		"data":    theater,
	})
}

func (h *TheaterHandler) DeleteTheater(c *gin.Context) {
	id := c.Param("id")

	var theater models.Theaters
	if err := config.DB.First(&theater, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Theater not found"})
		return
	}

	if err := config.DB.Delete(&theater).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete theater"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete theater success",
	})
}
