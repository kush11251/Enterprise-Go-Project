package controller

import (
	"github.com/enterprise-go-project/pkg/model"
	"github.com/enterprise-go-project/pkg/service"
)

// Controller handles incoming requests
func NewController(service service.Service) *Controller {
	return &Controller{service: service}
}

type Controller struct {
	service service.Service
}

// Run starts the controller
func (c *Controller) Run() {
	// Handle incoming requests
	c.handleRequests()
}

func (c *Controller) handleRequests() {
	// Mock handling requests
	model := model.NewModel()
	c.service.ProcessModel(model)
}