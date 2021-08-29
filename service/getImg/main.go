package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"ihome/service/getImg/handler"

	"github.com/micro/go-micro/registry/consul"
	"ihome/service/getImg/model"
	getImg "ihome/service/getImg/proto/getImg"
)

func main() {

	//使用consul做服务发现
	consulReg := consul.NewRegistry()

	model.InitRedis()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.getImg"),
		micro.Version("latest"),
		micro.Registry(consulReg),
		micro.Address(":9981"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	getImg.RegisterGetImgHandler(service.Server(), new(handler.GetImg))



	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
