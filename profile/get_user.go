package profile

import (
	"be/auth"
	"be/database/entities"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"gorm.io/gorm"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	userData := &auth.AccessToken{}

	err := userData.FromMap(claims)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := context.Background()
	db, _ := r.Context().Value("DB").(*gorm.DB)

	user, err := gorm.G[entities.User](db).Where(entities.User{ID: userData.ID}).First(ctx)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := ProfileResponseDto{
		Username:  user.Username,
		Email:     user.Email,
		ID:        user.ID,
		UpdatedAt: user.UpdatedAt.String(),
		CreatedAt: user.CreatedAt.String(),
	}

	responseJson, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJson)
}
