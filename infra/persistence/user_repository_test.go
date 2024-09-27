package persistence_test

import (
	"testing"

	"github.com/sakaguchi-0725/echo-onion-arch/domain/apperr"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/repository"
	"github.com/sakaguchi-0725/echo-onion-arch/infra/persistence"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupUserRepositoryTest() repository.UserRepository {
	cleanUpTables(testDB, "users")
	return persistence.NewUserRepository(testDB)
}

func TestUserRepository_Insert_Success(t *testing.T) {
	userRepo := setupUserRepositoryTest()

	userID := model.GenerateNewUserID()
	email := "test@example.com"
	password := "securepassword"

	insertedID, err := userRepo.Insert(userID, email, password)

	require.NoError(t, err)
	assert.Equal(t, userID, insertedID)
}

func TestUserRepository_Insert_DuplicateEmail(t *testing.T) {
	userRepo := setupUserRepositoryTest()

	userID := model.GenerateNewUserID()
	email := "test@example.com"
	password := "securepassword"

	// 1回目の挿入
	_, _ = userRepo.Insert(userID, email, password)

	// 2回目の挿入（重複エラー）
	userID2 := model.GenerateNewUserID()
	_, err := userRepo.Insert(userID2, email, password)

	require.Error(t, err)

	appErr, ok := err.(*apperr.ApplicationError)
	require.True(t, ok)
	assert.Equal(t, apperr.ErrBadReqeust, appErr.Code)
	assert.Equal(t, "This email address cannot be used", appErr.Message)
}

func TestUserRepository_FindByEmail_Success(t *testing.T) {
	userRepo := setupUserRepositoryTest()

	userID := model.GenerateNewUserID()
	email := "findme@example.com"
	password := "securepassword"

	// データを事前に挿入
	_, _ = userRepo.Insert(userID, email, password)

	foundUser, err := userRepo.FindByEmail(email)

	require.NoError(t, err)
	assert.Equal(t, userID, foundUser.ID)
	assert.Equal(t, email, foundUser.Email)
}

func TestUserRepository_FindByEmail_NotFound(t *testing.T) {
	userRepo := setupUserRepositoryTest()

	_, err := userRepo.FindByEmail("notfound@example.com")

	require.Error(t, err)

	appErr, ok := err.(*apperr.ApplicationError)
	require.True(t, ok)
	assert.Equal(t, apperr.ErrNotFound, appErr.Code)
	assert.Equal(t, "User with email notfound@example.com not found", appErr.Message)
}
