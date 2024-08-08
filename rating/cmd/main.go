package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/allancordeiro/movieapp/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"

	"github.com/allancordeiro/movieapp/pkg/discovery"
	"github.com/allancordeiro/movieapp/pkg/discovery/consul"
	"github.com/allancordeiro/movieapp/rating/internal/controller/rating"
	grpchandler "github.com/allancordeiro/movieapp/rating/internal/handler/grpc"
	"github.com/allancordeiro/movieapp/rating/internal/repository/memory"
)

const serviceName = "rating"

func main() {
	var port int
	flag.IntVar(&port, "port", 8082, "API handler port")
	flag.Parse()

	log.Printf("starting the rating service on port %d\n", port)
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

	repo := memory.New()
	ctrl := rating.New(repo)
	h := grpchandler.New(ctrl)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	reflection.Register(srv)
	gen.RegisterRatingServiceServer(srv, h)
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}

}
