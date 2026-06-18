package config

import "fmt"

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
	Sslmode  string
}

func (d *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		d.Host,
		d.User,
		d.Password,
		d.Dbname,
		d.Port,
		d.Sslmode,
	)
}
