package handlers

import (
	"network_graph_api/jagw_connector"
	"network_graph_api/models"

	"github.com/gin-gonic/gin"
)

func GetAllNodeNames(c *gin.Context) {
	nodes := jagw_connector.GetAllNodes()
	var result []models.Nodes
	for _, x := range nodes.LsNodes {
		result = append(result, models.Nodes{Name: *x.Name, Key: *x.Key})
	}
	c.JSON(200, result)

}
