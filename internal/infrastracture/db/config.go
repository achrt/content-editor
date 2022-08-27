package db

import "fmt"

type DBConfigSection struct {
	Host     string `env:"DBHOST" envDefault:"localhost"`
	Port     string `env:"DBPORT" envDefault:"5432"`
	Database string `env:"DBNAME"`
	Username string `env:"DB_USER_NAME"`
	Password string `env:"DB_PASS"`
}

func (cfg *DBConfigSection) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Database,
		cfg.Password,
	)
}
