package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/handler"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/validator"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/middleware"
)

type HandlerDependencies struct {
	AuthHandler handler.AuthHandler
}

func NewRouter(e *echo.Echo, deps *HandlerDependencies) {
	e.Validator = validator.NewValidator()

	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"}, // デプロイする場合は、オリジンを指定する
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(echoMiddleware.CSRFWithConfig(echoMiddleware.CSRFConfig{
		TokenLookup:    "header:X-CSRF-Token",
		CookieName:     "csrf_token",
		CookiePath:     "/",
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode, // デプロイ時はhttp.SamesiteStrictMode
	}))
	e.Use(middleware.ErrorMiddleware())

	e.GET("/csrf-token", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "CSRF token set"})
	})
	e.POST("/signin", deps.AuthHandler.SignIn)
	e.POST("/signup", deps.AuthHandler.SignUpForGeneral)
	e.POST("/signup/admin", deps.AuthHandler.SignUpForAdmin)
}
