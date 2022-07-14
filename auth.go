package auth

type User struct {
	ID   string
	Name string
}

type Authenticator interface {
	Authenticate() error
}

type UserRepository interface {
	FindByID(id string) User
}
