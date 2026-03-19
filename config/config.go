package config

import (
	"github.com/enterprise-go-project/pkg/utils"
)

// Config represents application configuration
func NewConfig() *Config {
	return &Config{}
}

type Config struct {}

// Load loads configuration from a file
func (c *Config) Load(filename string) {
	// Mock loading configuration from a file
	utils.Log("Loaded configuration from file: " + filename)
}