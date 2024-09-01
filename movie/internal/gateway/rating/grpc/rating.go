package grpc

import (
	"context"
	"github.com/allancordeiro/microservices-with-go/gen"
	"github.com/allancordeiro/microservices-with-go/internal/grpcutil"
	"github.com/allancordeiro/microservices-with-go/pkg/discovery"
	model "github.com/allancordeiro/microservices-with-go/rating/pkg"
)

// Gateway define an GRPC gateway for a rating service
type Gateway struct {
	registry discovery.Registry
}

// New creates a new GRPC gateway for a rating service
func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry: registry}
}

// GetAggregatedRating returns the aggregated rating for a record
// or ErrNotFound if there are no ratings for it
func (g *Gateway) GetAggregatedRating(ctx context.Context, recordId model.RecordID, recordType model.RecordType) (float64, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "rating", g.registry)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	client := gen.NewRatingServiceClient(conn)
	resp, err := client.GetAggregatedRating(ctx, &gen.GetAggregatedRatingRequest{
		RecordId:   string(recordId),
		RecordType: string(recordType),
	})
	if err != nil {
		return 0, err
	}
	return resp.RatingValue, nil
}
