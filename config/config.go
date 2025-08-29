package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig(){
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	if err := viper.ReadInConfig(); err!=nil{
		log.Fatal("Error reading config file, ", err)
	}
}