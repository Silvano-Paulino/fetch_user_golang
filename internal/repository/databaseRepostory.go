package repository

import "github.com/silvano-paulino/internal/domain"

type Repository interface {
	GetUserById(id string) (*domain.User, error)
}
