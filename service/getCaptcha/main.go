package main

import (
	"getCaptcha/handler"
	pb "getCaptcha/proto"
	"github.com/go-micro/plugins/v4/registry/consul"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "getcaptcha"
	version = "latest"
)

func main() {
	consulReg := consul.NewRegistry()
	// Create service
	srv := micro.NewService(

		micro.Name(service),
		micro.Registry(consulReg), // 添加注册
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterGetCaptchaHandler(srv.Server(), new(handler.GetCaptcha))
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
