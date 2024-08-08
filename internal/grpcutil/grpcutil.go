package grpcutil

import (
	"context"
	"github.com/allancordeiro/movieapp/pkg/discovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math/rand"
)

// ServiceConnection attempts to select a random service instance and returns a GRPC connection to it
func ServiceConnection(ctx context.Context, serviceName string, registry discovery.Registry) (*grpc.ClientConn, error) {
	addrs, err := registry.ServiceAddress(ctx, serviceName)
	if err != nil {
		return nil, err
	}
	return grpc.DialContext(ctx, addrs[rand.Intn(len(addrs))], grpc.WithTransportCredentials(insecure.NewCredentials()))
}
