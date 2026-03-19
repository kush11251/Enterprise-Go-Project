package config

import (
	"os"
)

// EnvConfig represents environment configuration
func NewEnvConfig() *EnvConfig {
	return &EnvConfig{}
}

type EnvConfig struct {}

// Load loads environment configuration
func (c *EnvConfig) Load() {
	// Mock loading environment configuration
	env := os.Getenv("ENVIRONMENT")
	if env == "dev" {
		// Load development configuration
	} else {
		// Load production configuration
	}
}