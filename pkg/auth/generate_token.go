package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
)

type JwtClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID model.UserID, secretKey string) (string, error) {
	claims := &JwtClaims{
		UserID: userID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}
