package application_test

import (
	"errors"
	"testing"

	"github.com/silvano-paulino/internal/application"
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

func TestFetchUser(t *testing.T) {
	t.Run("must exists an user", func(t *testing.T) {
		// Arrange
		mockRepo := &MockDatabase{
			MockUserData: map[string]*domain.User{
				"1": {Id: "1", Name: "Silvano Paulino"},
			},
		}

		// Act
		user, _ := application.FetchUser(mockRepo, "1")

		// assert
		if user.Name == "" {
			t.Error("Expected 'user'")
		}

	})
}
