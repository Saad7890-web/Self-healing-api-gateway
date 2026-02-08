package proxy

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/Saad7890-web/self-healing-gateway/internal/registry"
)

type Proxy struct {
	registry *registry.Registry
	timeout  time.Duration
}

func New(reg *registry.Registry, timeout time.Duration) *Proxy {
	return &Proxy{
		registry: reg,
		timeout:  timeout,
	}
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	services := p.registry.List()
	if len(services) == 0 {
		http.Error(w, "No backend available", http.StatusServiceUnavailable)
		return
	}

	target := services[0] 

	rp := httputil.NewSingleHostReverseProxy(target.URL)

	rp.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("proxy error: %v", err)
		http.Error(w, "Backend error", http.StatusBadGateway)
	}

	ctx, cancel := context.WithTimeout(r.Context(), p.timeout)
	defer cancel()

	r = r.WithContext(ctx)
	rp.ServeHTTP(w, r)
}
