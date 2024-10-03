package handler_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	mocks "github.com/sakaguchi-0725/echo-onion-arch/mocks/application/usecase"
	"github.com/sakaguchi-0725/echo-onion-arch/pkg/config"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/handler"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/router"
)

var cfg config.AppConfig

func TestMain(m *testing.M) {
	cfg.JWTSecret = "test_secret"

	code := m.Run()
	os.Exit(code)
}

func SetupRequest(e *echo.Echo, method, path, requestBody string) (*httptest.ResponseRecorder, *http.Request) {
	csrfReq := httptest.NewRequest(http.MethodGet, "/csrf-token", nil)
	csrfRec := httptest.NewRecorder()
	e.ServeHTTP(csrfRec, csrfReq)

	csrfToken := csrfRec.Result().Cookies()[0].Value

	req := httptest.NewRequest(method, path, strings.NewReader(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("X-CSRF-Token", csrfToken)
	req.AddCookie(&http.Cookie{Name: "csrf_token", Value: csrfToken})

	rec := httptest.NewRecorder()

	return rec, req
}

type MockHandlersDependencyOptionFunc func(*router.HandlerDependencies)

func NewMockHandlersDependency(ctrl *gomock.Controller, options ...MockHandlersDependencyOptionFunc) *router.HandlerDependencies {
	hd := &router.HandlerDependencies{
		AuthHandler:    handler.NewAuthHandler(mocks.NewMockAuthUsecase(ctrl), &config.AppConfig{}),
		ProfileHandler: handler.NewProfileHandler(mocks.NewMockProfileUsecase(ctrl)),
	}

	for _, option := range options {
		option(hd)
	}
	return hd
}

func SetAuthHandler(h handler.AuthHandler) MockHandlersDependencyOptionFunc {
	return func(hd *router.HandlerDependencies) {
		hd.AuthHandler = h
	}
}

func SetProfileHandler(h handler.ProfileHandler) MockHandlersDependencyOptionFunc {
	return func(hd *router.HandlerDependencies) {
		hd.ProfileHandler = h
	}
}
