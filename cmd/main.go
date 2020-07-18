package main

import (
	"context"
	"log"

	gc "github.com/phev8/go_service_A/pkg/grpc/clients"
	"github.com/phev8/go_service_A/pkg/grpc/server_a"
	"github.com/phev8/go_service_A/pkg/types"
)

func main() {
	// ---> client connections
	clients := &types.APIClients{}
	bc, close := gc.ConnectToBService("localhost:4302")
	defer close()
	clients.ServiceB = bc

	// <---
	ctx := context.Background()

	if err := server_a.RunServer(
		ctx,
		"4301",
		clients,
	); err != nil {
		log.Fatal(err)
	}

}
