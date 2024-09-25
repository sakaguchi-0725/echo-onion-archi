package repository

import (
	"errors"
	"fmt"

	domain "github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/repository"
	"github.com/sakaguchi-0725/echo-onion-arch/infra/repository/model"
	"gorm.io/gorm"
)

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) repository.ProfileRepository {
	return &profileRepository{db}
}

func (p *profileRepository) DeleteByID(id domain.UserID) error {
	result := p.db.Delete(&model.Profile{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("profile with ID %s not found", id)
	}

	return nil
}

func (p *profileRepository) FindAll() ([]domain.Profile, error) {
	var profiles []model.Profile

	if err := p.db.Find(&profiles).Error; err != nil {
		return []domain.Profile{}, err
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
			return domain.Profile{}, fmt.Errorf("profile with ID %s not found", id)
		}
		return domain.Profile{}, err
	}

	return model.ToDomainProfile(profile), nil
}

func (p *profileRepository) Insert(profile domain.Profile) (domain.Profile, error) {
	modelProfile := model.ToModelProfile(profile)

	if err := p.db.Create(&modelProfile).Error; err != nil {
		return domain.Profile{}, err
	}

	return model.ToDomainProfile(modelProfile), nil
}
