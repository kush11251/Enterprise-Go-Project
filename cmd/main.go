package main

import (
	"github.com/enterprise-go-project/pkg/controller"
	"github.com/enterprise-go-project/pkg/service"
)

func main() {
	service := service.NewService()
	controller := controller.NewController(service)
	controller.Run()
}