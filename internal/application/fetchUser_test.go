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

	t.Run("must fetch an user", func(t *testing.T) {
		// Arrange
		mockRepo := &MockDatabase{
			MockUserData: map[string]*domain.User{
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
	})

	t.Run("must fetch all users", func(t *testing.T) {
		// Arrange 
		mockRepo := &MockDatabase{
			MockAllUserData: []domain.User{
				{Id: "1", Name: "John Doe"},
				{Id: "2", Name: "Jane Smith"},
				{Id: "3", Name: "Alice Johnson"},
			},
		}

		// Act
		users, err := application.FetchAllUsers(mockRepo)

		// Assert
		if err != nil {
			t.Fatalf("expeccted an error, got %v", err)
		}

		if len(users) != 3 {
			t.Errorf("expeted 3 users, got %d", len(users))
		}
	})

	t.Run("Verify all data of users", func(t *testing.T) {
		// Arrange 
		mockRepo := &MockDatabase{
			MockAllUserData: []domain.User{
				{Id: "1", Name: "John Doe"},
				{Id: "2", Name: "Jane Smith"},
				{Id: "3", Name: "Alice Johnson"},
			},
		}

		expectedUsers := []string{"John Doe", "Jane Smith", "Alice Johnson"}

		// Act
		users, _ := application.FetchAllUsers(mockRepo)

		// Assert
		for i, user := range users {
			if user.Name != expectedUsers[i] {
				t.Errorf("expeted %s, got %s", expectedUsers[i], user.Name)
			}
		}
		
	})
}
