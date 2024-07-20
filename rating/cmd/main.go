package main

import (
	"log"
	"net/http"

	"github.com/allancordeiro/movieapp/rating/internal/controller/rating"
	httphandler "github.com/allancordeiro/movieapp/rating/internal/handler/http"
	"github.com/allancordeiro/movieapp/rating/internal/repository/memory"
)

func main() {
	log.Println("starting the rating service")
	repo := memory.New()
	ctrl := rating.New(repo)
	h := httphandler.New(ctrl)

	http.Handle("/rating", http.HandlerFunc(h.Hande))
	log.Fatal(http.ListenAndServe(":8082", nil))
}
