package handler_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/sakaguchi-0725/echo-onion-arch/application/dto"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/apperr"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	mocks "github.com/sakaguchi-0725/echo-onion-arch/mocks/application/usecase"
	"github.com/sakaguchi-0725/echo-onion-arch/pkg/auth"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/handler"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/router"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupProfileHandler(t *testing.T) (*mocks.MockProfileUsecase, *echo.Echo) {
	ctrl := gomock.NewController(t)
	profileUsecase := mocks.NewMockProfileUsecase(ctrl)
	profileHandler := handler.NewProfileHandler(profileUsecase)

	e := echo.New()
	deps := NewMockHandlersDependency(ctrl, SetProfileHandler(profileHandler))
	router.NewRouter(e, deps, cfg)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	return profileUsecase, e
}

func TestProfileHandler_GetProfile_Success(t *testing.T) {
	profileUsecase, e := setupProfileHandler(t)
	userID := model.GenerateNewUserID()
	usecaseOutput := dto.ProfileOutput{
		Name: "John",
		Role: "admin",
	}

	profileUsecase.EXPECT().FindByUserID(userID.String()).Return(usecaseOutput, nil)

	token, err := auth.GenerateToken(userID, cfg.JWTSecret)
	require.NoError(t, err)

	rec, req := SetupRequest(e, http.MethodGet, "/profile", "")
	req.AddCookie(&http.Cookie{
		Name:  "access_token",
		Value: token,
	})
	e.ServeHTTP(rec, req)

	expectedRes := `{
		"name": "John",
		"role": "admin"
	}`

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, expectedRes, rec.Body.String())
}

func TestProfileHandler_GetProfile_InvalidUserID(t *testing.T) {
	profileUsecase, e := setupProfileHandler(t)
	userID := model.UserID("invalid_id")

	profileUsecase.EXPECT().FindByUserID(userID.String()).Return(dto.ProfileOutput{}, apperr.NewApplicationError(apperr.ErrBadReqeust, "invalid request", errors.New("error")))

	token, err := auth.GenerateToken(userID, cfg.JWTSecret)
	require.NoError(t, err)

	rec, req := SetupRequest(e, http.MethodGet, "/profile", "")
	req.AddCookie(&http.Cookie{
		Name:  "access_token",
		Value: token,
	})
	e.ServeHTTP(rec, req)

	expectedRes := `{
		"code": "BadRequest",
		"message": "invalid request"
	}`

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.JSONEq(t, expectedRes, rec.Body.String())
}
