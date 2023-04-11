package marketplacegatewayv2

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

const (
	TRANSFER_TOPIC = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
)

type AxieTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
}

func (m *Marketplacegatewayv2OrderMatched) formatPrice() string {
	// Create a big.Int with the value 10^18
	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

	// Divide num by divisor
	result := new(big.Float).SetInt(m.SellerReceived)
	result.Quo(result, new(big.Float).SetInt(divisor))

	// Convert the result to a string with 6 decimal places
	return result.Text('f', 6)

}

func (m *Marketplacegatewayv2OrderMatched) getAxieTransfer(chain *ethclient.Client, axieContract common.Address) (*AxieTransfer, error) {

	tx, err := chain.TransactionReceipt(context.Background(), m.Raw.TxHash)

	if err != nil {
		return nil, err
	}

	if len(tx.Logs) == 0 {
		return nil, errors.New("no logs found")
	}

	for _, log := range tx.Logs {
		if log.Topics[0].Hex() == TRANSFER_TOPIC {
			if log.Address == axieContract {
				return &AxieTransfer{
					From:    common.HexToAddress(log.Topics[1].Hex()),
					To:      common.HexToAddress(log.Topics[2].Hex()),
					TokenId: log.Topics[3].Big(),
				}, nil
			}
		}
	}

	return nil, errors.New("no axie transfer found")
}
