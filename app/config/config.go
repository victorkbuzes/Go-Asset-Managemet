package config

import (
	"fmt"
	"os"
)

func Config(key string) string {

	// Return the value of the variable
	return os.Getenv(key)
}
func GetOrFail(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("could not get value for %s ", key))
	}
	return value
}
