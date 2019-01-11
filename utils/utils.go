package utils

import (
	"bytom-gin/complex/types"
	"bytom-gin/dto"
	"bytom-gin/providers"
)

// Utils - The Utils Module
type Utils struct {
	provider providers.ProviderInterface
}

// NewUtils - Utils Module constructor to set the default provider
func NewUtils(provider providers.ProviderInterface) *Utils {
	utils := new(Utils)
	utils.provider = provider
	return utils
}

func (utils *Utils) Sha3(data types.ComplexString) (string, error) {

	params := make([]string, 1)
	params[0] = data.ToHex()

	pointer := &dto.RequestResult{}

	err := utils.provider.SendRequest(pointer, "web3_sha3", params)

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}
