package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sakaguchi-0725/echo-onion-arch/application/usecase"
	"github.com/sakaguchi-0725/echo-onion-arch/config"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/apperr"
	"github.com/sakaguchi-0725/echo-onion-arch/pkg/auth"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/dto"
)

type AuthHandler interface {
	SignIn(c echo.Context) error
	SignUpForAdmin(c echo.Context) error
	SignUpForGeneral(c echo.Context) error
}

type authHandler struct {
	usecase usecase.AuthUsecase
	config  *config.AppConfig
}

func NewAuthHandler(usecase usecase.AuthUsecase, config *config.AppConfig) AuthHandler {
	return &authHandler{usecase, config}
}

func (a *authHandler) SignIn(c echo.Context) error {
	var req dto.SignInRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return apperr.NewApplicationError(apperr.ErrBadReqeust, "invalid input", err)
	}

	userID, err := a.usecase.SignIn(req.Email, req.Password)
	if err != nil {
		return err
	}

	token, err := auth.GenerateToken(userID)
	if err != nil {
		return err
	}

	a.setAuthCookie(c, token)

	return c.JSON(http.StatusOK, map[string]string{"message": "sign in successful"})
}

func (a *authHandler) SignUpForAdmin(c echo.Context) error {
	var req dto.SignUpRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return apperr.NewApplicationError(apperr.ErrBadReqeust, "invalid input", err)
	}

	userID, err := a.usecase.SignUpForAdmin(req.Name, req.Email, req.Password)
	if err != nil {
		return err
	}

	token, err := auth.GenerateToken(userID)
	if err != nil {
		return err
	}

	a.setAuthCookie(c, token)

	return c.JSON(http.StatusOK, map[string]string{"message": "sign up successful"})
}

func (a *authHandler) SignUpForGeneral(c echo.Context) error {
	var req dto.SignUpRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return apperr.NewApplicationError(apperr.ErrBadReqeust, "invalid input", err)
	}

	userID, err := a.usecase.SignUpForGeneral(req.Name, req.Email, req.Password)
	if err != nil {
		return err
	}

	token, err := auth.GenerateToken(userID)
	if err != nil {
		return err
	}

	a.setAuthCookie(c, token)

	return c.JSON(http.StatusOK, map[string]string{"message": "sign up successful"})
}

func (a *authHandler) setAuthCookie(c echo.Context, token string) {
	cookie := new(http.Cookie)
	cookie.Name = "access_token"
	cookie.Value = token
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(12 * time.Hour)
	cookie.HttpOnly = true
	cookie.Secure = false                   // デプロイする場合はtrue
	cookie.SameSite = http.SameSiteNoneMode // デプロイする場合はhttp.SameSiteStrictMode

	c.SetCookie(cookie)
}
