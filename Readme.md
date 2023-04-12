# Axie Sale Text Messages

This tool sends you a text message every time one of your Axies is sold on the marketplace.

## Requirements

- [GoLang](https://golang.org/dl/) (1.20+)
- [Twilio](https://www.twilio.com/) account
- [Axie Infinity](https://axieinfinity.com/) account
- [Ronin Address](https://roninchain.com/)

## Installation

1. Clone the repository and navigate to the directory

```bash
git clone https://github.com/wehmoen/axiesaletextmessages.git
cd axiesaletextmessages
```

2. Install the dependencies

```bash
go get ./...
```

3. Build the executable

```bash
# Linux / Mac

go build -o axiesaletextmessages.run app/main.go

# Windows

go build -o axiesaletextmessages.exe .\app\main.go
```

## Usage

To run the program, you need to provide the following arguments:

- `-twilio_account_sid` - Your Twilio account SID
- `-twilio_auth_token` - Your Twilio auth token
- `-twilio_from` - Your Twilio phone number
- `-twilio_to` - Your phone number
- `-seller` - The address you want to monitor

For a full list of arguments, run the following command:

```bash 
# Linux / Mac

./axiesaletextmessages.run -help

# Windows

axiesaletextmessages.exe -help
```

## Move to Mainnet

By default the program will run on the Saigon Testnet!

To run the program on the Mainnet, you need to add the following arguments:

- `-axie_contract 0x32950db2a7164ae833121501c797d79e7b79d74c` - The mainnet Axie contract address
- `-marketplace 0xfff9ce5f71ca6178d3beecedb61e7eff1602950e` - The mainnet marketplace contract address
- `-rpc wss://rpc2.ronin.rest/ronin-mainnet/ws` - The mainnet RPC endpoint (Must be a websocket)
