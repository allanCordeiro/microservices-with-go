package main

import (
	"log"
	"net/http"

	"github.com/allancordeiro/movieapp/movie/internal/controller/movie"
	metadatagateway "github.com/allancordeiro/movieapp/movie/internal/gateway/metadata/http"
	ratinggateway "github.com/allancordeiro/movieapp/movie/internal/gateway/rating/http"
	httphandler "github.com/allancordeiro/movieapp/movie/internal/handler/http"
)

func main() {
	log.Println("starting the movie service")
	medataGateway := metadatagateway.New("http://localhost:8081")
	ratingGateway := ratinggateway.New("http://localhost:8082")
	ctrl := movie.New(ratingGateway, medataGateway)
	h := httphandler.New(ctrl)

	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	log.Fatal(http.ListenAndServe(":8083", nil))
}
