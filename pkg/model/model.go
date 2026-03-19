package model

// Model represents a data structure
func NewModel() *Model {
	return &Model{}
}

type Model struct {
	Name string
	Age  int
}