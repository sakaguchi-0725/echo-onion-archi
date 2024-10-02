package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/apperr"
	"github.com/sakaguchi-0725/echo-onion-arch/pkg/config"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("access_token")
		if err != nil {
			return apperr.NewApplicationError(apperr.ErrUnauthorized, "access token not found", err)
		}

		tokenString := cookie.Value
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, apperr.NewApplicationError(apperr.ErrUnauthorized, "", err)
			}
			return []byte(config.GetJWTSecret()), nil
		})

		if err != nil || !token.Valid {
			return apperr.NewApplicationError(apperr.ErrUnauthorized, "invalid access token", err)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return apperr.NewApplicationError(apperr.ErrUnauthorized, "invalid access token", nil)
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			return apperr.NewApplicationError(apperr.ErrUnauthorized, "invalid access token", nil)
		}

		c.Set("userID", userID)
		return next(c)
	}
}
