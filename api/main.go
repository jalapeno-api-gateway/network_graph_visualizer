package main

import (
	"network_graph_api/handlers"
	"network_graph_api/jagw_connector"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	jagw_connector.GetJagwRequestClient()
	defer jagw_connector.CloseJagwRequestClient()

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}

	router.Use(cors.New(config))

	router.GET("/getNodes", handlers.GetAllNodeNames)
	router.GET("/getLinks", handlers.GetAllLinkConnections)

	router.Run()
}
