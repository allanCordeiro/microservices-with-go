package main

import (
	"context"
	"fmt"
	"github.com/allancordeiro/microservices-with-go/gen"
	grpchandler "github.com/allancordeiro/microservices-with-go/metadata/internal/handler/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v3"
	"log"
	"net"
	"os"
	"time"

	"github.com/allancordeiro/microservices-with-go/metadata/internal/controller/metadata"
	"github.com/allancordeiro/microservices-with-go/metadata/internal/repository/memory"
	"github.com/allancordeiro/microservices-with-go/pkg/discovery"
	"github.com/allancordeiro/microservices-with-go/pkg/discovery/consul"
)

const serviceName = "metadata"

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
	log.Printf("starting the movie metadata service on port %s\n", port)

	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%s", port)); err != nil {
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
