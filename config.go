package mongo

import "fmt"

type Config struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
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
