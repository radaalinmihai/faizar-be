package profile

import (
	"be/jwt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func Routes(r chi.Router) {
	r.Use(jwtauth.Verifier(jwt.TokenAuth))
	r.Use(jwtauth.Authenticator(jwt.TokenAuth))

	r.Get("/", GetUser)
	r.Patch("/", UpdateProfileHandler)
}
