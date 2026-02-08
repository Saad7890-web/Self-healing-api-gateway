package registry

import (
	"net/url"
	"sync"
)

type Service struct {
	ID string
	URL *url.URL
}

type Registry struct {
	mu sync.RWMutex
	services []*Service
}

func New() *Registry{
	return &Registry{}
}

func(r *Registry)Add(service *Service){
	r.mu.Lock()
	defer r.mu.Unlock()
	r.services = append(r.services, service)
}

func (r *Registry) List() []*Service {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.services
}