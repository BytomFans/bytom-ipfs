package api

import (
	"net/http"

	conn "bytom-ipfs/ipfs_connect"

	"github.com/gin-gonic/gin"
)

func AddFileToIpfs(c *gin.Context) {
	add := conn.ConnectionIpfs()
	c.JSON(http.StatusOK, gin.H{"fileds": add})
	return
}

func TestBytomBlock(c *gin.Context) {
	key, err := conn.ConnectionBytomNode()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"fileds": add})

}
