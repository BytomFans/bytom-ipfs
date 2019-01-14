package bytom_rpc

import "eth-gin/providers"

// retire
type Web struct {
	Provider providers.ProviderInterface
}

//Retire Module constructor to set the default provider, Bytom, Net and Personal
func NewWeb(provider providers.ProviderInterface) *Web {
	web := new(Web)
	web.Provider = provider
	return web
}
