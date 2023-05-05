package postgres

import (
	"backend-challenge-api/models"
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

	log.Println(dsn)
	log.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running models mogrations...")
	db.AutoMigrate(&models.Expression{})

	DB = DBInstance{
		DB: db,
	}

	return db
}

// func Connect() (*sql.DB, error) {
// 	db, err := sql.Open("postgres", "user=postgres dbname=expressions password=postgres host=0.0.0.0 port=5454 sslmode=disable")
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err = db.Ping(); err != nil {
// 		return nil, err
// 	}

// 	fmt.Println("Connected to PostgreSQL!")

// 	return db, nil
// }
