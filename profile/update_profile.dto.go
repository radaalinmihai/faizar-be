package profile

type UpdateProfileDTO struct {
	Email    string `json:"email" validate:"omitempty,email,max=255"`
	Username string `json:"username" validate:"omitempty,max=69"`
}

type ProfileErrorDTO struct {
	Code string `json:"code"`
}
