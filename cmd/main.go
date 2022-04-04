package main

import (
  "fmt"
  "context"
  "net"
  "log"
  "net/http"

  "api/config"
  "api/internal/interfaces/server"
  "api/proto/person/protobuf"

  "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
  "github.com/grpc-ecosystem/grpc-gateway/runtime"
)

func init() {
  config.Load()
}

func main() {
  grpcServer := grpc.NewServer()

  personServer := server.NewPersonManagementServer()
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
