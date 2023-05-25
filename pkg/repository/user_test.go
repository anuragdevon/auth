package repository

import (
	"auth/pkg/repository/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	t.Run("CreateUser method to successfully create user with valid data", func(t *testing.T) {
		user := &models.User{
			Email:    "test@example.com",
			Password: "test123",
		}
		userID, err := db.CreateUser(user)
		if err != nil {
			t.Fatalf("failed to create user: %v", err)
		}

		if userID == 0 {
			t.Errorf("CreateUser() did not set user ID")
		}
	})

	t.Run("CreateUser method to return error for duplicate emailID", func(t *testing.T) {
		user := &models.User{
			Email:    "anuragkar1@gmail.com",
			Password: "password123",
		}

		_, err := db.CreateUser(user)
		if err != nil {
			t.Fatalf("Failed to create user: %v", err)
		}

		duplicateUser := &models.User{
			Email:    "anuragkar1@gmail.com",
			Password: "password456",
		}

		userId, err := db.CreateUser(duplicateUser)
		assert.NotNil(t, err)
		assert.Equal(t, 0, userId)
	})

	t.Run("GetUserByID method to return valid user for valid userID", func(t *testing.T) {
		user := &models.User{
			Email:    "test4@example.com",
			Password: "test123",
		}
		userID, err := db.CreateUser(user)
		if err != nil {
			t.Fatalf("failed to create user: %v", err)
		}

		fetchedUser, err := db.GetUserByID(userID)
		if err != nil {
			t.Fatalf("GetUserByID() error = %v, want nil", err)
		}

		if fetchedUser.Email != user.Email {
			t.Errorf("GetUserByID() EmailID = %v, want %v", fetchedUser.Email, user.Email)
		}
	})

	t.Run("GetUserByEmail method to return valid user for valid emailID", func(t *testing.T) {
		userEmail := "testuser99@example.com"
		userPassword := "password123"
		userID, err := db.CreateUser(&models.User{Email: userEmail, Password: userPassword})
		if err != nil {
			t.Fatalf("failed to create test user: %v", err)
		}

		user, err := db.GetUserByEmail(userEmail)
		if err != nil {
			t.Fatalf("failed to get user by email: %v", err)
		}

		if int(user.Id) != userID {
			t.Errorf("GetUserByEmail() returned wrong user ID, got %d, want %d", user.Id, userID)
		}

		if user.Email != userEmail {
			t.Errorf("GetUserByEmail() returned wrong email, got %s, want %s", user.Email, userEmail)
		}
	})

	t.Run("GetUserByEmail method to return error for invalid emailID", func(t *testing.T) {
		userEmail := "nonexistent@example.com"
		_, err := db.GetUserByEmail(userEmail)
		assert.NotNil(t, err)
	})
}
