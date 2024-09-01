package model

import model "github.com/allancordeiro/microservices-with-go/metadata/pkg"

// MovieDetails includes movie metadata its aggregated rating
type MovieDetails struct {
	Rating   *float64       `json:"rating,omitEmpty"`
	Metadata model.Metadata `json:"metadata"`
}
