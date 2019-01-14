package api

import (
	"net/http"

	conn "bytom-ipfs/ipfs_connect"

	"github.com/gin-gonic/gin"
)

func AddFileToIpfs(c *gin.Context) {
	connect := conn.ConnectionIpfs()
	c.JSON(http.StatusOK, gin.H{"fileds": connect})
	return
}
//发送一笔交易

