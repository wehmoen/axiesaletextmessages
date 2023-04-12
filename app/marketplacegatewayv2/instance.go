package marketplacegatewayv2

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/wehmoen/axiesaletextmessages/app/types"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func New(chain *ethclient.Client, address string) *Marketplacegatewayv2 {
	contractAddress := common.HexToAddress(
		strings.Replace(address, "ronin:", "0x", 1),
	)
	contract, err := NewMarketplacegatewayv2(contractAddress, chain)

	if err != nil {
		log.Fatalf("Failed to instantiate the Axie contract: %v", err)
	}

	return contract
}

func (m *Marketplacegatewayv2) OnOrderMatched(config *types.AlertConfig) (event.Subscription, error) {
	sink := make(chan *Marketplacegatewayv2OrderMatched)

	sub, err := m.WatchOrderMatched(&bind.WatchOpts{
		Context: context.Background(),
	}, sink)

	if err != nil {
		return nil, err
	}

	go m.monitorOrderMatched(sink, config)

	return sub, err
}

func (m *Marketplacegatewayv2) Wait(sub event.Subscription) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case err := <-sub.Err():
			sub.Unsubscribe()
			log.Fatalf("Error in subscription: %v", err)
		case <-interrupt:
			sub.Unsubscribe()
			log.Printf("Received interrupt signal, shutting down...")
			return
		}
	}
}

func containsHash(h common.Hash, hashSlice *[]common.Hash) bool {
	for _, hash := range *hashSlice {
		if h == hash {
			return true
		}
	}

	// If the hash is not found in the slice, add it and return false
	*hashSlice = append(*hashSlice, h)
	return false
}

func (m *Marketplacegatewayv2) monitorOrderMatched(sink chan *Marketplacegatewayv2OrderMatched, config *types.AlertConfig) {

	log.Printf("Begin to monitor the chain for sales from %s", config.Seller)

	var processed []common.Hash

	for {
		select {
		case e := <-sink:

			if strings.ToLower(e.Maker.String()) == strings.ToLower(config.Seller) {
				if containsHash(e.Raw.TxHash, &processed) {
					continue
				}
				log.Printf("We sold something... Processing...")

				transfer, err := e.getAxieTransfer(config.Chain, config.AxieContract)

				if err != nil {
					log.Printf("Sale is not an Axie sale!")
				}

				message := fmt.Sprintf("Hi! Your axie %s was sold to %s for %s WETH", transfer.TokenId.Text(10), e.Matcher.Hex(), e.formatPrice())

				log.Println(message)

				sms, _, err := config.Twilio.SendSMS(message)
				if err != nil {
					log.Printf("[GOERR] Failed to send SMS: %v", err)
				}

				log.Printf("SMS sent: %s", sms.Sid)
			}
		}
	}
}
