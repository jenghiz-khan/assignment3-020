package database

import (
	"assignment-3/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host		= "localhost"
	port		= 5432
	user		= "postgres"
	password	= "postgres"
	dbname		= "postgres"
)

func InitDB() (*gorm.DB, error) {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    
	db.Debug().AutoMigrate(models.Status{})
    
    return db, nil
}