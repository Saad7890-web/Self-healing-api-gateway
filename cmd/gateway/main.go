package main

import (
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/Saad7890-web/self-healing-gateway/internal/balancer"
	"github.com/Saad7890-web/self-healing-gateway/internal/config"
	"github.com/Saad7890-web/self-healing-gateway/internal/health"
	"github.com/Saad7890-web/self-healing-gateway/internal/proxy"
	"github.com/Saad7890-web/self-healing-gateway/internal/registry"
)


func main() {
	cfg := config.Load()
	reg := registry.New()

	for _,b:= range cfg.Backends{
		u,err := url.Parse(b.BaseURL)
		if err != nil {
			log.Fatalf("invalid backend url %s", b.BaseURL)
		}

		reg.Add(&registry.Service{
			ID: b.ID,
			URL: u,
		})
	}

	checker := health.New(reg, 5*time.Second, 2*time.Second)
	checker.Start()
	rr := balancer.NewRoundRobin(reg)
	p := proxy.New(rr, cfg.Server.WriteTimeout)

	server := &http.Server{
		Addr: cfg.Server.Port,
		Handler: p,
		ReadTimeout: cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	log.Printf("API Gateway running on %s", cfg.Server.Port)
	log.Fatal(server.ListenAndServe())
}