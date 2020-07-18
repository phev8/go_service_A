package server_a

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/phev8/go_service_A/pkg/api"
	"github.com/phev8/go_service_A/pkg/types"
	"google.golang.org/grpc"
)

type serviceAServer struct {
	clients *types.APIClients
}

func NewServiceAServer(
	clients *types.APIClients,
) api.ServiceAServer {
	return &serviceAServer{
		clients: clients,
	}
}

// RunServer runs gRPC service to publish ToDo service
func RunServer(ctx context.Context, port string,
	clients *types.APIClients,
) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// register service
	server := grpc.NewServer()
	api.RegisterServiceAServer(server, NewServiceAServer(
		clients,
	))

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")
	log.Println("wait connections on port " + port)
	return server.Serve(lis)
}
