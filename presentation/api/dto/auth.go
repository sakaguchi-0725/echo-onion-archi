package dto

import "github.com/golang-jwt/jwt/v5"

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignUpRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type JwtClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}
