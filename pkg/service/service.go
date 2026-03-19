package service

import (
	"github.com/enterprise-go-project/pkg/model"
	"github.com/enterprise-go-project/pkg/repository"
)

// Service handles business logic
func NewService() *Service {
	return &Service{repository: repository.NewRepository()}
}

type Service struct {
	repository repository.Repository
}

// ProcessModel processes a model
func (s *Service) ProcessModel(model model.Model) {
	s.repository.SaveModel(model)
}