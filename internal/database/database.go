package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDatabase() (*gorm.DB, error) {
	u := os.Getenv("DB_USERNAME")
	p := os.Getenv("DB_PWD")
	h := os.Getenv("DB_HOST")
	t := os.Getenv("DB_TABLE")
	port := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", h, port, u, t, p)
	fmt.Println(connectionString)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return db, err
	}

	if err := db.DB().Ping(); err != nil {
		return db, err
	}

	return db, nil
}
