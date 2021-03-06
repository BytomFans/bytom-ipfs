package main

import (
	"bytom-ipfs/api"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("ginblock.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.POST("/ipfs", api.AddFileToIpfs)
	router.POST("/block", api.TestBytomBlock)
	router.Run(":8090")
}
