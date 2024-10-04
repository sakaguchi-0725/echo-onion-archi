//go:generate mockgen -source=profile_usecase.go -destination=../../mocks/application/usecase/mock_profile_usecase.go -package=mocks
package usecase

import (
	"github.com/sakaguchi-0725/echo-onion-arch/application/dto"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/apperr"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/repository"
)

type ProfileUsecase interface {
	FindByUserID(userID string) (dto.ProfileOutput, error)
	FindAll() ([]dto.ProfileOutput, error)
}

type profileUsecase struct {
	repo repository.ProfileRepository
}

func NewProfileUsecase(repo repository.ProfileRepository) ProfileUsecase {
	return &profileUsecase{repo}
}

func (p *profileUsecase) FindByUserID(userID string) (dto.ProfileOutput, error) {
	id, err := model.NewUserID(userID)
	if err != nil {
		return dto.ProfileOutput{}, apperr.NewApplicationError(apperr.ErrBadReqeust, "invalid request", err)
	}

	profile, err := p.repo.FindByID(id)
	if err != nil {
		return dto.ProfileOutput{}, err
	}

	return dto.ProfileOutput{
		Name: profile.Name,
		Role: profile.Role.String(),
	}, nil
}

func (p *profileUsecase) FindAll() ([]dto.ProfileOutput, error) {
	profiles, err := p.repo.FindAll()
	if err != nil {
		return []dto.ProfileOutput{}, err
	}

	output := make([]dto.ProfileOutput, len(profiles))
	for i, v := range profiles {
		p := dto.ProfileOutput{
			Name: v.Name,
			Role: v.Role.String(),
		}
		output[i] = p
	}

	return output, nil
}
