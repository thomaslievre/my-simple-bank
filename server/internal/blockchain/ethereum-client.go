package blockchain

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func NewEthClient(rpcURL string) *ethclient.Client {
	client, err := ethclient.DialContext(context.Background(), rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum node: %v", err)
	}
	return client
}
