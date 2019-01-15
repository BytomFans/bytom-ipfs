package ipfs_connect

import (
	"fmt"
	"os"
	"strings"

	ipfs "github.com/ipfs/go-ipfs-api"
)

var port_str = "localhost:5001"

//add file
func ConnectionIpfs() string {
	// Where your local node is running on localhost:5001
	port := port_str
	connect := ipfs.NewShell(port)
	cid, err := connect.Add(strings.NewReader("hello world!,everyone"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	fmt.Printf("added %s", cid)
	return cid
}

