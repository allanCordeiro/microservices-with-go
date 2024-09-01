package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/allancordeiro/microservices-with-go/movie/internal/gateway"
	"github.com/allancordeiro/microservices-with-go/pkg/discovery"
	model "github.com/allancordeiro/microservices-with-go/rating/pkg"
)

// Gateway defines an HTTP gateway for a rating service
type Gateway struct {
	registry discovery.Registry
}

// New creates a new HTTP gateway for a rating service
func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry: registry}
}

// GetAggregatedRating returns the aggregated rating for a record
// or ErrNotFound if there are no ratings for it
func (g *Gateway) GetAggregatedRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType) (float64, error) {
	addrs, err := g.getRandomServiceList(ctx, "rating")
	if err != nil {
		return 0, err
	}

	url := "http://" + addrs + "/rating"
	log.Printf("calling rating service GET: " + url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}

	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("id", string(recordID))
	values.Add("type", fmt.Sprintf("%v", recordType))
	req.URL.RawQuery = values.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return 0, gateway.ErrNotFound
	} else if resp.StatusCode/100 != 2 {
		return 0, fmt.Errorf("non-2xx response: %v", resp)
	}

	var v float64
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return 0, err
	}

	return v, nil
}

// PutRating writes a rating
func (g *Gateway) PutRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error {
	addrs, err := g.getRandomServiceList(ctx, "rating")
	if err != nil {
		return err
	}

	url := "http://" + addrs + "/rating"
	log.Printf("calling rating service GET: " + url)
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("id", string(recordID))
	values.Add("type", fmt.Sprintf("%v", recordType))
	values.Add("userId", string(rating.UserID))
	values.Add("value", fmt.Sprintf("%v", rating.Value))
	req.URL.RawQuery = values.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("non-2xx response: %v", resp)
	}

	return nil
}

func (g *Gateway) getRandomServiceList(ctx context.Context, serviceName string) (string, error) {
	addrs, err := g.registry.ServiceAddress(ctx, serviceName)
	if err != nil {
		return "", err
	}

	return addrs[rand.Intn(len(addrs))], nil
}
