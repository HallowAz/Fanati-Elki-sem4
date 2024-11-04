package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const (
	ConfigEnvFilePath  = "config/.env"
	ServerPort         = ":8000"
	Host               = "83.166.237.142"
	DBType             = "postgres"
	connectToDBTimeout = 5
	SessionExpTime     = time.Hour * 128
	CookieLen          = 10
	CookieHeader       = "X-Auth-Token"
	RedisHost          = "localhost"
	RedisPort          = ":6379"
)

var (
	DBPort       string
	DBHost       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBConnString string
)

func InitConfig(configPath string) error {
	err := godotenv.Load(configPath)
	if err != nil {
		return err
	}

	DBHost = os.Getenv("HOST")
	DBPort = os.Getenv("POSTGRES_PORT")
	DBUser = os.Getenv("POSTGRES_USER")
	DBPassword = os.Getenv("POSTGRES_PASSWORD")
	DBName = os.Getenv("POSTGRES_DB")
	DBConnString = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=%d",
		DBType,
		DBUser,     // Замените на имя пользователя
		DBPassword, // Замените на пароль
		DBHost,     // Хост базы данных
		DBPort,     // Порт базы данных
		DBName,     // Название базы данных
		connectToDBTimeout)

	return nil
}
