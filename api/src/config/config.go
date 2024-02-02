package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DBConnection string = ""
	Port         uint64 = 0
	SecretKey    []byte
)

// Load environment variables
func Load() {
	var err error

	if err := godotenv.Load(); err != nil {
		log.Fatal("Config:", err)
	}

	Port, err = strconv.ParseUint(os.Getenv("API_PORT"), 10, 64)
	if err != nil {
		Port = 9000
	}

	DBConnection = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
