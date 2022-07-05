package cmd

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aerodinamicat/thisisme01/pkg/protocol/grpc"
	v1 "github.com/aerodinamicat/thisisme01/pkg/service/v1"
)

type Config struct {
	GRPCPort         string
	DatabaseHost     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseSchema   string
}

func RunServer() error {
	ctx := context.Background()

	cfg := &Config{
		GRPCPort:         "5070",
		DatabaseHost:     "localhost:54321",
		DatabaseUser:     "postgres",
		DatabasePassword: "mysecretpassword",
		DatabaseSchema:   "thisisme",
	}

	dbAddress := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseHost, cfg.DatabaseSchema)
	db, err := sql.Open("postgres", dbAddress)
	if err != nil {
		return fmt.Errorf("Failed to open database: '%v'", err.Error())
	}
	defer db.Close()

	v1API := v1.NewUserServiceServer(db)

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
