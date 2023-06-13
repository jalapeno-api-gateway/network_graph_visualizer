package jagw_connector

import (
	"log"
	"network_graph_api/configs"

	"github.com/jalapeno-api-gateway/jagw-go/jagw"
	"google.golang.org/grpc"
)

/*
TODO:
A better way to connect to GRPC should be found.
Unfortunately the defer function closes the connection to early and I don't want to clutter the main.go
*/
var rsConnection *grpc.ClientConn

func getJagwRequestClient() jagw.RequestServiceClient {
	config := configs.LoadConfig()
	requestServiceEndpoint := jagw.JagwEndpoint{
		EndpointAddress: config.JAGW.Server,
		EndpointPort:    config.JAGW.RequestServicePort,
	}
	var rsErr error
	rsConnection, rsErr = jagw.NewJagwConnection(requestServiceEndpoint)
	if rsErr != nil {
		log.Fatalf("Failed to setup request service connection: %s", rsErr)
	}
	// defer rsConnection.Close()

	return jagw.NewRequestServiceClient(rsConnection)
}

func closeJagwRequestClient() {
	rsConnection.Close()
}
