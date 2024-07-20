package main

import (
	"log"
	"net/http"

	"github.com/allancordeiro/movieapp/metadata/internal/controller/metadata"
	httphandler "github.com/allancordeiro/movieapp/metadata/internal/handler/http"
	"github.com/allancordeiro/movieapp/metadata/internal/repository/memory"
)

func main() {
	log.Println("starting the movie metadata service")
	repo := memory.New()
	ctrl := metadata.New(repo)
	h := httphandler.New(ctrl)

	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
