package adapters

import (
	"github.com/enterprise-go-project/pkg/model"
)

// Adapter converts a model to a different format
func NewAdapter() *Adapter {
	return &Adapter{}
}

type Adapter struct {}

// Convert converts a model to a different format
func (a *Adapter) Convert(model model.Model) *ConvertedModel {
	return &ConvertedModel{Name: model.Name, Age: model.Age}
}

type ConvertedModel struct {
	Name string
	Age  int
}