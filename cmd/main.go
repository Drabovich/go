package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"todo-app"
	"todo-app/pkg/handler"
	"todo-app/pkg/repository"
	"todo-app/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	_ "github.com/lib/pq"
)

func main() {
	// JSON Errors
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// Старт конфига с переменными
	if err := initConfig(); err != nil {
		logrus.Fatalf("error ininitConfig")
	}

	// Старт env переменных
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error godotenv")
	}

	// Подключение к DB
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("error DB")
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// Старт сервера
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
