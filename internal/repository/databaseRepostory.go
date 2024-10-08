package repository

import "github.com/silvano-paulino/internal/domain"

type Repository interface {
	GetUserById(id string) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
	GetUserByIDWithDelay(id string) (*domain.User, error)
}
