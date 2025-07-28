package database

import (
	"be/config"
	"be/database/entities"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() {
	db := GetDBConnection()

	err := db.AutoMigrate(&entities.User{})

	if err != nil {
		panic(err)
	}
}

func GetDBConnection() *gorm.DB {
	_config := config.LoadConfig()
	dsn := fmt.Sprint("host=", _config.Database.Host, " user=", _config.Database.User, " password=", _config.Database.Password, " dbname=", _config.Database.Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
