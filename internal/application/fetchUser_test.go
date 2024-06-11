package application_test

import (
	"errors"
	"testing"

	"github.com/silvano-paulino/internal/application"
	"github.com/silvano-paulino/internal/domain"
)

type MockDatabase struct {
	MockData map[string]*domain.User
}

func (m *MockDatabase) GetUserById(id string) (*domain.User, error) {
	user, ok := m.MockData[id]

	if !ok {
		return nil, errors.New("User Not found")
	}
	return user, nil
}

func TestFetchUser(t *testing.T) {
	// Arrange
	mockRepo := &MockDatabase{
		MockData: map[string]*domain.User{
			"1": {Id: "1", Name: "Silvano Paulino"},
		},
	}

	// Act
	user, err := application.FetchUser(mockRepo, "1")

	// assert
	if err != nil {
		t.Fatalf("Expeted no error, got %v", err)
	}

	if user.Name != "Silvano Paulino" {
		t.Errorf("Expected 'Silvano Paulino', got %s", user.Name)
	}
}
