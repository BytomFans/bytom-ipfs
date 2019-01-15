package dto

import "math/big"

type SyncingResponse struct {
	StartingBlock *big.Int `json:"startingBlock"`
	CurrentBlock  *big.Int `json:"currentBlock"`
	HighestBlock  *big.Int `json:"highestBlock"`
}
