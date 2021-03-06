package input

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUsernameRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
