package main

import (
	"github.com/micro/go-micro"
	"time"
	"github.com/Rakanixu/factorial-api/handler"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.factorial"),
		micro.RegisterTTL(time.Minute),
		micro.RegisterInterval(time.Second * 30),
	)

	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			new(handler.Factorial),
		),
	)

	if err := service.Run(); err!= nil {
		log.Fatal(err)
	}
}
