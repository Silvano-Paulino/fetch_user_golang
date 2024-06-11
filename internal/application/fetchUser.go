package application

import (
	"errors"
	"time"

	"github.com/silvano-paulino/internal/domain"
	"github.com/silvano-paulino/internal/repository"
)

func FetchUser(db repository.Repository, id string) (*domain.User, error) {
	if id == "" {
		return nil, errors.New("empty user ID")
	}
	return db.GetUserById(id)
}

func FetchUserWithDelay(db repository.Repository, id string) (*domain.User, error) {
	time.Sleep(2 * time.Second)
	return db.GetUserByIDWithDelay(id)
}

func FetchAllUsers(db repository.Repository) ([]domain.User, error) {
	return db.GetAllUsers()
}
