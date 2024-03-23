package configs

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func GetEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func EnvRedisAddr() string {
	return os.Getenv("REDISADDR")
}

func EnvRedisPassword() string {
	return os.Getenv("REDISPASSWORD")
}

func EnvRedisDBName() int {
	name, _ := strconv.Atoi(os.Getenv("REDISDBNAME"))
	return name
}

func EnvMaxReq() int {
	K, _ := strconv.Atoi(os.Getenv("K"))
	return K
}

func EnvInterval() int64 {
	N, _ := strconv.Atoi(os.Getenv("N"))
	N *= 1000000000 
	return int64(time.Duration(N).Seconds())
}