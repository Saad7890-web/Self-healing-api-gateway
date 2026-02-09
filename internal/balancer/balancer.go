package balancer

import "github.com/Saad7890-web/self-healing-gateway/internal/registry"




type Balancer interface {
	Next() (*registry.Service, error)
}