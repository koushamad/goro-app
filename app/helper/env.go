package helper

import (
	"github.com/joho/godotenv"
	"github.com/koushamad/goro-core/app/exception/throw"
	"os"
	"strconv"
	"strings"
)

func Env(env string, def string) string {
	p := Path(".env")
	err := godotenv.Load(p)
	throw.Fatal("Cannot Load Env File", 107, err)
	if os.Getenv(strings.ToUpper(env)) == "" {
		return def
	}
	return os.Getenv(strings.ToUpper(env))
}

func EnvToInt(env string, def int) int {
	value := Env(env, "")
	if value == "" {
		return def
	}
	config, err := strconv.Atoi(value)
	throw.Fatal("Cannot Get Env Worker Number", 108, err)
	return config
}

func EnvToBool(env string, def bool) bool {
	value := Env(env, "")
	if value == "" {
		return def
	}
	config, err := strconv.ParseBool(value)
	throw.Fatal("Cannot Get Env Worker Number", 109, err)
	return config
}
