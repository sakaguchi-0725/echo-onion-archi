package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/apperr"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	mocks "github.com/sakaguchi-0725/echo-onion-arch/mocks/application/usecase"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/dto"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/handler"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/router"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupAuthHandler(t *testing.T) (*mocks.MockAuthUsecase, *echo.Echo) {
	ctrl := gomock.NewController(t)
	authUsecase := mocks.NewMockAuthUsecase(ctrl)
	authHandler := handler.NewAuthHandler(authUsecase, &cfg)

	e := echo.New()
	deps := NewMockHandlersDependency(ctrl, SetAuthHandler(authHandler))
	router.NewRouter(e, deps, cfg)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	return authUsecase, e
}

func TestAuthHandler_SignIn_Success(t *testing.T) {
	authUsecase, e := setupAuthHandler(t)

	email := "test@example.com"
	password := "password"
	expectedUserID := model.GenerateNewUserID()

	signInRequest := dto.SignInRequest{
		Email:    email,
		Password: password,
	}

	reqBody, err := json.Marshal(signInRequest)
	require.NoError(t, err)

	rec, req := SetupRequest(e, http.MethodPost, "/signin", string(reqBody))

	authUsecase.EXPECT().SignIn(email, password).Return(expectedUserID, nil)

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"message": "sign in successful"}`, rec.Body.String())
}

func TestAuthHandler_SignIn_BadRequest(t *testing.T) {
	_, e := setupAuthHandler(t)

	email := "test"
	password := "password"

	signInRequest := dto.SignInRequest{
		Email:    email,
		Password: password,
	}

	reqBody, err := json.Marshal(signInRequest)
	require.NoError(t, err)

	rec, req := SetupRequest(e, http.MethodPost, "/signin", string(reqBody))

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.JSONEq(t, `
	{
		"code": "BadRequest",
		"message": "invalid input"
	}`, rec.Body.String())
}

func TestAuthHandler_SignIn_Unauhorized(t *testing.T) {
	authUsecase, e := setupAuthHandler(t)

	email := "test@example.com"
	password := "password"

	signInReqest := dto.SignInRequest{
		Email:    email,
		Password: password,
	}

	reqBody, err := json.Marshal(signInReqest)
	require.NoError(t, err)

	rec, req := SetupRequest(e, http.MethodPost, "/signin", string(reqBody))

	authUsecase.EXPECT().SignIn(email, password).Return(model.UserID(""), apperr.NewApplicationError(apperr.ErrUnauthorized, "Authentication failed. Please check your email and password.", errors.New("error")))
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.JSONEq(t, `
	{
		"code": "Unauthorized",
		"message": "Authentication failed. Please check your email and password."
	}`, rec.Body.String())
}

func TestAuthHandler_SignUpForAdmin_Success(t *testing.T) {
	authUsecase, e := setupAuthHandler(t)

	expectedID := model.GenerateNewUserID()
	name := "John"
	email := "test@example.com"
	password := "password"

	signUpRequest := dto.SignUpRequest{
		Name:     name,
		Email:    email,
		Password: password,
	}

	reqBody, err := json.Marshal(signUpRequest)
	require.NoError(t, err)

	rec, req := SetupRequest(e, http.MethodPost, "/signup/admin", string(reqBody))

	authUsecase.EXPECT().SignUpForAdmin(name, email, password).Return(expectedID, nil)

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"message": "sign up successful"}`, rec.Body.String())
}

func TestAuthHandler_SignUpForAdmin_InvalidInput(t *testing.T) {
	_, e := setupAuthHandler(t)

	name := ""
	email := "test@example.com"
	password := "password"

	signUpRequest := dto.SignUpRequest{
		Name:     name,
		Email:    email,
		Password: password,
	}

	reqBody, err := json.Marshal(signUpRequest)
	require.NoError(t, err)

	rec, req := SetupRequest(e, http.MethodPost, "/signup/admin", string(reqBody))
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.JSONEq(t, `
	{
		"code": "BadRequest",
		"message": "invalid input"
	}`, rec.Body.String())
}

func TestAuthHandler_SignUpForAdmin_BadRequest(t *testing.T) {
	authUsecase, e := setupAuthHandler(t)

	name := "John"
	email := "test@example.com"
	password := "password"

	signUpRequest := dto.SignUpRequest{
		Name:     name,
		Email:    email,
		Password: password,
	}

	reqBody, err := json.Marshal(signUpRequest)
	require.NoError(t, err)

	rec, req := SetupRequest(e, http.MethodPost, "/signup/admin", string(reqBody))

	authUsecase.EXPECT().SignUpForAdmin(name, email, password).Return(model.UserID(""), apperr.NewApplicationError(apperr.ErrBadReqeust, "This email address cannot be used", errors.New("error")))

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.JSONEq(t, `
	{
		"code": "BadRequest",
		"message": "This email address cannot be used"
	}`, rec.Body.String())
}

func TestAuthHandler_SignUpForGeneral_Success(t *testing.T) {
	authUsecase, e := setupAuthHandler(t)

	expectedID := model.GenerateNewUserID()
	name := "John"
	email := "test@example.com"
	password := "password"

	signUpRequest := dto.SignUpRequest{
		Name:     name,
		Email:    email,
		Password: password,
	}

	reqBody, err := json.Marshal(signUpRequest)
	require.NoError(t, err)

	rec, req := SetupRequest(e, http.MethodPost, "/signup", string(reqBody))

	authUsecase.EXPECT().SignUpForGeneral(name, email, password).Return(expectedID, nil)

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"message": "sign up successful"}`, rec.Body.String())
}

func TestAuthHandler_SignUpForGeneral_InvalidInput(t *testing.T) {
	_, e := setupAuthHandler(t)

	name := ""
	email := "test@example.com"
	password := "password"

	signUpRequest := dto.SignUpRequest{
		Name:     name,
		Email:    email,
		Password: password,
	}

	reqBody, err := json.Marshal(signUpRequest)
	require.NoError(t, err)

	rec, req := SetupRequest(e, http.MethodPost, "/signup", string(reqBody))

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.JSONEq(t, `
	{
		"code": "BadRequest",
		"message": "invalid input"
	}`, rec.Body.String())
}

func TestAuthHandler_SignUpForGeneral_BadRequest(t *testing.T) {
	authUsecase, e := setupAuthHandler(t)

	name := "John"
	email := "test@example.com"
	password := "password"

	signUpRequest := dto.SignUpRequest{
		Name:     name,
		Email:    email,
		Password: password,
	}

	reqBody, err := json.Marshal(signUpRequest)
	require.NoError(t, err)

	rec, req := SetupRequest(e, http.MethodPost, "/signup", string(reqBody))

	authUsecase.EXPECT().SignUpForGeneral(name, email, password).Return(model.UserID(""), apperr.NewApplicationError(apperr.ErrBadReqeust, "This email address cannot be used", errors.New("error")))

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.JSONEq(t, `
	{
		"code": "BadRequest",
		"message": "This email address cannot be used"
	}`, rec.Body.String())
}
