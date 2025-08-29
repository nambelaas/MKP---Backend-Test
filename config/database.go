package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	cred := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s search_path=tiket_bioskop",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(cred), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	sqlDB, _ := db.DB()
	_, err = sqlDB.Exec("CREATE SCHEMA IF NOT EXISTS tiket_bioskop")
	if err != nil {
		log.Fatal("Failed to create schema:", err)
	}

	sqlExt, _ := db.DB()
	_, err = sqlExt.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	if err != nil {
		log.Fatal("Failed to create extension:", err)
	}

	if err := db.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
				CREATE TYPE tiket_bioskop.user_role AS ENUM ('Admin', 'User');
			END IF;
		END$$;
	`).Error; err != nil {
		log.Fatal(err)
	}

	if err := db.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'seat_status') THEN
				CREATE TYPE tiket_bioskop.seat_status AS ENUM ('Available', 'Hold', 'Paid', 'Cancelled', 'Expired');
			END IF;
		END$$;
	`).Error; err != nil {
		log.Fatal(err)
	}

	if err := db.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'order_status') THEN
				CREATE TYPE tiket_bioskop.order_status AS ENUM ('Pending', 'Paid', 'Completed', 'Cancelled', 'Refunded');
			END IF;
		END$$;
	`).Error; err != nil {
		log.Fatal(err)
	}


	DB = db
	log.Println("Database connected")
}