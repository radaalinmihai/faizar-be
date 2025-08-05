package auth

import (
	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.Post("/login", LoginHandler)
	r.Post("/register", RegisterHandler)
}
