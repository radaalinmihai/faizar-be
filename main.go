package main

import (
	"be/database"
	"be/database/entities"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func main() {
	database.InitDatabase()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("yokoso"))
	})
	r.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		db := database.GetDBConnection()
		username := "shackicko"
		user := entities.User{Name: "Alin", Username: &username, ID: 1}

		ctx := context.Background()
		result := gorm.WithResult()
		err := gorm.G[entities.User](db, result).Create(ctx, &user)

		if err != nil {
			panic(err)
		}

		data, err := json.Marshal(user)

		if err != nil {
			panic(err)
		}

		w.Write(data)
	})

	http.ListenAndServe(":3000", r)
}
