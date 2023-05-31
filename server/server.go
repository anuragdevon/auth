package server

import (
	"auth/pkg/config"
	"auth/pkg/pb"
	"auth/pkg/repository"
	"auth/pkg/services"
	"auth/pkg/utils"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func Run() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db := &repository.Database{}
	err = db.Connect(&c)
	if err != nil {
		log.Panic("failed to connect to database:", err)
	}
	defer db.Close()

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "auth",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth on", c.Port)

	newauthservice := services.NewAuthService(db.DB, jwt)

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, newauthservice)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
