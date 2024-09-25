package repository_test

import (
	"testing"

	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	domain "github.com/sakaguchi-0725/echo-onion-arch/domain/repository"
	"github.com/sakaguchi-0725/echo-onion-arch/infra/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupProfileRepository() domain.ProfileRepository {
	cleanUpTables(testDB, "profiles")
	return repository.NewProfileRepository(testDB)
}

func TestProfileRepository_Insert_Success(t *testing.T) {
	profileRepo := setupProfileRepository()

	profileID := model.GenerateNewUserID()
	profile := model.Profile{
		UserID: profileID,
		Name:   "John Doe",
		Role:   "admin",
	}

	insertedProfile, err := profileRepo.Insert(profile)

	require.NoError(t, err)
	assert.Equal(t, profileID, insertedProfile.UserID)
	assert.Equal(t, profile.Name, insertedProfile.Name)
	assert.Equal(t, profile.Role, insertedProfile.Role)
}

func TestProfileRepository_FindByID_Success(t *testing.T) {
	profileRepo := setupProfileRepository()

	profileID := model.GenerateNewUserID()
	profile := model.Profile{
		UserID: profileID,
		Name:   "Jane Doe",
		Role:   "general",
	}

	_, _ = profileRepo.Insert(profile)

	foundProfile, err := profileRepo.FindByID(profileID)

	require.NoError(t, err)
	assert.Equal(t, profileID, foundProfile.UserID)
	assert.Equal(t, profile.Name, foundProfile.Name)
	assert.Equal(t, profile.Role, foundProfile.Role)
}

func TestProfileRepository_FindByID_NotFound(t *testing.T) {
	profileRepo := setupProfileRepository()

	nonExistentID := model.GenerateNewUserID()

	_, err := profileRepo.FindByID(nonExistentID)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestProfileRepository_FindAll(t *testing.T) {
	profileRepo := setupProfileRepository()

	profile1 := model.Profile{
		UserID: model.GenerateNewUserID(),
		Name:   "John Doe",
		Role:   "admin",
	}
	profile2 := model.Profile{
		UserID: model.GenerateNewUserID(),
		Name:   "Jane Doe",
		Role:   "general",
	}

	_, _ = profileRepo.Insert(profile1)
	_, _ = profileRepo.Insert(profile2)

	profiles, err := profileRepo.FindAll()

	require.NoError(t, err)
	assert.Len(t, profiles, 2)

	assert.Equal(t, profile1.UserID, profiles[0].UserID)
	assert.Equal(t, profile1.Name, profiles[0].Name)
	assert.Equal(t, profile1.Role, profiles[0].Role)

	assert.Equal(t, profile2.UserID, profiles[1].UserID)
	assert.Equal(t, profile2.Name, profiles[1].Name)
	assert.Equal(t, profile2.Role, profiles[1].Role)
}

func TestProfileRepository_DeleteByID_Success(t *testing.T) {
	profileRepo := setupProfileRepository()

	profileID := model.GenerateNewUserID()
	profile := model.Profile{
		UserID: profileID,
		Name:   "John Doe",
		Role:   "admin",
	}

	_, _ = profileRepo.Insert(profile)

	err := profileRepo.DeleteByID(profileID)

	require.NoError(t, err)

	_, err = profileRepo.FindByID(profileID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestProfileRepository_DeleteByID_NotFound(t *testing.T) {
	profileRepo := setupProfileRepository()

	nonExistentID := model.GenerateNewUserID()

	err := profileRepo.DeleteByID(nonExistentID)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}
