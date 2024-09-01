package main

import (
	"context"
	"fmt"
	"github.com/allancordeiro/microservices-with-go/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v3"
	"log"
	"net"
	"os"
	"time"

	"github.com/allancordeiro/microservices-with-go/movie/internal/controller/movie"
	metadatagateway "github.com/allancordeiro/microservices-with-go/movie/internal/gateway/metadata/http"
	ratinggateway "github.com/allancordeiro/microservices-with-go/movie/internal/gateway/rating/http"
	grpchandler "github.com/allancordeiro/microservices-with-go/movie/internal/handler/grpc"
	"github.com/allancordeiro/microservices-with-go/pkg/discovery"
	"github.com/allancordeiro/microservices-with-go/pkg/discovery/consul"
)

const serviceName = "movie"

func main() {
	f, err := os.Open("configs/base.yaml")
	if err != nil {
		panic(err)
	}
	var cfg config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		panic(err)
	}
	port := cfg.APIConfig.Port
	log.Printf("starting the movie service on port %d\n", port)
	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	err = registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port))
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			err := registry.ReportHealthyState(instanceID, serviceName)
			if err != nil {
				log.Println("failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer registry.Deregister(ctx, instanceID, serviceName)

	medataGateway := metadatagateway.New(registry)
	ratingGateway := ratinggateway.New(registry)
	ctrl := movie.New(ratingGateway, medataGateway)
	h := grpchandler.New(ctrl)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	reflection.Register(srv)
	gen.RegisterMovieServiceServer(srv, h)
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}

}
