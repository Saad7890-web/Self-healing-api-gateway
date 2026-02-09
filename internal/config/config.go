package config

import "time"

type Config struct {
	Server ServerConfig
	Backends []BackendConfig
}

type ServerConfig struct {
	Port string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	
}
	
type BackendConfig struct {
	ID string
	BaseURL string
	
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port:         ":8080",
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		Backends: []BackendConfig{
			{ID:"backend-1", BaseURL: "http://localhost:9000"},
			{ID: "backend-2", BaseURL: "http://localhost:9001"},
		},
	}
}