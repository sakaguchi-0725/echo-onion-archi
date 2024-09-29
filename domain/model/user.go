package model

type User struct {
	ID       UserID
	Email    string
	Password string
}

func NewUser(id UserID, email, password string) User {
	return User{
		ID:       id,
		Email:    email,
		Password: password,
	}
}
