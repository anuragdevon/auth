package utils

import (
	"auth/pkg/pb"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidUserType(t *testing.T) {
	validUserTypes := map[pb.UserType]bool{
		pb.UserType_CUSTOMER: true,
		pb.UserType_ADMIN:    true,
	}

	t.Run("Valid user types should return true", func(t *testing.T) {
		for userType := range validUserTypes {
			isValid := IsValidUserType(userType)
			assert.True(t, isValid, "Expected valid user type")
		}
	})

	t.Run("Invalid user type should return false", func(t *testing.T) {
		invalidUserType := pb.UserType(-1)
		isValid := IsValidUserType(invalidUserType)
		assert.False(t, isValid, "Expected invalid user type")
	})
}
