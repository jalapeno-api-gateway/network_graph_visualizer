package handlers

import (
	"network_graph_api/jagw_connector"
	"network_graph_api/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func removeLsNodePrefix(s string) string {
	return strings.Replace(s, "ls_node/", "", -1)
}

func GetAllLinkConnections(c *gin.Context) {
	nodes := jagw_connector.GetAllLinks()
	var result []models.Edges
	for _, x := range nodes.LsNodeEdges {
		result = append(result, models.Edges{Origin: removeLsNodePrefix(*x.From), Destination: removeLsNodePrefix(*x.To)})
	}
	c.JSON(200, result)

}
