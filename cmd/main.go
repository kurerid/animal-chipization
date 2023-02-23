package main

import (
	"animal-chipization/pkg/handler"
	"animal-chipization/pkg/repository"
	"animal-chipization/pkg/service"
	"encoding/json"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("failed to read environments: %s", err.Error())
	}
	config, err := loadConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		logrus.Fatal(err)
	}
	db, err := repository.NewPostgresDB(repository.Config{
		DBConfig: repository.DBConfig{
			Host:     config.DBConfig.Host,
			Port:     config.DBConfig.Port,
			Username: config.DBConfig.Username,
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   config.DBConfig.DBName,
			SSLMode:  config.DBConfig.SSLMode,
		},
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	if err := run(config, handlers); err != nil {
		logrus.Fatal(err)
	}
}

func loadConfig(path string) (*repository.Config, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		logrus.Println(err)
		return nil, err
	}
	var config repository.Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		logrus.Println(err)
		return nil, err
	}
	return &config, err
}

func run(cfg *repository.Config, handler *handler.Handler) error {
	server := handler.InitRoutes()
	err := server.Run(":" + cfg.ServerConfig.Port)
	if err != nil {
		return err
	}
	return nil
}
