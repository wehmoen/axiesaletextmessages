package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wehmoen/axiesaletextmessages/app/twilio"
)

type AlertConfig struct {
	Seller       string            `json:"seller"`
	Twilio       *twilio.Twilio    `json:"twilio"`
	Chain        *ethclient.Client `json:"chain"`
	AxieContract common.Address
}
