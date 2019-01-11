package ipfs

import (
	"fmt"
	"os"
	"strings"

	ipfs "github.com/ipfs/go-ipfs-api"
)

func Connection() {
	// Where your local node is running on localhost:5001
	port := "localhost:5001"
	connect := ipfs.NewShell(port)
	cid, err := connect.Add(strings.NewReader("hello world!,everyone"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	fmt.Printf("added %s", cid)
}
