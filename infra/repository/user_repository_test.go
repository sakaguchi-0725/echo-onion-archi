package repository_test

import (
	"testing"

	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	domain "github.com/sakaguchi-0725/echo-onion-arch/domain/repository"
	"github.com/sakaguchi-0725/echo-onion-arch/infra/repository"
	"github.com/stretchr/testify/assert"
)

func setupUserRepositoryTest() domain.UserRepository {
	cleanUpTables(testDB, "users")
	return repository.NewUserRepository(testDB)
}

func TestUserRepository_Insert_Success(t *testing.T) {
	userRepo := setupUserRepositoryTest()

	userID := model.GenerateNewUserID()
	email := "test@example.com"
	password := "securepassword"

	insertedID, err := userRepo.Insert(userID, email, password)

	assert.NoError(t, err)
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
	_, err := userRepo.Insert(userID, email, password)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "SQLSTATE 23505")
}

func TestUserRepository_FindByEmail_Success(t *testing.T) {
	userRepo := setupUserRepositoryTest()

	userID := model.GenerateNewUserID()
	email := "findme@example.com"
	password := "securepassword"

	// データを事前に挿入
	_, _ = userRepo.Insert(userID, email, password)

	foundID, err := userRepo.FindByEmail(email)

	assert.NoError(t, err)
	assert.Equal(t, userID, foundID)
}

func TestUserRepository_FindByEmail_NotFound(t *testing.T) {
	userRepo := setupUserRepositoryTest()

	_, err := userRepo.FindByEmail("notfound@example.com")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}
