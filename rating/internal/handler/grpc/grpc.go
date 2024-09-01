package grpc

import (
	"context"
	"errors"
	"github.com/allancordeiro/microservices-with-go/gen"
	"github.com/allancordeiro/microservices-with-go/rating/internal/controller/rating"
	model "github.com/allancordeiro/microservices-with-go/rating/pkg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handler defines a GRPC rating API handler
type Handler struct {
	gen.UnimplementedRatingServiceServer
	ctrl *rating.Controller
}

// New creates a new movie metadata GRPC handler
func New(ctrl *rating.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

// GetAggregatedRating return the aggregated rating for a record
func (h *Handler) GetAggregatedRating(ctx context.Context, req *gen.GetAggregatedRatingRequest) (*gen.GetAggregatedRatingResponse, error) {
	if req == nil || req.RecordId == "" || req.RecordType == "" {
		return nil, status.Error(codes.InvalidArgument, "nil req or empty id")
	}
	v, err := h.ctrl.GetAggregateRating(ctx, model.RecordID(req.RecordId), model.RecordType(req.RecordType))
	if err != nil && errors.Is(err, rating.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &gen.GetAggregatedRatingResponse{RatingValue: v}, nil
}

// PutRating writes a rating for a given record
func (h *Handler) PutRating(ctx context.Context, req *gen.PutRatingRequest) (*gen.PutRatingResponse, error) {
	if req == nil || req.RecordId == "" || req.UserId == "" {
		return nil, status.Error(codes.InvalidArgument, "nil req or empty user or empty id")
	}
	if err := h.ctrl.PutRating(ctx, model.RecordID(req.RecordId), model.RecordType(req.RecordType), &model.Rating{
		UserID: model.UserID(req.UserId),
		Value:  model.RatingValue(req.RatingValue),
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &gen.PutRatingResponse{}, nil
}
