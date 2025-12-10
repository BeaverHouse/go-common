package env

import (
	"os"
	"strconv"
)

// GetEnv retrieves an environment variable with a default string value
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetIntEnv retrieves an environment variable with a default integer value
func GetIntEnv(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return valueInt
}

// IsGoEnv checks if the environment (GO_ENV) is the specified value
func IsGoEnv(env EnvType) bool {
	return GetEnv("GO_ENV", string(LocalEnv)) == string(env)
}
