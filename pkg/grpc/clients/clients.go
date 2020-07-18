package clients

import (
	"log"

	apiB "github.com/phev8/go_service_B/pkg/api"
	"google.golang.org/grpc"
)

func connectToGRPCServer(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to %s: %v", addr, err)
	}
	return conn
}

func ConnectToBService(addr string) (client apiB.ServiceBClient, close func() error) {
	serverConn := connectToGRPCServer(addr)
	return apiB.NewServiceBClient(serverConn), serverConn.Close
}
