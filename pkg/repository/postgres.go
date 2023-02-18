package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ServerConfig struct {
	Port string `json:"port"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"db_Name"`
	SSLMode  string `json:"ssl_mode"`
}

type Config struct {
	ServerConfig `json:"server"`
	DBConfig     `json:"db"`
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DBConfig.Host, cfg.DBConfig.Port,
		cfg.DBConfig.Username, cfg.DBConfig.DBName,
		cfg.DBConfig.Password, cfg.DBConfig.SSLMode))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
