package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"api/config"
	"api/internal/dependency"
	"api/internal/interface/server"
	"api/proto/person/protobuf"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	config.Load()
}

func main() {
	grpcServer := grpc.NewServer()

	personAS := dependency.NewPersonApplicationService(dependency.NewPersonRepository())
	personServer := server.NewPersonManagementServer(personAS)
	protobuf.RegisterPersonManagementServer(grpcServer, personServer)

	reflection.Register(grpcServer)

	grpcAddress := fmt.Sprintf(":%s", config.Global.GRPCPort)
	grpcListener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf("Cannot start server: %s\n", err)
	}

	go grpcServer.Serve(grpcListener)

	go func() {
		if err := runGateway(); err != nil {
			log.Fatalf("Cannot serve gateway: %s\n", err)
		}
	}()
	select {}
}

func runGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := fmt.Sprintf("localhost:%s", config.Global.GRPCPort)
	err := protobuf.RegisterPersonManagementHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(fmt.Sprintf(":%s", config.Global.RESTPort), mux)
}
