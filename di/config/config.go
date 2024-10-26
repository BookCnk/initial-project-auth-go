package config

import "github.com/kelseyhightower/envconfig"

type AppConfig struct {
	CommonConfig   CommonConfig
	ServerConfig   ServerConfig
	DatabaseConfig DatabaseConfig
}

type CommonConfig struct {
	Env string `envconfig:"ENV" default:"local"`
}

type ServerConfig struct {
	Port string `envconfig:"PORT" default:"8080"`
}

type DatabaseConfig struct {
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	Port     int    `envconfig:"DB_PORT" default:"3306"`
	User     string `envconfig:"DB_USER" default:"root"`
	Password string `envconfig:"DB_PASSWORD" default:"root"`
	DBName   string `envconfig:"DB_NAME" default:"mydb"`
}

func GetConfig() AppConfig {
	var app AppConfig
	envconfig.MustProcess("APP", &app.CommonConfig)
	envconfig.MustProcess("APP", &app.ServerConfig)
	envconfig.MustProcess("APP", &app.DatabaseConfig)
	return app
}
