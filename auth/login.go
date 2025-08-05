package auth

import (
	"be/database/entities"
	"context"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var userPayload LoginUserDTO
	err := json.NewDecoder(r.Body).Decode(&userPayload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	db, _ := r.Context().Value("DB").(*gorm.DB)

	data, err := gorm.G[entities.User](db).Where(&entities.User{Username: userPayload.Username}).First(ctx)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write(jsonData)
}
