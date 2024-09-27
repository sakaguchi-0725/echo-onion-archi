package persistence

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/apperr"
	domain "github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/repository"
	"github.com/sakaguchi-0725/echo-onion-arch/infra/persistence/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

func (u *userRepository) FindByEmail(email string) (domain.User, error) {
	var model model.User

	err := u.db.Where("email = ?", email).First(&model).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, apperr.NewApplicationError(apperr.ErrNotFound, fmt.Sprintf("User with email %s not found", email), err)
		}
		return domain.User{}, apperr.NewApplicationError(apperr.ErrInternalError, "Failed to retrieve user by email", err)
	}

	user := domain.NewUser(domain.UserID(model.ID), model.Email, model.Password)

	return user, nil
}

func (u *userRepository) Insert(id domain.UserID, email string, password string) (domain.UserID, error) {
	user := model.NewUser(id, email, password)

	if err := u.db.Create(&user).Error; err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return domain.UserID(""), apperr.NewApplicationError(apperr.ErrBadReqeust, "This email address cannot be used", err)
		}
		return domain.UserID(""), apperr.NewApplicationError(apperr.ErrInternalError, "Failed to insert user", err)
	}

	return id, nil
}
