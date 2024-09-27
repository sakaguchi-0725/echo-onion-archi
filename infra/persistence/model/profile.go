package model

import (
	"time"

	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
)

type Profile struct {
	UserID    string    `gorm:"primaryKey;not null"`
	Name      string    `gorm:"not null"`
	Role      string    `gorm:"type:enum('admin', 'general') not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func ToModelProfile(profile model.Profile) Profile {
	return Profile{
		UserID: profile.UserID.String(),
		Name:   profile.Name,
		Role:   profile.Role.String(),
	}
}

func ToDomainProfile(profile Profile) model.Profile {
	return model.RecreateProfile(model.UserID(profile.UserID), profile.Name, model.UserRole(profile.Role))
}
