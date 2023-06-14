package jagw_connector

import (
	"context"
	"log"
	"time"

	"github.com/jalapeno-api-gateway/jagw-go/jagw"
)

func GetAllNodes() *jagw.LsNodeResponse {
	var response *jagw.LsNodeResponse
	var err error

	request := &jagw.TopologyRequest{}

	for i := 0; i < 3; i++ {
		response, err = rsClient.GetLsNodes(context.Background(), request)
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Error when calling GetLsNodes on request service: %s", err)
	}

	return response
}
