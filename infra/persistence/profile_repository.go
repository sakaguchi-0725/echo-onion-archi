package persistence

import (
	"errors"
	"fmt"

	"github.com/sakaguchi-0725/echo-onion-arch/domain/apperr"
	domain "github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/repository"
	"github.com/sakaguchi-0725/echo-onion-arch/infra/persistence/model"
	"gorm.io/gorm"
)

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) repository.ProfileRepository {
	return &profileRepository{db}
}

func (p *profileRepository) DeleteByID(id domain.UserID) error {
	result := p.db.Delete(&model.Profile{}, "user_id = ?", id.String())
	if result.Error != nil {
		return apperr.NewApplicationError(apperr.ErrInternalError, "Failed to delete profile", result.Error)
	}

	if result.RowsAffected == 0 {
		return apperr.NewApplicationError(apperr.ErrNotFound, fmt.Sprintf("Profile with ID %s not found", id.String()), nil)
	}

	return nil
}

func (p *profileRepository) FindAll() ([]domain.Profile, error) {
	var profiles []model.Profile

	if err := p.db.Find(&profiles).Error; err != nil {
		return []domain.Profile{}, apperr.NewApplicationError(apperr.ErrInternalError, "Failed to retrieve profiles", err)
	}

	results := make([]domain.Profile, len(profiles))
	for i, v := range profiles {
		results[i] = model.ToDomainProfile(v)
	}

	return results, nil
}

func (p *profileRepository) FindByID(id domain.UserID) (domain.Profile, error) {
	var profile model.Profile

	if err := p.db.Where("user_id = ?", id.String()).First(&profile).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Profile{}, apperr.NewApplicationError(apperr.ErrNotFound, "Profile not found", fmt.Errorf("profile with ID %s not found", id))
		}

		return domain.Profile{}, apperr.NewApplicationError(apperr.ErrInternalError, "Failed to retrieve profile by ID", err)
	}

	return model.ToDomainProfile(profile), nil
}

func (p *profileRepository) Insert(profile domain.Profile) (domain.Profile, error) {
	modelProfile := model.ToModelProfile(profile)

	if err := p.db.Create(&modelProfile).Error; err != nil {
		return domain.Profile{}, apperr.NewApplicationError(apperr.ErrInternalError, "Failed to insert profile", err)
	}

	return model.ToDomainProfile(modelProfile), nil
}
