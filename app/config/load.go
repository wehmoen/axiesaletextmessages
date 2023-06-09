package config

import "flag"

const (
	defaultRPC          = "wss://rpc2.ronin.rest/ronin-testnet/ws"
	defaultAxieContract = "0xcaca1c072d26e46686d932686015207fbe08fdb8"
	defaultMarketplace  = "0xa7be1f3c47f191217a6e0bf3f7d05ca39427e0af"
)

func New() map[string]string {
	rpc := flag.String("rpc", defaultRPC, "Ronin Websocket RPC")
	axie_contract := flag.String("axie_contract", defaultAxieContract, "Axie Contract")
	marketplace := flag.String("marketplace", defaultMarketplace, "Marketplace Contract")

	seller := flag.String("seller", "", "Seller Address to monitor")

	twilioAccountSid := flag.String("twilio_account_sid", "", "Twilio Account SID")
	twilioAuthToken := flag.String("twilio_auth_token", "", "Twilio Auth Token")
	twilioFrom := flag.String("twilio_from", "", "Twilio From Number")
	twilioTo := flag.String("twilio_to", "", "Number that will receive the Text Messages")

	flag.Parse()

	return map[string]string{
		"rpc":           *rpc,
		"axie_contract": *axie_contract,
		"marketplace":   *marketplace,
		"seller":        *seller,

		"twilio_account_sid": *twilioAccountSid,
		"twilio_auth_token":  *twilioAuthToken,
		"twilio_from":        *twilioFrom,
		"twilio_to":          *twilioTo,
	}

}
