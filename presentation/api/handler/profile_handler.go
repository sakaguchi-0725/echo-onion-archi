package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sakaguchi-0725/echo-onion-arch/application/usecase"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/dto"
)

type ProfileHandler interface {
	GetProfile(c echo.Context) error
	GetAllProfiles(c echo.Context) error
}

type profileHandler struct {
	usecase usecase.ProfileUsecase
}

func NewProfileHandler(usecase usecase.ProfileUsecase) ProfileHandler {
	return &profileHandler{usecase}
}

func (p *profileHandler) GetProfile(c echo.Context) error {
	userID := c.Get("userID").(string)

	profile, err := p.usecase.FindByUserID(userID)
	if err != nil {
		return err
	}

	res := dto.ProfileResponse{
		Name: profile.Name,
		Role: profile.Role,
	}

	return c.JSON(http.StatusOK, res)
}

func (p *profileHandler) GetAllProfiles(c echo.Context) error {
	profiles, err := p.usecase.FindAll()
	if err != nil {
		return err
	}

	res := make([]dto.ProfileResponse, len(profiles))
	for i, v := range profiles {
		p := dto.ProfileResponse{
			Name: v.Name,
			Role: v.Role,
		}

		res[i] = p
	}

	return c.JSON(http.StatusOK, res)
}
