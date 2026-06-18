package config

import "fmt"

type AppConfig struct {
	AppPort string
	AppEnv  string
}

func (a *AppConfig) IsDevelopment() bool {
	return a.AppEnv == "development"
}

func (a *AppConfig) IsProduction() bool {
	return a.AppEnv == "production"
}

func (a *AppConfig) GetPort() string {
	return fmt.Sprintf(":%s", a.AppPort)
}
