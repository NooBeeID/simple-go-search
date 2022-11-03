package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadConfig(filename string) error {
	err := godotenv.Load(filename)
	return err
}

func GetString(key ConfigKey, def string) string {
	val := os.Getenv(string(key))
	if val == "" {
		return def
	}
	return val
}

func GetInt8(key ConfigKey, def int8) int8 {
	val := os.Getenv(string(key))
	if val == "" {
		return def
	}

	valInt, err := strconv.Atoi(val)
	if err != nil {
		log.Printf("error : value `%v` cant be cast to int. Will used default value", val)
		return def
	}

	return int8(valInt)
}
