package auth

type LoginUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
