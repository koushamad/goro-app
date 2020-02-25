package helper

import (
	"github.com/joho/godotenv"
	"github.com/koushamad/goro/app/exception/throw"
	"os"
	"strconv"
	"strings"
)

func Env(env string) string {
	err := godotenv.Load()
	throw.Fatal("Cannot Load Env File", 107, err)
	return os.Getenv(strings.ToUpper(env))
}

func EnvToInt(env string) int {
	err := godotenv.Load()
	throw.Fatal("Cannot Load Env File", 108, err)
	config, err := strconv.Atoi(Env(env))
	throw.Fatal("Cannot Get Env Worker Number", 109, err)
	return config
}

func EnvToBool(env string) bool {
	err := godotenv.Load()
	throw.Fatal("Cannot Load Env File", 110, err)
	config, err := strconv.ParseBool(Env(env))
	throw.Fatal("Cannot Get Env Worker Number", 111, err)
	return config
}
