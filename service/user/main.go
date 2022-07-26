package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	"user/handler"
	"user/model"
	pb "user/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "user"
	version = "latest"
)

func main() {
	model.InitRedis()
	model.InitDb()

	consulReg := consul.NewRegistry()

	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulReg),
	)

	// Register handler
	err := pb.RegisterUserHandler(srv.Server(), new(handler.User))
	if err != nil {
		fmt.Println("RegisterUserHandler err: ", err)
		return
	}
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
