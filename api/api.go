package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RetireOnBytomChain(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"fileds": ""})
	return
}
