package model

import "errors"

type Category struct {
	ID   uint
	Name string
}

func NewCategory(name string) (Category, error) {
	if name == "" {
		return Category{}, errors.New("name is required")
	}

	return Category{
		Name: name,
	}, nil
}

func RecreateCategory(id uint, name string) Category {
	return Category{
		ID:   id,
		Name: name,
	}
}
