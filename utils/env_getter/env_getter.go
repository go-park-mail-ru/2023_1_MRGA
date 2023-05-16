package env_getter

import (
	"os"
	"strconv"
)

func GetEnvAsBool(name string, defaultVal bool) bool {
	valStr := os.Getenv(name)
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultVal
}

func GetHostFromEnv(name string) string {
	isLocal := GetEnvAsBool("IS_LOCAL", true)
	host := "0.0.0.0"
	if !isLocal {
		host = os.Getenv(name)
	}
	return host
}
