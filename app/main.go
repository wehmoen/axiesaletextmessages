package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/inancgumus/screen"
	"github.com/wehmoen/axiesaletextmessages/app/chain"
	"github.com/wehmoen/axiesaletextmessages/app/config"
	"github.com/wehmoen/axiesaletextmessages/app/marketplacegatewayv2"
	"github.com/wehmoen/axiesaletextmessages/app/twilio"
	"github.com/wehmoen/axiesaletextmessages/app/types"
	"log"
)

func main() {

	screen.Clear()

	cfg := config.New()

	ronin := chain.New(cfg["rpc"])

	twilioClient := twilio.NewTwilio(cfg["twilio_account_sid"], cfg["twilio_auth_token"], cfg["twilio_from"], cfg["twilio_to"])

	marketplace := marketplacegatewayv2.New(ronin, cfg["marketplace"])

	sub, err := marketplace.OnOrderMatched(&types.AlertConfig{
		Seller:       cfg["seller"],
		Twilio:       twilioClient,
		Chain:        ronin,
		AxieContract: common.HexToAddress(cfg["axie_contract"]),
	})

	if err != nil {
		sub.Unsubscribe()
		log.Fatalf("Failed to subscribe to OrderMatched event: %v", err)
	}

	marketplace.Wait(sub)
}
