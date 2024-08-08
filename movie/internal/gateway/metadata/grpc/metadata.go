package grpc

import (
	"context"
	"github.com/allancordeiro/movieapp/gen"
	"github.com/allancordeiro/movieapp/internal/grpcutil"
	"github.com/allancordeiro/movieapp/metadata/pkg"
	pkgmodel "github.com/allancordeiro/movieapp/metadata/pkg/model"
	"github.com/allancordeiro/movieapp/pkg/discovery"
)

// Gateway defines a movie metadata GRPC gateway
type Gateway struct {
	registry discovery.Registry
}

// New creates a new GRPC gateway for a movie metadata service
func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry: registry}
}

// Get returns movie metadata by a movie id
func (g *Gateway) Get(ctx context.Context, id string) (*model.Metadata, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "metadata", g.registry)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := gen.NewMetadataServiceClient(conn)
	resp, err := client.GetMetadata(ctx, &gen.GetMetadataRequest{MovieId: id})
	if err != nil {
		return nil, err
	}

	return pkgmodel.MetadataFromProto(resp.Metadata), nil
}
