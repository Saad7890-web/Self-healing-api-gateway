package registry

import (
	"net/url"
	"sync"
	"sync/atomic"

	"github.com/Saad7890-web/self-healing-gateway/internal/circuitbreaker"
)

type Status int32

const (
	Healthy Status = iota
	Unhealthy
)

type Service struct {
	ID     string
	URL    *url.URL
	status atomic.Int32
	Breaker *circuitbreaker.Breaker
}

func (s *Service) SetStatus(st Status) {
	s.status.Store(int32(st))
}

func (s *Service) Status() Status {
	return Status(s.status.Load())
}

type Registry struct {
	mu       sync.RWMutex
	services []*Service
}

func New() *Registry {
	return &Registry{}
}

func (r *Registry) Add(service *Service) {
	service.SetStatus(Healthy)
	r.mu.Lock()
	defer r.mu.Unlock()
	r.services = append(r.services, service)
}

func (r *Registry) HealthyServices() []*Service {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var healthy []*Service
	for _, s := range r.services {
		if s.Status() == Healthy {
			healthy = append(healthy, s)
		}
	}
	return healthy
}

func (r *Registry) All() []*Service {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.services
}
