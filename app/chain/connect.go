package chain

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func New(rpc string) *ethclient.Client {

	if rpc[:3] != "wss" {
		panic("RPC must be a websocket")
	}

	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
