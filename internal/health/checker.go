package health

import (
	"log"
	"net/http"
	"time"

	"github.com/Saad7890-web/self-healing-gateway/internal/registry"
)


type Checker struct {
	registry *registry.Registry
	interval time.Duration
	timeout time.Duration
}

func New(reg *registry.Registry, interval, timeout time.Duration) *Checker{
	return &Checker{
		registry: reg,
		interval: interval,
		timeout: timeout,
	}
}

func (c *Checker) Start(){
	ticker := time.NewTicker(c.interval)

	go func(){
		for range ticker.C {
			c.checkAll()
		}
	}()
}

func (c *Checker) checkAll(){
	for _,service := range c.registry.All(){
		go c.check(service)
	}
}

func (c *Checker) check(s *registry.Service){
	client := http.Client{
		Timeout: c.timeout,
	}

	resp, err := client.Get(s.URL.String())
	if err != nil || resp.StatusCode >= 500 {
		s.SetStatus(registry.Unhealthy)
		log.Printf("[HEALTH] %s -> UNHEALTHY", s.ID)
		return
	}

	s.SetStatus(registry.Healthy)
	log.Printf("[HEALTH] %s -> HEALTY", s.ID)
}