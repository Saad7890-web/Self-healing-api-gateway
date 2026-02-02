package config

import "time"

type Config struct {
	Server ServerConfig
	Backend BackendConfig
}

type ServerConfig struct {
	Port string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	
}
	
type BackendConfig struct {
	BaseURL string
	Timeout time.Duration
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port:         ":8080",
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		Backend: BackendConfig{
			BaseURL: "http://localhost:9000",
			Timeout: 5 * time.Second,
		},
	}
}