package consul

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/allancordeiro/movieapp/pkg/discovery"
	consul "github.com/hashicorp/consul/api"
)

type Registry struct {
	client *consul.Client
}

// NewRegistry creates a new Consult-based service registry instance
func NewRegistry(addr string) (*Registry, error) {
	config := consul.DefaultConfig()
	config.Address = addr
	client, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Registry{client: client}, nil
}

// Register creates a service instance record in the registry
func (r *Registry) Register(ctx context.Context, instanceID string, serviceName string, hostPort string) error {
	parts := strings.Split(hostPort, ":")
	if len(parts) != 2 {
		return errors.New("hostPort must be in a form of <host>:<port>")
	}

	port, err := strconv.Atoi(parts[1])
	if err != nil {
		return err
	}

	return r.client.Agent().ServiceRegister(&consul.AgentServiceRegistration{
		Address: parts[0],
		ID:      instanceID,
		Name:    serviceName,
		Port:    port,
		Check:   &consul.AgentServiceCheck{CheckID: instanceID, TTL: "5s"},
	})
}

// Deregister removes a service instance record from the registry
func (r *Registry) Deregister(ctx context.Context, instanceID string, serviceName string) error {
	return r.client.Agent().ServiceDeregister(instanceID)
}

// ServiceAddress returns list of addressess of given service
func (r *Registry) ServiceAddress(ctx context.Context, serviceName string) ([]string, error) {
	entries, _, err := r.client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, err
	}
	if len(entries) == 0 {
		return nil, discovery.ErrNotFound
	}

	var res []string
	for _, e := range entries {
		res = append(res, fmt.Sprintf("%s:%d", e.Service.Address, e.Service.Port))
	}
	return res, nil
}

// ReportHealthyState is a push mechanism for reporting service's healthy state
func (r *Registry) ReportHealthyState(instanceID string, serviceName string) error {
	return r.client.Agent().PassTTL(instanceID, "")
	//return r.client.Agent().UpdateTTL(instanceID, "", "")
}
