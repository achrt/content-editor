package application

// import "github.com/caarlos0/env/v6" - для чтения переменных окружения

import (
	"content-editor/internal/infrastracture/db"
)

type Config struct {
	db.DBConfigSection
	Address string `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
}

// TODO: дописать метод чтения адреса из переменной окружения
// данные для постреса в docker_postgres_init.sql
func (c *Config) loadConfiguration() error {

	// c.Address =
	return nil
}
