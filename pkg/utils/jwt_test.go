package utils

import (
	"auth/pkg/repository/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJwtWrapperMethod(t *testing.T) {
	t.Run("GenerateToken test to successfully generate a token provided valid inputs", func(t *testing.T) {
		secretKey := "secret"
		issuer := "example.com"
		expirationHours := int64(1)

		jwtWrapper := &JwtWrapper{
			SecretKey:       secretKey,
			Issuer:          issuer,
			ExpirationHours: expirationHours,
		}

		user := models.User{
			Id:    1,
			Email: "test@example.com",
		}

		token, err := jwtWrapper.GenerateToken(user)
		assert.Nil(t, err)
		assert.NotEqual(t, "", token)
	})

	t.Run("ValidateToken test to successfully validate the claims of token provided", func(t *testing.T) {
		secretKey := "secret"
		issuer := "example.com"
		expirationHours := int64(1)

		jwtWrapper := &JwtWrapper{
			SecretKey:       secretKey,
			Issuer:          issuer,
			ExpirationHours: expirationHours,
		}

		user := models.User{
			Id:    1,
			Email: "test@example.com",
		}

		token, _ := jwtWrapper.GenerateToken(user)

		claims, err := jwtWrapper.ValidateToken(token)
		assert.Nil(t, err)
		assert.NotNil(t, claims)
		assert.Equal(t, user.Id, claims.Id)
		assert.Equal(t, user.Email, claims.Email)
	})
}
