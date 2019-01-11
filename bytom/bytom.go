package bytom

import (
	"bytom-ipfs/providers"
	"ginblock/dto"
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

func (bytom *Bytom) Contract(jsonInterface string) (*Contract, error) {
	return bytom.NewContract(jsonInterface)
}

func (bytom *Bytom) sendTransaction(hash string) (*dto.Receipt, error) {
	params := make([]string, 1)
	params[0] = hash

	err := bytom.provider.SendRequest(params)
	if err != nil {
		return nil, err
	}

	return nil, err

}
