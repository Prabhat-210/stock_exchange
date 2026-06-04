package requestDTO

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignUpRequest struct {
	Email	string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	UserName string `json:"username" validate:"required"`
}