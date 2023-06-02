package dto

import "auth/pkg/pb"

var ValidUserTypes = map[pb.UserType]bool{
	pb.UserType_CUSTOMER: true,
	pb.UserType_ADMIN:    true,
}
