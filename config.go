package mongo

import "fmt"

type Config struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func (c Config) getUri() string {
	return fmt.Sprintf("mongodb://%s:%d", c.Host, c.Port)
}

func GetDefaultConfig() Config {
	return Config{
		Host:     "localhost",
		Port:     27017,
		Database: "database",
		Username: "username",
		Password: "password",
	}
}
