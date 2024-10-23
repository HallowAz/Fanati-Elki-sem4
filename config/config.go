package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Здесь все конфиг данные располагай, кроме паролей, пароли на виртуалке и локально будем хранить

const (
	ConfigEnvFilePath  = "config/.env"
	ServerPort         = ":8000"
	Host               = "83.166.237.142"
	DBType             = "postgres"
	connectToDBTimeout = 5
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
