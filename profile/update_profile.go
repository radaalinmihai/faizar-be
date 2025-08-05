package profile

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"be/auth"
	"be/database/entities"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func validateDTO(updateProfileDTO UpdateProfileDTO) ([]ProfileErrorDTO, bool) {
	err := validate.Struct(updateProfileDTO)

	if err != nil {
		var validateErrs validator.ValidationErrors

		if errors.As(err, &validateErrs) {
			var errorsResponse []ProfileErrorDTO
			for _, e := range validateErrs {
				errorType := strings.ToUpper(e.Tag())
				errorsResponse = append(errorsResponse, ProfileErrorDTO{Code: errorType})
			}
			return errorsResponse, false
		}
	}
	return []ProfileErrorDTO{}, false
}

func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	userData := &auth.AccessToken{}

	err := userData.FromMap(claims)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var updateProfileDTO UpdateProfileDTO
	err = json.NewDecoder(r.Body).Decode(&updateProfileDTO)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	errors, ok := validateDTO(updateProfileDTO)
	if !ok {
		errorsBody, _ := json.Marshal(errors)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, string(errorsBody), http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	db, _ := ctx.Value("DB").(*gorm.DB)

	gorm.G[entities.User](db).Where(entities.User{ID: userData.ID}).Updates(ctx, entities.User{
		Email:    updateProfileDTO.Email,
		Username: updateProfileDTO.Username,
	})

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
