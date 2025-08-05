package auth

type RegisterUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

type RegisterResponse struct {
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
}
