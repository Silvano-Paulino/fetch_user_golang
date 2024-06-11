package cmd

import (
	"errors"
	"time"

	"github.com/silvano-paulino/internal/domain"
)

type MockDatabase struct {
	MockUserData    map[string]*domain.User
	MockAllUserData []domain.User
}

func (m *MockDatabase) GetUserById(id string) (*domain.User, error) {
	user, ok := m.MockUserData[id]
	if !ok {
		return nil, errors.New("User Not found")
	}
	return user, nil
}

func (m *MockDatabase) GetAllUsers() ([]domain.User, error) {
	if len(m.MockAllUserData) == 0 {
		return nil, errors.New("no users found")
	}
	return m.MockAllUserData, nil
}

func (m *MockDatabase) GetUserByIDWithDelay(id string) (*domain.User, error) {
	time.Sleep(2 * time.Second) 
	user, exists := m.MockUserData[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}