package config

import (
	"os"
	"strconv"
)

type Config struct {
	Service Service `yaml:"service"`
}

type Service struct {
	Name      string `yaml:"name"`
	Version   string `yaml:"version"`
	Address   string `yaml:"address"`
	Port      int    `yaml:"port"`
	JwtSecret string `yaml:"jwt_secret"`
	Env       string `yaml:"env"`
}

func (c *Config) Init() {
	svcPort, _ := strconv.Atoi(os.Getenv("SERVICE_PORT"))
	c.Service = Service{
		Name:    os.Getenv("SERVICE_NAME"),
		Version: os.Getenv("SERVICE_VERSION"),
		Address: os.Getenv("SERVICE_ADDRESS"),
		Port:    svcPort,
		Env:     os.Getenv("SERVICE_ENV"),
	}
}
