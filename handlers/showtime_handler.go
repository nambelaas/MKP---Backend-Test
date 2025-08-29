package handlers

import (
	"net/http"
	"tiket-bioskop-mkp/config"
	"tiket-bioskop-mkp/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShowtimeHandler struct {
	DB *gorm.DB
}

func NewShowtimeHandler() *ShowtimeHandler {
	return &ShowtimeHandler{DB: config.DB}
}

func (h *ShowtimeHandler) GetAllShowtimes(c *gin.Context) {
	var showtimes []models.Showtimes

	if err := config.DB.Preload("Movie").Preload("Theater").Find(&showtimes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to fetch showtimes",
		})
		return
	}

	var result []models.ListShowtimes

	for _, s := range showtimes {
		var dataSeat []models.ListSeats

		if err := config.DB.Model(&models.Seats{}).Where("showtime_id = ?", s.ID).Find(&dataSeat).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to fetch seats",
			})
			return
		}

		result = append(result, models.ListShowtimes{
			ID:          s.ID,
			MovieName:   s.Movie.Title,
			TheaterName: s.Theater.Name,
			StartAt:     s.StartAt,
			BasePrice:   s.BasePrice,
			Seats:       dataSeat,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (h *ShowtimeHandler) GetShowtimeById(c *gin.Context) {
	var showtimes models.Showtimes
	id := c.Param("id")

	if err := config.DB.Preload("Movie").Preload("Theater").Where("id = ?", id).Find(&showtimes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to fetch showtimes",
		})
		return
	}

	var result models.ListShowtimes

	var dataSeat []models.ListSeats

	if err := config.DB.Model(&models.Seats{}).Where("showtime_id = ?", showtimes.ID).Find(&dataSeat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to fetch seats",
		})
		return
	}

	result = models.ListShowtimes{
		ID:          showtimes.ID,
		MovieName:   showtimes.Movie.Title,
		TheaterName: showtimes.Theater.Name,
		StartAt:     showtimes.StartAt,
		BasePrice:   showtimes.BasePrice,
		Seats:       dataSeat,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (h *ShowtimeHandler) CreateShowtime(c *gin.Context) {
	var input models.ShowtimeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	movieData := config.DB.Where("id = ?", input.MovieId).First(&models.Movies{})
	if movieData.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Movie not found",
		})
		return
	}

	theaterData := config.DB.Where("id = ?", input.TheaterId).First(&models.Theaters{})
	if theaterData.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Theater not found",
		})
		return
	}

	startAt, err := time.Parse("2006-01-02 15:04:05", input.StartAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid time format. Use YYYY-MM-DD HH:MM:SS",
		})
		return
	}

	dataShowtime := models.Showtimes{
		MovieId:   input.MovieId,
		TheaterId: input.TheaterId,
		StartAt:   startAt,
		BasePrice: input.BasePrice,
	}

	if err := config.DB.Create(&dataShowtime).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create new showtime",
		})
		return
	}

	for _, seat := range input.Seats {
		if err := config.DB.Create(&models.Seats{
			ShowtimeId: dataShowtime.ID,
			SeatCode:   seat.SeatCode,
			Status:     seat.Status,
		}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to create seats for the showtime",
			})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Showtime created successfully",
	})
}

func (h *ShowtimeHandler) UpdateShowtime(c *gin.Context) {
	id := c.Param("id")
	var input models.ShowtimeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var existing models.Showtimes
	if err := config.DB.First(&existing, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Showtime not found"})
		return
	}

	movieData := config.DB.Where("id = ?", input.MovieId).First(&models.Movies{})
	if movieData.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Movie not found",
		})
		return
	}

	theaterData := config.DB.Where("id = ?", input.TheaterId).First(&models.Theaters{})
	if theaterData.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Theater not found",
		})
		return
	}

	startAt, err := time.Parse("2006-01-02 15:04:05", input.StartAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid time format. Use YYYY-MM-DD HH:MM:SS",
		})
		return
	}

	dataShowtime := models.Showtimes{
		MovieId:   input.MovieId,
		TheaterId: input.TheaterId,
		StartAt:   startAt,
		BasePrice: input.BasePrice,
	}

	if err := config.DB.Model(&existing).Updates(dataShowtime).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update showtime"})
		return
	}

	for _, seat := range input.Seats {
		if err := config.DB.Model(&models.Seats{}).Where("id = ? AND showtime_id = ?", seat.ID, existing.ID).Updates(models.Seats{
			SeatCode: seat.SeatCode,
			Status:   seat.Status,
		}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to update seats for the showtime",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update showtime success",
		"data":    existing,
	})
}

func (h *ShowtimeHandler) DeleteShowtime(c *gin.Context) {
	id := c.Param("id")

	var showtime models.Showtimes
	if err := config.DB.First(&showtime, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Showtime not found"})
		return
	}

	if err := config.DB.Delete(&showtime).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete showtime"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete showtime success",
	})
}
