package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/allancordeiro/movieapp/gen"
	grpchandler "github.com/allancordeiro/movieapp/metadata/internal/handler/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"

	"github.com/allancordeiro/movieapp/metadata/internal/controller/metadata"
	"github.com/allancordeiro/movieapp/metadata/internal/repository/memory"
	"github.com/allancordeiro/movieapp/pkg/discovery"
	"github.com/allancordeiro/movieapp/pkg/discovery/consul"
)

const serviceName = "metadata"

func main() {
	var port int
	flag.IntVar(&port, "port", 8081, "API handler port")
	flag.Parse()
	log.Printf("starting the movie metadata service on port %d\n", port)
	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		panic(err)
	}
	go func() {
		for {
			err := registry.ReportHealthyState(instanceID, serviceName)
			if err != nil {
				log.Println("failed to report healty state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer registry.Deregister(ctx, instanceID, serviceName)

	repo := memory.New()
	ctrl := metadata.New(repo)
	h := grpchandler.New(ctrl)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	reflection.Register(srv)
	gen.RegisterMetadataServiceServer(srv, h)
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}

}
