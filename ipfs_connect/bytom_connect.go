package ipfs_connect

import (
	web "bytom-ipfs/http_web"
	"bytom-ipfs/providers"
)

var connection = web.NewWeb(providers.NewHTTPProvider("0.0.0.0:9888", 10, false))

func ConnectionBytomNode() (string, error) {
	bytom, err := connection.Bytom.CreateKey()
	if err != nil {
		return "", err

	}
	return bytom, nil

}
