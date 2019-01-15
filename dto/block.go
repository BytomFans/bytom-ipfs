package dto

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

type Block struct {
	Number       *big.Int           `json:"number"`
	Hash         string             `json:"hash"`
	ParentHash   string             `json:"parentHash"`
	Author       string             `json:"author,omitempty"`
	Miner        string             `json:"miner,omitempty"`
	Size         *big.Int           `json:"size"`
	GasUsed      *big.Int           `json:"gasUsed"`
	Nonce        *big.Int           `json:"nonce"`
	Timestamp    *big.Int           `json:"timestamp"`
	Transactions types.Transactions `transactions`
}

/**
 * How to un-marshal the block struct using the Big.Int rather than the
 * `complexReturn` type.
 */
func (b *Block) UnmarshalJSON(data []byte) error {
	type Alias Block
	temp := &struct {
		Number    string `json:"number"`
		Size      string `json:"size"`
		GasUsed   string `json:"gasUsed"`
		Nonce     string `json:"nonce"`
		Timestamp string `json:"timestamp"`
		*Alias
	}{
		Alias: (*Alias)(b),
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	num, success := big.NewInt(0).SetString(temp.Number[2:], 16)

	if !success {
		return errors.New(fmt.Sprintf("Error converting %s to bigInt", temp.Number))
	}

	size, success := big.NewInt(0).SetString(temp.Size[2:], 16)

	if !success {
		return errors.New(fmt.Sprintf("Error converting %s to bigInt", temp.Size))
	}

	gas, success := big.NewInt(0).SetString(temp.GasUsed[2:], 16)

	if !success {
		return errors.New(fmt.Sprintf("Error converting %s to bigInt", temp.GasUsed))
	}

	nonce, success := big.NewInt(0).SetString(temp.Nonce[2:], 16)

	if !success {
		return errors.New(fmt.Sprintf("Error converting %s to bigInt", temp.Nonce))
	}

	timestamp, success := big.NewInt(0).SetString(temp.Timestamp[2:], 16)

	if !success {
		return errors.New(fmt.Sprintf("Error converting %s to bigInt", temp.Timestamp))
	}

	b.Number = num
	b.Size = size
	b.GasUsed = gas
	b.Nonce = nonce
	b.Timestamp = timestamp

	return nil
}
