package routes

import (
	"tiket-bioskop-mkp/handlers"
	"tiket-bioskop-mkp/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	userHandlers := handlers.NewUserHandler()

	userRoutes := r.Group("/api/users")
	{
		userRoutes.POST("/register", userHandlers.Register)
		userRoutes.POST("/login", userHandlers.Login)
	}

	showtimeHandlers := handlers.NewShowtimeHandler()
	showtimeRoutes := r.Group("/api/showtimes")
	{
		showtimeRoutes.GET("/", showtimeHandlers.GetAllShowtimes)
		showtimeRoutes.GET("/:id", showtimeHandlers.GetShowtimeById)
		showtimeRoutes.POST("/", middleware.CheckJwt(), middleware.AdminOnly(), showtimeHandlers.CreateShowtime)
		showtimeRoutes.PUT("/:id", middleware.CheckJwt(), middleware.AdminOnly(), showtimeHandlers.UpdateShowtime)
		showtimeRoutes.DELETE("/:id", middleware.CheckJwt(), middleware.AdminOnly(), showtimeHandlers.DeleteShowtime)
	}

	movieHandlers := handlers.NewMovieHandler()
	movieRoutes := r.Group("/api/movies")
	{
		movieRoutes.GET("/", movieHandlers.GetAllMovies)
		movieRoutes.GET("/:id", movieHandlers.GetMovieById)
		movieRoutes.POST("/", middleware.CheckJwt(), middleware.AdminOnly(), movieHandlers.CreateMovie)
		movieRoutes.PUT("/:id", middleware.CheckJwt(), middleware.AdminOnly(), movieHandlers.UpdateMovie)
		movieRoutes.DELETE("/:id", middleware.CheckJwt(), middleware.AdminOnly(), movieHandlers.DeleteMovie)
	}

	theaterHandlers := handlers.NewTheaterHandler()
	theaterRoutes := r.Group("/api/theaters")
	{
		theaterRoutes.GET("/", theaterHandlers.GetAllTheaters)
		theaterRoutes.GET("/:id", theaterHandlers.GetTheaterById)
		theaterRoutes.POST("/", middleware.CheckJwt(), middleware.AdminOnly(), theaterHandlers.CreateTheater)
		theaterRoutes.PUT("/:id", middleware.CheckJwt(), middleware.AdminOnly(), theaterHandlers.UpdateTheater)
		theaterRoutes.DELETE("/:id", middleware.CheckJwt(), middleware.AdminOnly(), theaterHandlers.DeleteTheater)
	}
}
