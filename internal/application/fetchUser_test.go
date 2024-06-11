package application_test

import (
	"errors"
	"testing"
	"time"

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

func (m *MockDatabase) GetUserByIDWithDelay(id string) (*domain.User, error) {
	time.Sleep(2 * time.Second) // Simula um delay
	user, exists := m.MockUserData[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
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
		if user.Id != "1" {
			t.Error("Expected user")
		}

	})

	t.Run("must return error for non-existent user", func(t *testing.T) {
		// Arrange
		mockRepo := &MockDatabase{
			MockUserData: map[string]*domain.User{},
		}

		// Act
		user, err := application.FetchUser(mockRepo, "2")

		// Assert
		if err == nil {
			t.Fatal("Expected an error, but got nil")
		}

		if user != nil {
			t.Errorf("Expected nil user, but got %v", user)
		}
	})

	t.Run("must handle repository delay", func(t *testing.T) {
		// Arrange
		mockRepo := &MockDatabase{
			MockUserData: map[string]*domain.User{
				"4": {Id: "4", Name: "Silvano Paulino"},
			},
		}

		// Act
		user, err := application.FetchUserWithDelay(mockRepo, "4")

		// Assert
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if user == nil || user.Id != "4" {
			t.Errorf("Expected user with ID '1', but got %v", user)
		}
	})

	t.Run("must handle empty ID", func(t *testing.T) {
		// Arrange
		mockRepo := &MockDatabase{
			MockUserData: map[string]*domain.User{},
		}

		// Act
		user, err := application.FetchUser(mockRepo, "")

		// Assert
		if err == nil {
			t.Fatal("Expected an error for empty ID, but got nil")
		}

		if user != nil {
			t.Errorf("Expected nil user for empty ID, but got %v", user)
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

		if user.Id != "1" {
			t.Errorf("Expected '1', got %s", user.Id)
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
