package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/apperr"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func ErrorMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err == nil {
				return nil
			}

			log.Errorf("%v", err)

			if appErr, ok := err.(*apperr.ApplicationError); ok {
				switch appErr.Code {
				case apperr.ErrBadReqeust:
					return c.JSON(http.StatusBadRequest, &ErrorResponse{
						Code:    appErr.Code.String(),
						Message: appErr.Message,
					})
				case apperr.ErrNotFound:
					return c.JSON(http.StatusNotFound, &ErrorResponse{
						Code:    appErr.Code.String(),
						Message: appErr.Message,
					})
				case apperr.ErrUnauthorized:
					return c.JSON(http.StatusUnauthorized, &ErrorResponse{
						Code:    appErr.Code.String(),
						Message: appErr.Message,
					})
				default:
					return c.JSON(http.StatusInternalServerError, &ErrorResponse{
						Code:    appErr.Code.String(),
						Message: appErr.Message,
					})
				}
			}

			return c.NoContent(http.StatusInternalServerError)
		}
	}
}
