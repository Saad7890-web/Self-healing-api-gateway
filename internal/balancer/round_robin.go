package balancer

import (
	"errors"
	"sync/atomic"

	"github.com/Saad7890-web/self-healing-gateway/internal/registry"
)	

type RoundRobin struct {
	registry *registry.Registry
	counter uint64
}

func NewRoundRobin(reg *registry.Registry) *RoundRobin{
	return &RoundRobin{
		registry: reg,
	}
}

func (rr *RoundRobin) Next() (*registry.Service, error){
	services := rr.registry.List()
	if len(services) == 0 {
		return nil, errors.New("No backend available")
	}

	index := atomic.AddUint64(&rr.counter, 1)
	return services[index%uint64(len(services))], nil
}