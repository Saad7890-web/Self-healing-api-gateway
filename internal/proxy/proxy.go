package proxy

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/Saad7890-web/self-healing-gateway/internal/balancer"
)

type Proxy struct {
	balancer balancer.Balancer
	timeout  time.Duration
}

func New(b balancer.Balancer, timeout time.Duration) *Proxy {
	return &Proxy{
		balancer: b,
		timeout:  timeout,
	}
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	services, err := p.balancer.Next()
	if err != nil {
		http.Error(w, "No backend available", http.StatusServiceUnavailable)
		return
	}
	if !services.Breaker.Allow(){
		http.Error(w, "Service temporarily unavailable", http.StatusServiceUnavailable)
		return
	}
	rp := httputil.NewSingleHostReverseProxy(services.URL)

	rp.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("proxy error: %v", err)
		http.Error(w, "Backend error", http.StatusBadGateway)
	}

	ctx, cancel := context.WithTimeout(r.Context(), p.timeout)
	defer cancel()

	r = r.WithContext(ctx)
	rp.ServeHTTP(w, r)
}
