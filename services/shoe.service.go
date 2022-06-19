package services

import "example.com/sarang-apis/models"

type ShoeService interface {
	CreateShoe(*models.Shoe, *string) error
	GetShoe(*string) (*models.Shoe, error)
	GetAll() ([]*models.Shoe, error)
	UpdateShoe(*models.Shoe) error
	DeleteShoe(*string) error
}
