package jagw_connector

import (
	"context"
	"log"
	"time"

	"github.com/jalapeno-api-gateway/jagw-go/jagw"
)

func GetAllLinks() *jagw.LsNodeEdgeResponse {
	var response *jagw.LsNodeEdgeResponse
	var err error

	client := getJagwRequestClient()
	defer closeJagwRequestClient()

	request := &jagw.TopologyRequest{}

	for i := 0; i < 3; i++ {
		response, err = client.GetLsNodeEdges(context.Background(), request)
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Error when calling GetLsNodeEdges on request service: %s", err)
	}
	return response
}
