package repository

import (
	"errors"
	"fmt"

	domain "github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/repository"
	"github.com/sakaguchi-0725/echo-onion-arch/infra/repository/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

func (u *userRepository) FindByEmail(email string) (domain.UserID, error) {
	var model model.User

	err := u.db.Where("email = ?", email).First(&model).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.UserID(""), fmt.Errorf("user with email %s not found", email)
		}
		return domain.UserID(""), fmt.Errorf("failed to retrieve user by email: %w", err)
	}

	return domain.UserID(model.ID), nil
}

func (u *userRepository) Insert(id domain.UserID, email string, password string) (domain.UserID, error) {
	user := model.NewUser(id, email, password)

	if err := u.db.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.UserID(""), fmt.Errorf("user with email %s already exists: %w", email, err)
		}
		return domain.UserID(""), fmt.Errorf("failed to insert user: %w", err)
	}

	return id, nil
}
