package postgres

import (
	"backend-challenge-api/internal/domain/entities"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	DB *gorm.DB
}

var DB DBInstance

func Connect() *gorm.DB {
	dsn := fmt.Sprintf(`host=db user=%s password=%s dbname=%s port=5432 sslmode=disable`,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running models migrations...")
	db.AutoMigrate(&entities.Expression{})

	DB = DBInstance{
		DB: db,
	}

	return db
}
