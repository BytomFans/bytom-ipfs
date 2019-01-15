package bytom

import (
	"bytom-ipfs/dto"
	"bytom-ipfs/providers"
)

// The Block Modle
type Bytom struct {
	provider providers.ProviderInterface
}

// NewBytom
func NewBytom(provider providers.ProviderInterface) *Bytom {
	bytom := new(Bytom)
	bytom.provider = provider
	return bytom
}

//create-key
func (btm *Bytom) CreateKey() (string, error) {

	pointer := &dto.RequestResult{}

	err := btm.provider.SendRequest(pointer, "create_key", nil)

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}

//create-account
func (btm *Bytom) CreateAccount() (string, error) {
	pointer := &dto.RequestResult{}

	err := btm.provider.SendRequest(pointer, "create_account", nil)

	if err != nil {
		return "", err
	}

	return pointer.ToString()
}

//build-transaction
func (btm *Bytom) BuildTransaction() (string, error) {
	pointer := &dto.RequestResult{}

	err := btm.provider.SendRequest(pointer, "build_transaction", nil)

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}
