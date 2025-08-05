package database

import (
	"be/config"
	"be/database/entities"
	"context"
	"fmt"
	"net/http"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() {
	db := GetDBConnection()

	err := db.Debug().AutoMigrate(&entities.User{})

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

func SetDBMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := GetDBConnection()
		timeoutContext, _ := context.WithTimeout(context.Background(), time.Second)
		ctx := context.WithValue(r.Context(), "DB", db.WithContext(timeoutContext))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
