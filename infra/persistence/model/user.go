package model

import (
	"time"

	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
)

type User struct {
	ID        string    `gorm:"primaryKey;not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func NewUser(id model.UserID, email, password string) User {
	return User{
		ID:       id.String(),
		Email:    email,
		Password: password,
	}
}
