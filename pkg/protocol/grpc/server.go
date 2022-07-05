package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	v1 "github.com/aerodinamicat/thisisme01/pkg/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunServer(ctx context.Context, v1API v1.UserServiceServer, port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	v1.RegisterUserServiceServer(server, v1API)
	reflection.Register(server)

	//* Graceful stop
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("Shutting down gRPC server")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	log.Printf("Starting server and listening from:'%s'\n", port)
	return server.Serve(listener)
}
