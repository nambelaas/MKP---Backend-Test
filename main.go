package main

import (
	"tiket-bioskop-mkp/config"
	"tiket-bioskop-mkp/models"
	"tiket-bioskop-mkp/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main(){
	config.LoadConfig()
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Users{})
	config.DB.AutoMigrate(&models.Movies{})
	config.DB.AutoMigrate(&models.Theaters{})
	config.DB.AutoMigrate(&models.Showtimes{})
	config.DB.AutoMigrate(&models.Seats{})

	r := gin.Default()
	routes.InitRoutes(r)

	r.Run(":"+ viper.GetString("APP_PORT"))
}