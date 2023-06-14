package jagw_connector

import (
	"log"
	"network_graph_api/configs"

	"github.com/jalapeno-api-gateway/jagw-go/jagw"
	"google.golang.org/grpc"
)

var rsConnection *grpc.ClientConn
var rsClient jagw.RequestServiceClient

func GetJagwRequestClient() {
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
	rsClient = jagw.NewRequestServiceClient(rsConnection)
}

func CloseJagwRequestClient() {
	rsConnection.Close()
}
