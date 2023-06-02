package services

import (
	"context"
	"net/http"

	"auth/pkg/pb"
	"auth/pkg/repository"
	"auth/pkg/repository/models"
	"auth/pkg/utils"

	"gorm.io/gorm"
)

type AuthService struct {
	db  *gorm.DB
	jwt utils.JwtWrapper
	pb.UnimplementedAuthServiceServer
}

func NewAuthService(db *gorm.DB, jwt utils.JwtWrapper) *AuthService {
	return &AuthService{
		db:                             db,
		jwt:                            jwt,
		UnimplementedAuthServiceServer: pb.UnimplementedAuthServiceServer{},
	}
}

func (as *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	db := repository.Database{DB: as.db}
	_, err := db.GetUserByEmail(req.Email)

	if err == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "email already registered",
		}, nil
	}

	user.Email = req.Email
	user.Password = utils.HashPassword(req.Password)

	if !utils.IsValidUserType(req.UserType) {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  "invalid user type",
		}, nil
	}

	user.Usertype = req.UserType.String()

	_, err = db.CreateUser(&user)
	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
			Error:  "failed to register user",
		}, nil
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (as *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	db := repository.Database{DB: as.db}
	user, err := db.GetUserByEmail(req.Email)

	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "no user found with email",
		}, nil
	}

	match := utils.CheckPasswordHash(user.Password, req.Password)

	if !match {
		return &pb.LoginResponse{
			Status: http.StatusUnauthorized,
			Error:  "invalid login credentials",
		}, nil
	}

	token, _ := as.jwt.GenerateToken(*user)

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (as *AuthService) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	db := repository.Database{DB: as.db}
	claims, err := as.jwt.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	user, err := db.GetUserByEmail(claims.Email)
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "user not found",
		}, nil
	}
	userType := pb.UserType_CUSTOMER
	if user.Usertype == "ADMIN" {
		userType = pb.UserType_ADMIN
	}

	return &pb.ValidateResponse{
		Status:   http.StatusOK,
		UserId:   user.Id,
		UserType: userType,
	}, nil
}
