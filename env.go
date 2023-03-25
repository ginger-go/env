package env

import (
	"os"
	"reflect"
	"strings"
)

// GetEnv returns the value of the environment variable named by the key.
func GetEnv[T any](key string, defaultValue T) T {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return TestGetEnv(value, defaultValue) // This part of the code is extracted to a function for testing.
}

// TestGetEnv is used for testing. Don't use it in production.
func TestGetEnv[T any](value string, defaultValue T) T {
	output, err := convertString[T](value)
	if err != nil {
		return defaultValue
	}
	if reflect.ValueOf(output).IsZero() {
		return defaultValue
	}
	return output
}

// GetEnvList returns the list of value of the environment variable named by the key.
func GetEnvList[T any](key string, defaultValue []T) []T {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return TestGetEnvList(value, defaultValue) // This part of the code is extracted to a function for testing.
}

// TestGetEnvList is used for testing. Don't use it in production.
func TestGetEnvList[T any](value string, defaultValue []T) []T {
	li := strings.Split(value, ",")
	output := make([]T, 0)

	for _, v := range li {
		o, err := convertString[T](v)
		if err != nil {
			return defaultValue
		}
		output = append(output, o)
	}

	return output
}
