package usecase_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sakaguchi-0725/echo-onion-arch/application/dto"
	"github.com/sakaguchi-0725/echo-onion-arch/application/usecase"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/apperr"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	mocks "github.com/sakaguchi-0725/echo-onion-arch/mocks/domain/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupProfileUsecase(t *testing.T) (usecase.ProfileUsecase, *mocks.MockProfileRepository) {
	ctrl := gomock.NewController(t)

	profileRepo := mocks.NewMockProfileRepository(ctrl)
	profileUsecase := usecase.NewProfileUsecase(profileRepo)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	return profileUsecase, profileRepo
}

func TestProfileUsecase_FindByUserID_Success(t *testing.T) {
	profileUsecase, profileRepo := setupProfileUsecase(t)

	userID := model.GenerateNewUserID()
	profile := model.Profile{
		UserID: userID,
		Name:   "John",
		Role:   model.General,
	}

	profileRepo.EXPECT().FindByID(userID).Return(profile, nil)

	output, err := profileUsecase.FindByUserID(userID.String())

	require.NoError(t, err)
	assert.Equal(t, output.Name, profile.Name)
	assert.Equal(t, output.Role, profile.Role.String())
}

func TestProfileUsecase_FindByUserID_InvalidID(t *testing.T) {
	profileUsecase, _ := setupProfileUsecase(t)

	id := "invalid_id"

	output, err := profileUsecase.FindByUserID(id)

	assert.Error(t, err)
	assert.EqualError(t, err, "invalid user ID format")
	assert.Equal(t, dto.ProfileOutput{}, output)
}

func TestProfileUsecase_FindByUserID_NotFound(t *testing.T) {
	profileUsecase, profileRepo := setupProfileUsecase(t)

	id := model.GenerateNewUserID()

	profileRepo.EXPECT().FindByID(id).Return(model.Profile{}, apperr.NewApplicationError(apperr.ErrNotFound, "Profile not found", fmt.Errorf("profile with ID %s not found", id)))

	output, err := profileUsecase.FindByUserID(id.String())
	assert.Error(t, err)
	assert.Equal(t, dto.ProfileOutput{}, output)

	appErr, ok := err.(*apperr.ApplicationError)
	assert.True(t, ok)
	assert.Equal(t, apperr.ErrNotFound, appErr.Code)
	assert.Equal(t, "Profile not found", appErr.Message)
}
