package http_web

import (
	"bytom-ipfs/bytom"
	"bytom-ipfs/providers"
)

//web struct
type Web struct {
	Provider providers.ProviderInterface
	Bytom    *bytom.Bytom
}

func NewWeb(provider providers.ProviderInterface) *Web {
	web := new(Web)
	web.Provider = provider
	web.Bytom = bytom.NewBytom(provider)
	return web
}
