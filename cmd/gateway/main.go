package main

import (
	"log"
	"net/http"

	"github.com/Saad7890-web/self-healing-gateway/internal/config"
	"github.com/Saad7890-web/self-healing-gateway/internal/proxy"
)


func main() {
	cfg := config.Load()
	p,err := proxy.New(cfg.Backend.BaseURL, cfg.Backend.Timeout)

	if err != nil {
		log.Fatalf("Failed to create proxy: %v", err)
	}

	server := &http.Server{
		Addr: cfg.Server.Port,
		Handler: p,
		ReadTimeout: cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	log.Printf("API Gateway running on %s", cfg.Server.Port)
	log.Fatal(server.ListenAndServe())
}