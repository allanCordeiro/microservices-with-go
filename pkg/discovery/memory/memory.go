package memory

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/allancordeiro/movieapp/pkg/discovery"
)

type serviceName string
type instanceID string

type serviceInstance struct {
	hostPort   string
	lastActive time.Time
}

// Registry defines an in-memory service registry
type Registry struct {
	sync.RWMutex
	serviceAddrs map[serviceName]map[instanceID]*serviceInstance
}

// NewRegistry creates a new in-memory service registry instance
func NewRegistry() *Registry {
	return &Registry{serviceAddrs: map[serviceName]map[instanceID]*serviceInstance{}}
}

// Register creates a service instance record in the registry
func (r *Registry) Register(ctx context.Context, instID string, srvName string, hostPort string) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.serviceAddrs[serviceName(srvName)]; !ok {
		r.serviceAddrs[serviceName(srvName)] = map[instanceID]*serviceInstance{}
	}
	r.serviceAddrs[serviceName(srvName)][instanceID(instID)] = &serviceInstance{
		hostPort:   hostPort,
		lastActive: time.Now(),
	}
	return nil
}

// Deregister removes a service instance record from the registry
func (r *Registry) Deregister(ctx context.Context, instID string, srvName string) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.serviceAddrs[serviceName(srvName)]; !ok {
		return nil
	}
	delete(r.serviceAddrs[serviceName(srvName)], instanceID(instID))

	return nil
}

// ReportHealthyState is a push mechanism for reporting service's healthy state
func (r *Registry) ReportHealthyState(instID string, srvName string) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.serviceAddrs[serviceName(srvName)]; !ok {
		return errors.New("service is not registered yet")
	}

	if _, ok := r.serviceAddrs[serviceName(srvName)][instanceID(instID)]; !ok {
		return errors.New("service instance is not registered yet")
	}

	r.serviceAddrs[serviceName(srvName)][instanceID(instID)].lastActive = time.Now()
	return nil
}

// ServiceAddress returns list of addressess of given service
func (r *Registry) ServiceAddress(ctx context.Context, srvName string) ([]string, error) {
	r.RLock()
	defer r.RUnlock()

	if len(r.serviceAddrs[serviceName(srvName)]) == 0 {
		return nil, discovery.ErrNotFound
	}
	var res []string

	for _, i := range r.serviceAddrs[serviceName(srvName)] {
		if i.lastActive.Before(time.Now().Add(-5 * time.Second)) {
			continue
		}
		res = append(res, i.hostPort)
	}
	return res, nil
}
