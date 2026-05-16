package inboundPort

type AuthService interface {
	Login(email, password string) (string, error)
}
