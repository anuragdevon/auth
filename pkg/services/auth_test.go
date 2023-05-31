package services

import (
	"auth/pkg/pb"
	"auth/pkg/repository/models"
	"auth/pkg/utils"
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthService_Register(t *testing.T) {
	authService := &AuthService{
		db:  db.DB,
		jwt: utils.JwtWrapper{},
	}
	t.Run("Register method to return StatusCreated a new user with valid usertype", func(t *testing.T) {
		registerReq := &pb.RegisterRequest{
			Email:    "testregister1@example.com",
			Password: "password123",
			UserType: pb.UserType_ADMIN,
		}
		registerRes, err := authService.Register(context.Background(), registerReq)
		assert.NoError(t, err)
		assert.Equal(t, int64(http.StatusCreated), registerRes.Status)
	})
	t.Run("Register method with an existing email should return StatusConflict", func(t *testing.T) {
		existingUser := models.User{
			Email:    "testregister2@example.com",
			Password: utils.HashPassword("password456"),
		}
		_, err := db.CreateUser(&existingUser)
		assert.NoError(t, err)

		registerReq := &pb.RegisterRequest{
			Email:    "testregister2@example.com",
			Password: "password789",
			UserType: pb.UserType_CUSTOMER,
		}
		registerRes, err := authService.Register(context.Background(), registerReq)
		assert.NoError(t, err)
		assert.Equal(t, int64(http.StatusConflict), registerRes.Status)
		assert.Equal(t, "email already registered", registerRes.Error)
	})
}

func TestAuthService_Login(t *testing.T) {
	authService := &AuthService{
		db:  db.DB,
		jwt: utils.JwtWrapper{},
	}
	t.Run("Login method to return StatusOK and a valid token for valid credentials ", func(t *testing.T) {
		user := models.User{
			Email:    "testlogin1@example.com",
			Password: utils.HashPassword("password123"),
			Usertype: pb.UserType_CUSTOMER.String(),
		}
		_, err := db.CreateUser(&user)
		assert.NoError(t, err)

		loginReq := &pb.LoginRequest{
			Email:    "testlogin1@example.com",
			Password: "password123",
		}
		loginRes, err := authService.Login(context.Background(), loginReq)
		assert.NoError(t, err)
		assert.Equal(t, int64(http.StatusOK), loginRes.Status)
		assert.NotEmpty(t, loginRes.Token)
	})

	t.Run("Login method to return StatusUnauthorized for invalid credentials", func(t *testing.T) {
		loginReq := &pb.LoginRequest{
			Email:    "testlogin1@example.com",
			Password: "wrongpassword",
		}
		loginRes, err := authService.Login(context.Background(), loginReq)
		assert.NoError(t, err)
		assert.Equal(t, int64(http.StatusUnauthorized), loginRes.Status)
		assert.Equal(t, "invalid login credentials", loginRes.Error)
	})
}

func TestAuthService_Validate(t *testing.T) {
	authService := &AuthService{
		db:  db.DB,
		jwt: utils.JwtWrapper{},
	}

	t.Run("Validate method with a valid token should return StatusOK and the user details", func(t *testing.T) {
		user := models.User{
			Email:    "testvalidate1@example.com",
			Password: utils.HashPassword("password123"),
		}
		_, err := db.CreateUser(&user)
		assert.NoError(t, err)

		token, err := authService.jwt.GenerateToken(user)
		assert.NoError(t, err)

		validateReq := &pb.ValidateRequest{
			Token: token,
		}
		validateRes, err := authService.Validate(context.Background(), validateReq)
		assert.NoError(t, err)
		assert.Equal(t, int64(http.StatusOK), validateRes.Status)
		assert.Equal(t, user.Id, validateRes.UserId)
		assert.Equal(t, pb.UserType_CUSTOMER, validateRes.UserType)
	})

	t.Run("Validate method with an invalid token should return StatusBadRequest", func(t *testing.T) {
		validateReq := &pb.ValidateRequest{
			Token: "invalidtoken",
		}
		validateRes, err := authService.Validate(context.Background(), validateReq)
		assert.NoError(t, err)
		assert.Equal(t, int64(http.StatusBadRequest), validateRes.Status)
		assert.NotEmpty(t, validateRes.Error)
	})

	t.Run("Validate method with a token for a non-existent user should return StatusNotFound", func(t *testing.T) {
		otherUser := models.User{
			Email:    "other@example.com",
			Password: utils.HashPassword("password456"),
		}
		token, err := authService.jwt.GenerateToken(otherUser)
		assert.NoError(t, err)

		validateReq := &pb.ValidateRequest{
			Token: token,
		}
		validateRes, err := authService.Validate(context.Background(), validateReq)
		assert.NoError(t, err)
		assert.Equal(t, int64(http.StatusNotFound), validateRes.Status)
		assert.Equal(t, "user not found", validateRes.Error)
	})
}
