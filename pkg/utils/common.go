package utils

import (
	"auth/pkg/pb"
	"auth/pkg/services/dto"
)

func IsValidUserType(userType pb.UserType) bool {
	_, ok := dto.ValidUserTypes[userType]
	return ok
}
