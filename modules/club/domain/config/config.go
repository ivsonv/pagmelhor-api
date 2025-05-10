package config

import (
	"os"
	"strconv"
	"time"
)

var DefaultTimeout = time.Duration(getEnvInt("DEFAULT_CONTEXT_TIMEOUT", 30)) * time.Second

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
