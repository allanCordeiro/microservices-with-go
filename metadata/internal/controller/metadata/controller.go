package metadata

import (
	"context"
	"errors"

	"github.com/allancordeiro/microservices-with-go/metadata/internal/repository"
	model "github.com/allancordeiro/microservices-with-go/metadata/pkg"
)

// ErrNotFound is returned when a requested record is not found
var ErrNotFound = errors.New("not found")

type medataRepository interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
}

// Controller defines a metadata service controller
type Controller struct {
	repo medataRepository
}

// New creates a metadata service controller
func New(repo medataRepository) *Controller {
	return &Controller{repo: repo}
}

// Get returns movie metadata by id.
func (c *Controller) Get(ctx context.Context, id string) (*model.Metadata, error) {
	res, err := c.repo.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}

	return res, err
}
