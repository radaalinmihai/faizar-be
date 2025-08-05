package auth

import (
	"be/database/entities"
	"context"
	"encoding/json"
	"net/http"

	"be/jwt"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var userDto RegisterUserDTO

	err := json.NewDecoder(r.Body).Decode(&userDto)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	db, _ := r.Context().Value("DB").(*gorm.DB)

	password, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)

	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user := entities.User{
		Name:     &userDto.Name,
		Email:    userDto.Email,
		Username: userDto.Username,
		Password: string(password),
	}
	err = gorm.G[entities.User](db).Create(ctx, &user)

	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	accessTokenInterface := GenerateAccessToken(user)

	_, accessToken, err := jwt.TokenAuth.Encode(accessTokenInterface)

	response := RegisterResponse{Message: "User registered successfully", AccessToken: accessToken}
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	w.Write(jsonData)
}
