package main

import (
	"be/database"
	"be/profile"
	"net/http"

	"be/auth"

	_ "be/jwt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	database.InitDatabase()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(database.SetDBMiddleware)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("yokoso"))
	})
	r.Route("/auth", auth.Routes)
	r.Route("/profile", profile.Routes)

	http.ListenAndServe(":3000", r)
}
