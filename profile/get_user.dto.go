package profile

type ProfileUserDto struct {
	ID       uint
	Name     *string
	Email    string
	Username string
}

type ProfileResponseDto struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	ID        uint   `json:"id"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}
