package handlers

import (
	"net/http"
	"tiket-bioskop-mkp/config"
	"tiket-bioskop-mkp/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MovieHandler struct {
	DB *gorm.DB
}

func NewMovieHandler() *MovieHandler {
	return &MovieHandler{DB: config.DB}
}

func (h *MovieHandler) GetAllMovies(c *gin.Context) {
	var movies []models.Movies

	if err := config.DB.Find(&movies).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to fetch movies",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": movies,
	})
}

func (h *MovieHandler) GetMovieById(c *gin.Context) {
	var movie models.Movies
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).Find(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to fetch movies",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": movie,
	})
}

func (h *MovieHandler) CreateMovie(c *gin.Context) {
	var input models.Movies
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Check duplicate title
	var existing models.Movies
	if err := config.DB.Where("title = ?", input.Title).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Movie with this title already exists",
		})
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create add movie",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Add Movie success",
		"data":    input,
	})
}

func (h *MovieHandler) UpdateMovie(c *gin.Context) {
	id := c.Param("id")

	var input models.Movies
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var movie models.Movies
	if err := config.DB.First(&movie, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	if err := config.DB.Model(&movie).Updates(input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update movie"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update movie success",
		"data":    movie,
	})
}

func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	id := c.Param("id")

	var movie models.Movies
	if err := config.DB.First(&movie, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	if err := config.DB.Delete(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete movie"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete movie success",
	})
}
