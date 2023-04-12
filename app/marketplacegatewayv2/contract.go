// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package marketplacegatewayv2

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MarketAssetAsset is an auto generated low-level Go binding around an user-defined struct.
type MarketAssetAsset struct {
	Erc      uint8
	Addr     common.Address
	Id       *big.Int
	Quantity *big.Int
}

// MarketOrderOrder is an auto generated low-level Go binding around an user-defined struct.
type MarketOrderOrder struct {
	Maker               common.Address
	Kind                uint8
	Assets              []MarketAssetAsset
	ExpiredAt           *big.Int
	PaymentToken        common.Address
	StartedAt           *big.Int
	BasePrice           *big.Int
	EndedAt             *big.Int
	EndedPrice          *big.Int
	ExpectedState       *big.Int
	Nonce               *big.Int
	MarketFeePercentage *big.Int
}

// Marketplacegatewayv2MetaData contains all meta data concerning the Marketplacegatewayv2 contract.
var Marketplacegatewayv2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"MakerNonceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"matcher\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumMarketOrder.OrderKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settlePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sellerReceived\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketFeePercentage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketFeeTaken\",\"type\":\"uint256\"}],\"name\":\"OrderMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referralAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ReferralRewardTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INTERFACE_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MARKET_OPERATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"enumMarketOrder.OrderKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"enumMarketAsset.TokenStandard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structMarketAsset.Asset[]\",\"name\":\"assets\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endedPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marketFeePercentage\",\"type\":\"uint256\"}],\"internalType\":\"structMarketOrder.Order\",\"name\":\"_order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumMarketAsset.TokenStandard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structMarketAsset.Asset[]\",\"name\":\"_assets\",\"type\":\"tuple[]\"}],\"name\":\"getState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increaseNonceMaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_maker\",\"type\":\"address\"}],\"name\":\"makerNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"orderFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"enumMarketOrder.OrderKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"enumMarketAsset.TokenStandard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structMarketAsset.Asset[]\",\"name\":\"assets\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endedPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marketFeePercentage\",\"type\":\"uint256\"}],\"internalType\":\"structMarketOrder.Order\",\"name\":\"_order\",\"type\":\"tuple\"}],\"name\":\"orderValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_expectedState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_settlePrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_referralAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"enumMarketOrder.OrderKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"enumMarketAsset.TokenStandard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structMarketAsset.Asset[]\",\"name\":\"assets\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endedPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marketFeePercentage\",\"type\":\"uint256\"}],\"internalType\":\"structMarketOrder.Order\",\"name\":\"_order\",\"type\":\"tuple\"}],\"name\":\"settleOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_expectedState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_settlePrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_referralAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_path\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"enumMarketOrder.OrderKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"enumMarketAsset.TokenStandard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structMarketAsset.Asset[]\",\"name\":\"assets\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endedPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marketFeePercentage\",\"type\":\"uint256\"}],\"internalType\":\"structMarketOrder.Order\",\"name\":\"_order\",\"type\":\"tuple\"}],\"name\":\"swapTokensAndSettleOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_hashList\",\"type\":\"bytes32[]\"}],\"name\":\"tryBulkCancelOrderByHash\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"_orderAlreadyFinalized\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"AllowedAllPaymentTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"_interfaces\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"name\":\"InterfacesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIKatanaRouter\",\"name\":\"KatanaRouterContract\",\"type\":\"address\"}],\"name\":\"KatanaRouterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minMarketFeePercentage\",\"type\":\"uint256\"}],\"name\":\"MinMarketFeePercentageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"PaymentTokensAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"referralFeature\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"contractIReferral\",\"name\":\"referralContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"autoTransferReferralReward\",\"type\":\"bool\"}],\"name\":\"ReferralConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"TreasuryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"WRONUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WRON\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowedAllPaymentTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"allowedPaymentToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_interface\",\"type\":\"string\"}],\"name\":\"getInterface\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReferralConfig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"referralFeature_\",\"type\":\"bool\"},{\"internalType\":\"contractIReferral\",\"name\":\"referralContract_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"autoTransferReferralReward_\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wronAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasuryAddr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_allowedAllPaymentTokens\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_referralFeature\",\"type\":\"bool\"},{\"internalType\":\"contractIReferral\",\"name\":\"_referralContract\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_autoTransferReferralReward\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_minMarketFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_tokenPayments\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"_interfaces\",\"type\":\"string[]\"},{\"internalType\":\"address[][2]\",\"name\":\"_addresses\",\"type\":\"address[][2]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_interface\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"interactWith\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minMarketFeePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_flag\",\"type\":\"bool\"}],\"name\":\"setAllowedAllPaymentTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_interfaces\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"setInterfaces\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIKatanaRouter\",\"name\":\"_routerContract\",\"type\":\"address\"}],\"name\":\"setKatanaRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minMarketFeePercentage\",\"type\":\"uint256\"}],\"name\":\"setMinMarketFeePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"_allowed\",\"type\":\"bool\"}],\"name\":\"setPaymentTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"referralFeature_\",\"type\":\"bool\"},{\"internalType\":\"contractIReferral\",\"name\":\"referralContract_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"autoTransferReferralReward_\",\"type\":\"bool\"}],\"name\":\"setReferralConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wronAddr\",\"type\":\"address\"}],\"name\":\"setWRON\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Marketplacegatewayv2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Marketplacegatewayv2MetaData.ABI instead.
var Marketplacegatewayv2ABI = Marketplacegatewayv2MetaData.ABI

// Marketplacegatewayv2 is an auto generated Go binding around an Ethereum contract.
type Marketplacegatewayv2 struct {
	Marketplacegatewayv2Caller     // Read-only binding to the contract
	Marketplacegatewayv2Transactor // Write-only binding to the contract
	Marketplacegatewayv2Filterer   // Log filterer for contract events
}

// Marketplacegatewayv2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Marketplacegatewayv2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Marketplacegatewayv2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Marketplacegatewayv2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Marketplacegatewayv2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Marketplacegatewayv2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Marketplacegatewayv2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Marketplacegatewayv2Session struct {
	Contract     *Marketplacegatewayv2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Marketplacegatewayv2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Marketplacegatewayv2CallerSession struct {
	Contract *Marketplacegatewayv2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// Marketplacegatewayv2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Marketplacegatewayv2TransactorSession struct {
	Contract     *Marketplacegatewayv2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// Marketplacegatewayv2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Marketplacegatewayv2Raw struct {
	Contract *Marketplacegatewayv2 // Generic contract binding to access the raw methods on
}

// Marketplacegatewayv2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Marketplacegatewayv2CallerRaw struct {
	Contract *Marketplacegatewayv2Caller // Generic read-only contract binding to access the raw methods on
}

// Marketplacegatewayv2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Marketplacegatewayv2TransactorRaw struct {
	Contract *Marketplacegatewayv2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplacegatewayv2 creates a new instance of Marketplacegatewayv2, bound to a specific deployed contract.
func NewMarketplacegatewayv2(address common.Address, backend bind.ContractBackend) (*Marketplacegatewayv2, error) {
	contract, err := bindMarketplacegatewayv2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplacegatewayv2{Marketplacegatewayv2Caller: Marketplacegatewayv2Caller{contract: contract}, Marketplacegatewayv2Transactor: Marketplacegatewayv2Transactor{contract: contract}, Marketplacegatewayv2Filterer: Marketplacegatewayv2Filterer{contract: contract}}, nil
}

// NewMarketplacegatewayv2Caller creates a new read-only instance of Marketplacegatewayv2, bound to a specific deployed contract.
func NewMarketplacegatewayv2Caller(address common.Address, caller bind.ContractCaller) (*Marketplacegatewayv2Caller, error) {
	contract, err := bindMarketplacegatewayv2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Marketplacegatewayv2Caller{contract: contract}, nil
}

// NewMarketplacegatewayv2Transactor creates a new write-only instance of Marketplacegatewayv2, bound to a specific deployed contract.
func NewMarketplacegatewayv2Transactor(address common.Address, transactor bind.ContractTransactor) (*Marketplacegatewayv2Transactor, error) {
	contract, err := bindMarketplacegatewayv2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Marketplacegatewayv2Transactor{contract: contract}, nil
}

// NewMarketplacegatewayv2Filterer creates a new log filterer instance of Marketplacegatewayv2, bound to a specific deployed contract.
func NewMarketplacegatewayv2Filterer(address common.Address, filterer bind.ContractFilterer) (*Marketplacegatewayv2Filterer, error) {
	contract, err := bindMarketplacegatewayv2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Marketplacegatewayv2Filterer{contract: contract}, nil
}

// bindMarketplacegatewayv2 binds a generic wrapper to an already deployed contract.
func bindMarketplacegatewayv2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Marketplacegatewayv2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplacegatewayv2 *Marketplacegatewayv2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplacegatewayv2.Contract.Marketplacegatewayv2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplacegatewayv2 *Marketplacegatewayv2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.Marketplacegatewayv2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplacegatewayv2 *Marketplacegatewayv2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.Marketplacegatewayv2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplacegatewayv2 *Marketplacegatewayv2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplacegatewayv2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplacegatewayv2 *Marketplacegatewayv2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplacegatewayv2 *Marketplacegatewayv2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Marketplacegatewayv2.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _Marketplacegatewayv2.Contract.DEFAULTADMINROLE(&_Marketplacegatewayv2.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Marketplacegatewayv2 *Marketplacegatewayv2CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Marketplacegatewayv2.Contract.DEFAULTADMINROLE(&_Marketplacegatewayv2.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Marketplacegatewayv2.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _Marketplacegatewayv2.Contract.DOMAINSEPARATOR(&_Marketplacegatewayv2.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Marketplacegatewayv2 *Marketplacegatewayv2CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Marketplacegatewayv2.Contract.DOMAINSEPARATOR(&_Marketplacegatewayv2.CallOpts)
}

// INTERFACENAME is a free data retrieval call binding the contract method 0x25df3b9b.
//
// Solidity: function INTERFACE_NAME() view returns(string)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Caller) INTERFACENAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Marketplacegatewayv2.contract.Call(opts, &out, "INTERFACE_NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// INTERFACENAME is a free data retrieval call binding the contract method 0x25df3b9b.
//
// Solidity: function INTERFACE_NAME() view returns(string)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) INTERFACENAME() (string, error) {
	return _Marketplacegatewayv2.Contract.INTERFACENAME(&_Marketplacegatewayv2.CallOpts)
}

// INTERFACENAME is a free data retrieval call binding the contract method 0x25df3b9b.
//
// Solidity: function INTERFACE_NAME() view returns(string)
func (_Marketplacegatewayv2 *Marketplacegatewayv2CallerSession) INTERFACENAME() (string, error) {
	return _Marketplacegatewayv2.Contract.INTERFACENAME(&_Marketplacegatewayv2.CallOpts)
}

// MARKETOPERATOR is a free data retrieval call binding the contract method 0xbe958cc6.
//
// Solidity: function MARKET_OPERATOR() view returns(bytes32)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Caller) MARKETOPERATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Marketplacegatewayv2.contract.Call(opts, &out, "MARKET_OPERATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MARKETOPERATOR is a free data retrieval call binding the contract method 0xbe958cc6.
//
// Solidity: function MARKET_OPERATOR() view returns(bytes32)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) MARKETOPERATOR() ([32]byte, error) {
	return _Marketplacegatewayv2.Contract.MARKETOPERATOR(&_Marketplacegatewayv2.CallOpts)
}

// MARKETOPERATOR is a free data retrieval call binding the contract method 0xbe958cc6.
//
// Solidity: function MARKET_OPERATOR() view returns(bytes32)
func (_Marketplacegatewayv2 *Marketplacegatewayv2CallerSession) MARKETOPERATOR() ([32]byte, error) {
	return _Marketplacegatewayv2.Contract.MARKETOPERATOR(&_Marketplacegatewayv2.CallOpts)
}

// WRON is a free data retrieval call binding the contract method 0x281388a8.
//
// Solidity: function WRON() view returns(address)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Caller) WRON(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplacegatewayv2.contract.Call(opts, &out, "WRON")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WRON is a free data retrieval call binding the contract method 0x281388a8.
//
// Solidity: function WRON() view returns(address)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) WRON() (common.Address, error) {
	return _Marketplacegatewayv2.Contract.WRON(&_Marketplacegatewayv2.CallOpts)
}

// WRON is a free data retrieval call binding the contract method 0x281388a8.
//
// Solidity: function WRON() view returns(address)
func (_Marketplacegatewayv2 *Marketplacegatewayv2CallerSession) WRON() (common.Address, error) {
	return _Marketplacegatewayv2.Contract.WRON(&_Marketplacegatewayv2.CallOpts)
}

// AllowedAllPaymentTokens is a free data retrieval call binding the contract method 0xd080fc88.
//
// Solidity: function allowedAllPaymentTokens() view returns(bool)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Caller) AllowedAllPaymentTokens(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Marketplacegatewayv2.contract.Call(opts, &out, "allowedAllPaymentTokens")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedAllPaymentTokens is a free data retrieval call binding the contract method 0xd080fc88.
//
// Solidity: function allowedAllPaymentTokens() view returns(bool)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) AllowedAllPaymentTokens() (bool, error) {
	return _Marketplacegatewayv2.Contract.AllowedAllPaymentTokens(&_Marketplacegatewayv2.CallOpts)
}

// AllowedAllPaymentTokens is a free data retrieval call binding the contract method 0xd080fc88.
//
// Solidity: function allowedAllPaymentTokens() view returns(bool)
func (_Marketplacegatewayv2 *Marketplacegatewayv2CallerSession) AllowedAllPaymentTokens() (bool, error) {
	return _Marketplacegatewayv2.Contract.AllowedAllPaymentTokens(&_Marketplacegatewayv2.CallOpts)
}

// AllowedPaymentToken is a free data retrieval call binding the contract method 0x8a095a57.
//
// Solidity: function allowedPaymentToken(address _token) view returns(bool)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Caller) AllowedPaymentToken(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _Marketplacegatewayv2.contract.Call(opts, &out, "allowedPaymentToken", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedPaymentToken is a free data retrieval call binding the contract method 0x8a095a57.
//
// Solidity: function allowedPaymentToken(address _token) view returns(bool)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) AllowedPaymentToken(_token common.Address) (bool, error) {
	return _Marketplacegatewayv2.Contract.AllowedPaymentToken(&_Marketplacegatewayv2.CallOpts, _token)
}

// AllowedPaymentToken is a free data retrieval call binding the contract method 0x8a095a57.
//
// Solidity: function allowedPaymentToken(address _token) view returns(bool)
func (_Marketplacegatewayv2 *Marketplacegatewayv2CallerSession) AllowedPaymentToken(_token common.Address) (bool, error) {
	return _Marketplacegatewayv2.Contract.AllowedPaymentToken(&_Marketplacegatewayv2.CallOpts, _token)
}

// GetInterface is a free data retrieval call binding the contract method 0xd8503fee.
//
// Solidity: function getInterface(string _interface) view returns(address)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Caller) GetInterface(opts *bind.CallOpts, _interface string) (common.Address, error) {
	var out []interface{}
	err := _Marketplacegatewayv2.contract.Call(opts, &out, "getInterface", _interface)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInterface is a free data retrieval call binding the contract method 0xd8503fee.
//
// Solidity: function getInterface(string _interface) view returns(address)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) GetInterface(_interface string) (common.Address, error) {
	return _Marketplacegatewayv2.Contract.GetInterface(&_Marketplacegatewayv2.CallOpts, _interface)
}

// GetInterface is a free data retrieval call binding the contract method 0xd8503fee.
//
// Solidity: function getInterface(string _interface) view returns(address)
func (_Marketplacegatewayv2 *Marketplacegatewayv2CallerSession) GetInterface(_interface string) (common.Address, error) {
	return _Marketplacegatewayv2.Contract.GetInterface(&_Marketplacegatewayv2.CallOpts, _interface)
}

// GetReferralConfig is a free data retrieval call binding the contract method 0x1e58a872.
//
// Solidity: function getReferralConfig() view returns(bool referralFeature_, address referralContract_, bool autoTransferReferralReward_)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Caller) GetReferralConfig(opts *bind.CallOpts) (struct {
	ReferralFeature            bool
	ReferralContract           common.Address
	AutoTransferReferralReward bool
}, error) {
	var out []interface{}
	err := _Marketplacegatewayv2.contract.Call(opts, &out, "getReferralConfig")

	outstruct := new(struct {
		ReferralFeature            bool
		ReferralContract           common.Address
		AutoTransferReferralReward bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReferralFeature = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ReferralContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AutoTransferReferralReward = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// GetReferralConfig is a free data retrieval call binding the contract method 0x1e58a872.
//
// Solidity: function getReferralConfig() view returns(bool referralFeature_, address referralContract_, bool autoTransferReferralReward_)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) GetReferralConfig() (struct {
	ReferralFeature            bool
	ReferralContract           common.Address
	AutoTransferReferralReward bool
}, error) {
	return _Marketplacegatewayv2.Contract.GetReferralConfig(&_Marketplacegatewayv2.CallOpts)
}

// GetReferralConfig is a free data retrieval call binding the contract method 0x1e58a872.
//
// Solidity: function getReferralConfig() view returns(bool referralFeature_, address referralContract_, bool autoTransferReferralReward_)
func (_Marketplacegatewayv2 *Marketplacegatewayv2CallerSession) GetReferralConfig() (struct {
	ReferralFeature            bool
	ReferralContract           common.Address
	AutoTransferReferralReward bool
}, error) {
	return _Marketplacegatewayv2.Contract.GetReferralConfig(&_Marketplacegatewayv2.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Marketplacegatewayv2.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Marketplacegatewayv2.Contract.GetRoleAdmin(&_Marketplacegatewayv2.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Marketplacegatewayv2 *Marketplacegatewayv2CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Marketplacegatewayv2.Contract.GetRoleAdmin(&_Marketplacegatewayv2.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Caller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Marketplacegatewayv2.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Marketplacegatewayv2.Contract.GetRoleMember(&_Marketplacegatewayv2.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Marketplacegatewayv2 *Marketplacegatewayv2CallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Marketplacegatewayv2.Contract.GetRoleMember(&_Marketplacegatewayv2.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Caller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Marketplacegatewayv2.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Marketplacegatewayv2.Contract.GetRoleMemberCount(&_Marketplacegatewayv2.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Marketplacegatewayv2 *Marketplacegatewayv2CallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Marketplacegatewayv2.Contract.GetRoleMemberCount(&_Marketplacegatewayv2.CallOpts, role)
}

// GetState is a free data retrieval call binding the contract method 0xf99fdd28.
//
// Solidity: function getState((uint8,address,uint256,uint256)[] _assets) view returns(uint256)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Caller) GetState(opts *bind.CallOpts, _assets []MarketAssetAsset) (*big.Int, error) {
	var out []interface{}
	err := _Marketplacegatewayv2.contract.Call(opts, &out, "getState", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetState is a free data retrieval call binding the contract method 0xf99fdd28.
//
// Solidity: function getState((uint8,address,uint256,uint256)[] _assets) view returns(uint256)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) GetState(_assets []MarketAssetAsset) (*big.Int, error) {
	return _Marketplacegatewayv2.Contract.GetState(&_Marketplacegatewayv2.CallOpts, _assets)
}

// GetState is a free data retrieval call binding the contract method 0xf99fdd28.
//
// Solidity: function getState((uint8,address,uint256,uint256)[] _assets) view returns(uint256)
func (_Marketplacegatewayv2 *Marketplacegatewayv2CallerSession) GetState(_assets []MarketAssetAsset) (*big.Int, error) {
	return _Marketplacegatewayv2.Contract.GetState(&_Marketplacegatewayv2.CallOpts, _assets)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Marketplacegatewayv2.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Marketplacegatewayv2.Contract.HasRole(&_Marketplacegatewayv2.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Marketplacegatewayv2 *Marketplacegatewayv2CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Marketplacegatewayv2.Contract.HasRole(&_Marketplacegatewayv2.CallOpts, role, account)
}

// MakerNonce is a free data retrieval call binding the contract method 0xcd370bb2.
//
// Solidity: function makerNonce(address _maker) view returns(uint256)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Caller) MakerNonce(opts *bind.CallOpts, _maker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Marketplacegatewayv2.contract.Call(opts, &out, "makerNonce", _maker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MakerNonce is a free data retrieval call binding the contract method 0xcd370bb2.
//
// Solidity: function makerNonce(address _maker) view returns(uint256)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) MakerNonce(_maker common.Address) (*big.Int, error) {
	return _Marketplacegatewayv2.Contract.MakerNonce(&_Marketplacegatewayv2.CallOpts, _maker)
}

// MakerNonce is a free data retrieval call binding the contract method 0xcd370bb2.
//
// Solidity: function makerNonce(address _maker) view returns(uint256)
func (_Marketplacegatewayv2 *Marketplacegatewayv2CallerSession) MakerNonce(_maker common.Address) (*big.Int, error) {
	return _Marketplacegatewayv2.Contract.MakerNonce(&_Marketplacegatewayv2.CallOpts, _maker)
}

// MinMarketFeePercentage is a free data retrieval call binding the contract method 0x3acb705b.
//
// Solidity: function minMarketFeePercentage() view returns(uint256)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Caller) MinMarketFeePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplacegatewayv2.contract.Call(opts, &out, "minMarketFeePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinMarketFeePercentage is a free data retrieval call binding the contract method 0x3acb705b.
//
// Solidity: function minMarketFeePercentage() view returns(uint256)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) MinMarketFeePercentage() (*big.Int, error) {
	return _Marketplacegatewayv2.Contract.MinMarketFeePercentage(&_Marketplacegatewayv2.CallOpts)
}

// MinMarketFeePercentage is a free data retrieval call binding the contract method 0x3acb705b.
//
// Solidity: function minMarketFeePercentage() view returns(uint256)
func (_Marketplacegatewayv2 *Marketplacegatewayv2CallerSession) MinMarketFeePercentage() (*big.Int, error) {
	return _Marketplacegatewayv2.Contract.MinMarketFeePercentage(&_Marketplacegatewayv2.CallOpts)
}

// OrderFinalized is a free data retrieval call binding the contract method 0xfe8082d9.
//
// Solidity: function orderFinalized(bytes32 _hash) view returns(bool)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Caller) OrderFinalized(opts *bind.CallOpts, _hash [32]byte) (bool, error) {
	var out []interface{}
	err := _Marketplacegatewayv2.contract.Call(opts, &out, "orderFinalized", _hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OrderFinalized is a free data retrieval call binding the contract method 0xfe8082d9.
//
// Solidity: function orderFinalized(bytes32 _hash) view returns(bool)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) OrderFinalized(_hash [32]byte) (bool, error) {
	return _Marketplacegatewayv2.Contract.OrderFinalized(&_Marketplacegatewayv2.CallOpts, _hash)
}

// OrderFinalized is a free data retrieval call binding the contract method 0xfe8082d9.
//
// Solidity: function orderFinalized(bytes32 _hash) view returns(bool)
func (_Marketplacegatewayv2 *Marketplacegatewayv2CallerSession) OrderFinalized(_hash [32]byte) (bool, error) {
	return _Marketplacegatewayv2.Contract.OrderFinalized(&_Marketplacegatewayv2.CallOpts, _hash)
}

// OrderValid is a free data retrieval call binding the contract method 0x93342c69.
//
// Solidity: function orderValid(bytes32 _hash, (address,uint8,(uint8,address,uint256,uint256)[],uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256) _order) view returns(bool)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Caller) OrderValid(opts *bind.CallOpts, _hash [32]byte, _order MarketOrderOrder) (bool, error) {
	var out []interface{}
	err := _Marketplacegatewayv2.contract.Call(opts, &out, "orderValid", _hash, _order)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OrderValid is a free data retrieval call binding the contract method 0x93342c69.
//
// Solidity: function orderValid(bytes32 _hash, (address,uint8,(uint8,address,uint256,uint256)[],uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256) _order) view returns(bool)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) OrderValid(_hash [32]byte, _order MarketOrderOrder) (bool, error) {
	return _Marketplacegatewayv2.Contract.OrderValid(&_Marketplacegatewayv2.CallOpts, _hash, _order)
}

// OrderValid is a free data retrieval call binding the contract method 0x93342c69.
//
// Solidity: function orderValid(bytes32 _hash, (address,uint8,(uint8,address,uint256,uint256)[],uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256) _order) view returns(bool)
func (_Marketplacegatewayv2 *Marketplacegatewayv2CallerSession) OrderValid(_hash [32]byte, _order MarketOrderOrder) (bool, error) {
	return _Marketplacegatewayv2.Contract.OrderValid(&_Marketplacegatewayv2.CallOpts, _hash, _order)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Marketplacegatewayv2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Marketplacegatewayv2.Contract.SupportsInterface(&_Marketplacegatewayv2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Marketplacegatewayv2 *Marketplacegatewayv2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Marketplacegatewayv2.Contract.SupportsInterface(&_Marketplacegatewayv2.CallOpts, interfaceId)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Caller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplacegatewayv2.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) Treasury() (common.Address, error) {
	return _Marketplacegatewayv2.Contract.Treasury(&_Marketplacegatewayv2.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Marketplacegatewayv2 *Marketplacegatewayv2CallerSession) Treasury() (common.Address, error) {
	return _Marketplacegatewayv2.Contract.Treasury(&_Marketplacegatewayv2.CallOpts)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xc9f7c467.
//
// Solidity: function cancelOrder((address,uint8,(uint8,address,uint256,uint256)[],uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256) _order) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Transactor) CancelOrder(opts *bind.TransactOpts, _order MarketOrderOrder) (*types.Transaction, error) {
	return _Marketplacegatewayv2.contract.Transact(opts, "cancelOrder", _order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xc9f7c467.
//
// Solidity: function cancelOrder((address,uint8,(uint8,address,uint256,uint256)[],uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256) _order) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) CancelOrder(_order MarketOrderOrder) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.CancelOrder(&_Marketplacegatewayv2.TransactOpts, _order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xc9f7c467.
//
// Solidity: function cancelOrder((address,uint8,(uint8,address,uint256,uint256)[],uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256) _order) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2TransactorSession) CancelOrder(_order MarketOrderOrder) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.CancelOrder(&_Marketplacegatewayv2.TransactOpts, _order)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.GrantRole(&_Marketplacegatewayv2.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.GrantRole(&_Marketplacegatewayv2.TransactOpts, role, account)
}

// IncreaseNonceMaker is a paid mutator transaction binding the contract method 0xe66ba89f.
//
// Solidity: function increaseNonceMaker() returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Transactor) IncreaseNonceMaker(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplacegatewayv2.contract.Transact(opts, "increaseNonceMaker")
}

// IncreaseNonceMaker is a paid mutator transaction binding the contract method 0xe66ba89f.
//
// Solidity: function increaseNonceMaker() returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) IncreaseNonceMaker() (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.IncreaseNonceMaker(&_Marketplacegatewayv2.TransactOpts)
}

// IncreaseNonceMaker is a paid mutator transaction binding the contract method 0xe66ba89f.
//
// Solidity: function increaseNonceMaker() returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2TransactorSession) IncreaseNonceMaker() (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.IncreaseNonceMaker(&_Marketplacegatewayv2.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x05ee9518.
//
// Solidity: function initialize(address _wronAddr, address _treasuryAddr, bool _allowedAllPaymentTokens, bool _referralFeature, address _referralContract, bool _autoTransferReferralReward, uint256 _minMarketFeePercentage, address[] _tokenPayments, string[] _interfaces, address[][2] _addresses) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Transactor) Initialize(opts *bind.TransactOpts, _wronAddr common.Address, _treasuryAddr common.Address, _allowedAllPaymentTokens bool, _referralFeature bool, _referralContract common.Address, _autoTransferReferralReward bool, _minMarketFeePercentage *big.Int, _tokenPayments []common.Address, _interfaces []string, _addresses [2][]common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.contract.Transact(opts, "initialize", _wronAddr, _treasuryAddr, _allowedAllPaymentTokens, _referralFeature, _referralContract, _autoTransferReferralReward, _minMarketFeePercentage, _tokenPayments, _interfaces, _addresses)
}

// Initialize is a paid mutator transaction binding the contract method 0x05ee9518.
//
// Solidity: function initialize(address _wronAddr, address _treasuryAddr, bool _allowedAllPaymentTokens, bool _referralFeature, address _referralContract, bool _autoTransferReferralReward, uint256 _minMarketFeePercentage, address[] _tokenPayments, string[] _interfaces, address[][2] _addresses) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) Initialize(_wronAddr common.Address, _treasuryAddr common.Address, _allowedAllPaymentTokens bool, _referralFeature bool, _referralContract common.Address, _autoTransferReferralReward bool, _minMarketFeePercentage *big.Int, _tokenPayments []common.Address, _interfaces []string, _addresses [2][]common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.Initialize(&_Marketplacegatewayv2.TransactOpts, _wronAddr, _treasuryAddr, _allowedAllPaymentTokens, _referralFeature, _referralContract, _autoTransferReferralReward, _minMarketFeePercentage, _tokenPayments, _interfaces, _addresses)
}

// Initialize is a paid mutator transaction binding the contract method 0x05ee9518.
//
// Solidity: function initialize(address _wronAddr, address _treasuryAddr, bool _allowedAllPaymentTokens, bool _referralFeature, address _referralContract, bool _autoTransferReferralReward, uint256 _minMarketFeePercentage, address[] _tokenPayments, string[] _interfaces, address[][2] _addresses) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2TransactorSession) Initialize(_wronAddr common.Address, _treasuryAddr common.Address, _allowedAllPaymentTokens bool, _referralFeature bool, _referralContract common.Address, _autoTransferReferralReward bool, _minMarketFeePercentage *big.Int, _tokenPayments []common.Address, _interfaces []string, _addresses [2][]common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.Initialize(&_Marketplacegatewayv2.TransactOpts, _wronAddr, _treasuryAddr, _allowedAllPaymentTokens, _referralFeature, _referralContract, _autoTransferReferralReward, _minMarketFeePercentage, _tokenPayments, _interfaces, _addresses)
}

// InteractWith is a paid mutator transaction binding the contract method 0x95a4ec00.
//
// Solidity: function interactWith(string _interface, bytes _data) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Transactor) InteractWith(opts *bind.TransactOpts, _interface string, _data []byte) (*types.Transaction, error) {
	return _Marketplacegatewayv2.contract.Transact(opts, "interactWith", _interface, _data)
}

// InteractWith is a paid mutator transaction binding the contract method 0x95a4ec00.
//
// Solidity: function interactWith(string _interface, bytes _data) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) InteractWith(_interface string, _data []byte) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.InteractWith(&_Marketplacegatewayv2.TransactOpts, _interface, _data)
}

// InteractWith is a paid mutator transaction binding the contract method 0x95a4ec00.
//
// Solidity: function interactWith(string _interface, bytes _data) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2TransactorSession) InteractWith(_interface string, _data []byte) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.InteractWith(&_Marketplacegatewayv2.TransactOpts, _interface, _data)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.RenounceRole(&_Marketplacegatewayv2.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.RenounceRole(&_Marketplacegatewayv2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.RevokeRole(&_Marketplacegatewayv2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.RevokeRole(&_Marketplacegatewayv2.TransactOpts, role, account)
}

// SetAllowedAllPaymentTokens is a paid mutator transaction binding the contract method 0x47dc7d28.
//
// Solidity: function setAllowedAllPaymentTokens(bool _flag) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Transactor) SetAllowedAllPaymentTokens(opts *bind.TransactOpts, _flag bool) (*types.Transaction, error) {
	return _Marketplacegatewayv2.contract.Transact(opts, "setAllowedAllPaymentTokens", _flag)
}

// SetAllowedAllPaymentTokens is a paid mutator transaction binding the contract method 0x47dc7d28.
//
// Solidity: function setAllowedAllPaymentTokens(bool _flag) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) SetAllowedAllPaymentTokens(_flag bool) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.SetAllowedAllPaymentTokens(&_Marketplacegatewayv2.TransactOpts, _flag)
}

// SetAllowedAllPaymentTokens is a paid mutator transaction binding the contract method 0x47dc7d28.
//
// Solidity: function setAllowedAllPaymentTokens(bool _flag) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2TransactorSession) SetAllowedAllPaymentTokens(_flag bool) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.SetAllowedAllPaymentTokens(&_Marketplacegatewayv2.TransactOpts, _flag)
}

// SetInterfaces is a paid mutator transaction binding the contract method 0x22d57b49.
//
// Solidity: function setInterfaces(string[] _interfaces, address[] _addresses) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Transactor) SetInterfaces(opts *bind.TransactOpts, _interfaces []string, _addresses []common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.contract.Transact(opts, "setInterfaces", _interfaces, _addresses)
}

// SetInterfaces is a paid mutator transaction binding the contract method 0x22d57b49.
//
// Solidity: function setInterfaces(string[] _interfaces, address[] _addresses) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) SetInterfaces(_interfaces []string, _addresses []common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.SetInterfaces(&_Marketplacegatewayv2.TransactOpts, _interfaces, _addresses)
}

// SetInterfaces is a paid mutator transaction binding the contract method 0x22d57b49.
//
// Solidity: function setInterfaces(string[] _interfaces, address[] _addresses) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2TransactorSession) SetInterfaces(_interfaces []string, _addresses []common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.SetInterfaces(&_Marketplacegatewayv2.TransactOpts, _interfaces, _addresses)
}

// SetKatanaRouter is a paid mutator transaction binding the contract method 0x8d2d4d2b.
//
// Solidity: function setKatanaRouter(address _routerContract) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Transactor) SetKatanaRouter(opts *bind.TransactOpts, _routerContract common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.contract.Transact(opts, "setKatanaRouter", _routerContract)
}

// SetKatanaRouter is a paid mutator transaction binding the contract method 0x8d2d4d2b.
//
// Solidity: function setKatanaRouter(address _routerContract) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) SetKatanaRouter(_routerContract common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.SetKatanaRouter(&_Marketplacegatewayv2.TransactOpts, _routerContract)
}

// SetKatanaRouter is a paid mutator transaction binding the contract method 0x8d2d4d2b.
//
// Solidity: function setKatanaRouter(address _routerContract) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2TransactorSession) SetKatanaRouter(_routerContract common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.SetKatanaRouter(&_Marketplacegatewayv2.TransactOpts, _routerContract)
}

// SetMinMarketFeePercentage is a paid mutator transaction binding the contract method 0x2156dd09.
//
// Solidity: function setMinMarketFeePercentage(uint256 _minMarketFeePercentage) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Transactor) SetMinMarketFeePercentage(opts *bind.TransactOpts, _minMarketFeePercentage *big.Int) (*types.Transaction, error) {
	return _Marketplacegatewayv2.contract.Transact(opts, "setMinMarketFeePercentage", _minMarketFeePercentage)
}

// SetMinMarketFeePercentage is a paid mutator transaction binding the contract method 0x2156dd09.
//
// Solidity: function setMinMarketFeePercentage(uint256 _minMarketFeePercentage) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) SetMinMarketFeePercentage(_minMarketFeePercentage *big.Int) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.SetMinMarketFeePercentage(&_Marketplacegatewayv2.TransactOpts, _minMarketFeePercentage)
}

// SetMinMarketFeePercentage is a paid mutator transaction binding the contract method 0x2156dd09.
//
// Solidity: function setMinMarketFeePercentage(uint256 _minMarketFeePercentage) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2TransactorSession) SetMinMarketFeePercentage(_minMarketFeePercentage *big.Int) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.SetMinMarketFeePercentage(&_Marketplacegatewayv2.TransactOpts, _minMarketFeePercentage)
}

// SetPaymentTokens is a paid mutator transaction binding the contract method 0xf7e3870c.
//
// Solidity: function setPaymentTokens(address[] _tokens, bool _allowed) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Transactor) SetPaymentTokens(opts *bind.TransactOpts, _tokens []common.Address, _allowed bool) (*types.Transaction, error) {
	return _Marketplacegatewayv2.contract.Transact(opts, "setPaymentTokens", _tokens, _allowed)
}

// SetPaymentTokens is a paid mutator transaction binding the contract method 0xf7e3870c.
//
// Solidity: function setPaymentTokens(address[] _tokens, bool _allowed) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) SetPaymentTokens(_tokens []common.Address, _allowed bool) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.SetPaymentTokens(&_Marketplacegatewayv2.TransactOpts, _tokens, _allowed)
}

// SetPaymentTokens is a paid mutator transaction binding the contract method 0xf7e3870c.
//
// Solidity: function setPaymentTokens(address[] _tokens, bool _allowed) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2TransactorSession) SetPaymentTokens(_tokens []common.Address, _allowed bool) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.SetPaymentTokens(&_Marketplacegatewayv2.TransactOpts, _tokens, _allowed)
}

// SetReferralConfig is a paid mutator transaction binding the contract method 0xcf2fe908.
//
// Solidity: function setReferralConfig(bool referralFeature_, address referralContract_, bool autoTransferReferralReward_) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Transactor) SetReferralConfig(opts *bind.TransactOpts, referralFeature_ bool, referralContract_ common.Address, autoTransferReferralReward_ bool) (*types.Transaction, error) {
	return _Marketplacegatewayv2.contract.Transact(opts, "setReferralConfig", referralFeature_, referralContract_, autoTransferReferralReward_)
}

// SetReferralConfig is a paid mutator transaction binding the contract method 0xcf2fe908.
//
// Solidity: function setReferralConfig(bool referralFeature_, address referralContract_, bool autoTransferReferralReward_) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) SetReferralConfig(referralFeature_ bool, referralContract_ common.Address, autoTransferReferralReward_ bool) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.SetReferralConfig(&_Marketplacegatewayv2.TransactOpts, referralFeature_, referralContract_, autoTransferReferralReward_)
}

// SetReferralConfig is a paid mutator transaction binding the contract method 0xcf2fe908.
//
// Solidity: function setReferralConfig(bool referralFeature_, address referralContract_, bool autoTransferReferralReward_) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2TransactorSession) SetReferralConfig(referralFeature_ bool, referralContract_ common.Address, autoTransferReferralReward_ bool) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.SetReferralConfig(&_Marketplacegatewayv2.TransactOpts, referralFeature_, referralContract_, autoTransferReferralReward_)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Transactor) SetTreasury(opts *bind.TransactOpts, _treasury common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.contract.Transact(opts, "setTreasury", _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.SetTreasury(&_Marketplacegatewayv2.TransactOpts, _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2TransactorSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.SetTreasury(&_Marketplacegatewayv2.TransactOpts, _treasury)
}

// SetWRON is a paid mutator transaction binding the contract method 0x6fb05fdf.
//
// Solidity: function setWRON(address _wronAddr) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Transactor) SetWRON(opts *bind.TransactOpts, _wronAddr common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.contract.Transact(opts, "setWRON", _wronAddr)
}

// SetWRON is a paid mutator transaction binding the contract method 0x6fb05fdf.
//
// Solidity: function setWRON(address _wronAddr) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) SetWRON(_wronAddr common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.SetWRON(&_Marketplacegatewayv2.TransactOpts, _wronAddr)
}

// SetWRON is a paid mutator transaction binding the contract method 0x6fb05fdf.
//
// Solidity: function setWRON(address _wronAddr) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2TransactorSession) SetWRON(_wronAddr common.Address) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.SetWRON(&_Marketplacegatewayv2.TransactOpts, _wronAddr)
}

// SettleOrder is a paid mutator transaction binding the contract method 0xc54c66d3.
//
// Solidity: function settleOrder(uint256 _expectedState, uint256 _settlePrice, address _referralAddr, bytes _signature, (address,uint8,(uint8,address,uint256,uint256)[],uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256) _order) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Transactor) SettleOrder(opts *bind.TransactOpts, _expectedState *big.Int, _settlePrice *big.Int, _referralAddr common.Address, _signature []byte, _order MarketOrderOrder) (*types.Transaction, error) {
	return _Marketplacegatewayv2.contract.Transact(opts, "settleOrder", _expectedState, _settlePrice, _referralAddr, _signature, _order)
}

// SettleOrder is a paid mutator transaction binding the contract method 0xc54c66d3.
//
// Solidity: function settleOrder(uint256 _expectedState, uint256 _settlePrice, address _referralAddr, bytes _signature, (address,uint8,(uint8,address,uint256,uint256)[],uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256) _order) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) SettleOrder(_expectedState *big.Int, _settlePrice *big.Int, _referralAddr common.Address, _signature []byte, _order MarketOrderOrder) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.SettleOrder(&_Marketplacegatewayv2.TransactOpts, _expectedState, _settlePrice, _referralAddr, _signature, _order)
}

// SettleOrder is a paid mutator transaction binding the contract method 0xc54c66d3.
//
// Solidity: function settleOrder(uint256 _expectedState, uint256 _settlePrice, address _referralAddr, bytes _signature, (address,uint8,(uint8,address,uint256,uint256)[],uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256) _order) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2TransactorSession) SettleOrder(_expectedState *big.Int, _settlePrice *big.Int, _referralAddr common.Address, _signature []byte, _order MarketOrderOrder) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.SettleOrder(&_Marketplacegatewayv2.TransactOpts, _expectedState, _settlePrice, _referralAddr, _signature, _order)
}

// SwapTokensAndSettleOrder is a paid mutator transaction binding the contract method 0x77a1de8b.
//
// Solidity: function swapTokensAndSettleOrder(uint256 _expectedState, uint256 _settlePrice, address _referralAddr, uint256 _deadline, address[] _path, bytes _signature, (address,uint8,(uint8,address,uint256,uint256)[],uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256) _order) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Transactor) SwapTokensAndSettleOrder(opts *bind.TransactOpts, _expectedState *big.Int, _settlePrice *big.Int, _referralAddr common.Address, _deadline *big.Int, _path []common.Address, _signature []byte, _order MarketOrderOrder) (*types.Transaction, error) {
	return _Marketplacegatewayv2.contract.Transact(opts, "swapTokensAndSettleOrder", _expectedState, _settlePrice, _referralAddr, _deadline, _path, _signature, _order)
}

// SwapTokensAndSettleOrder is a paid mutator transaction binding the contract method 0x77a1de8b.
//
// Solidity: function swapTokensAndSettleOrder(uint256 _expectedState, uint256 _settlePrice, address _referralAddr, uint256 _deadline, address[] _path, bytes _signature, (address,uint8,(uint8,address,uint256,uint256)[],uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256) _order) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) SwapTokensAndSettleOrder(_expectedState *big.Int, _settlePrice *big.Int, _referralAddr common.Address, _deadline *big.Int, _path []common.Address, _signature []byte, _order MarketOrderOrder) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.SwapTokensAndSettleOrder(&_Marketplacegatewayv2.TransactOpts, _expectedState, _settlePrice, _referralAddr, _deadline, _path, _signature, _order)
}

// SwapTokensAndSettleOrder is a paid mutator transaction binding the contract method 0x77a1de8b.
//
// Solidity: function swapTokensAndSettleOrder(uint256 _expectedState, uint256 _settlePrice, address _referralAddr, uint256 _deadline, address[] _path, bytes _signature, (address,uint8,(uint8,address,uint256,uint256)[],uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256) _order) returns()
func (_Marketplacegatewayv2 *Marketplacegatewayv2TransactorSession) SwapTokensAndSettleOrder(_expectedState *big.Int, _settlePrice *big.Int, _referralAddr common.Address, _deadline *big.Int, _path []common.Address, _signature []byte, _order MarketOrderOrder) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.SwapTokensAndSettleOrder(&_Marketplacegatewayv2.TransactOpts, _expectedState, _settlePrice, _referralAddr, _deadline, _path, _signature, _order)
}

// TryBulkCancelOrderByHash is a paid mutator transaction binding the contract method 0xcd55e01e.
//
// Solidity: function tryBulkCancelOrderByHash(bytes32[] _hashList) returns(bool[] _orderAlreadyFinalized)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Transactor) TryBulkCancelOrderByHash(opts *bind.TransactOpts, _hashList [][32]byte) (*types.Transaction, error) {
	return _Marketplacegatewayv2.contract.Transact(opts, "tryBulkCancelOrderByHash", _hashList)
}

// TryBulkCancelOrderByHash is a paid mutator transaction binding the contract method 0xcd55e01e.
//
// Solidity: function tryBulkCancelOrderByHash(bytes32[] _hashList) returns(bool[] _orderAlreadyFinalized)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Session) TryBulkCancelOrderByHash(_hashList [][32]byte) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.TryBulkCancelOrderByHash(&_Marketplacegatewayv2.TransactOpts, _hashList)
}

// TryBulkCancelOrderByHash is a paid mutator transaction binding the contract method 0xcd55e01e.
//
// Solidity: function tryBulkCancelOrderByHash(bytes32[] _hashList) returns(bool[] _orderAlreadyFinalized)
func (_Marketplacegatewayv2 *Marketplacegatewayv2TransactorSession) TryBulkCancelOrderByHash(_hashList [][32]byte) (*types.Transaction, error) {
	return _Marketplacegatewayv2.Contract.TryBulkCancelOrderByHash(&_Marketplacegatewayv2.TransactOpts, _hashList)
}

// Marketplacegatewayv2AllowedAllPaymentTokensIterator is returned from FilterAllowedAllPaymentTokens and is used to iterate over the raw logs and unpacked data for AllowedAllPaymentTokens events raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2AllowedAllPaymentTokensIterator struct {
	Event *Marketplacegatewayv2AllowedAllPaymentTokens // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Marketplacegatewayv2AllowedAllPaymentTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Marketplacegatewayv2AllowedAllPaymentTokens)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Marketplacegatewayv2AllowedAllPaymentTokens)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Marketplacegatewayv2AllowedAllPaymentTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Marketplacegatewayv2AllowedAllPaymentTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Marketplacegatewayv2AllowedAllPaymentTokens represents a AllowedAllPaymentTokens event raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2AllowedAllPaymentTokens struct {
	Arg0 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAllowedAllPaymentTokens is a free log retrieval operation binding the contract event 0xe92fe5c07ca76c17786aacc94ab29d26d505cd07ec755c8f8595944587e41588.
//
// Solidity: event AllowedAllPaymentTokens(bool arg0)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) FilterAllowedAllPaymentTokens(opts *bind.FilterOpts) (*Marketplacegatewayv2AllowedAllPaymentTokensIterator, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.FilterLogs(opts, "AllowedAllPaymentTokens")
	if err != nil {
		return nil, err
	}
	return &Marketplacegatewayv2AllowedAllPaymentTokensIterator{contract: _Marketplacegatewayv2.contract, event: "AllowedAllPaymentTokens", logs: logs, sub: sub}, nil
}

// WatchAllowedAllPaymentTokens is a free log subscription operation binding the contract event 0xe92fe5c07ca76c17786aacc94ab29d26d505cd07ec755c8f8595944587e41588.
//
// Solidity: event AllowedAllPaymentTokens(bool arg0)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) WatchAllowedAllPaymentTokens(opts *bind.WatchOpts, sink chan<- *Marketplacegatewayv2AllowedAllPaymentTokens) (event.Subscription, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.WatchLogs(opts, "AllowedAllPaymentTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Marketplacegatewayv2AllowedAllPaymentTokens)
				if err := _Marketplacegatewayv2.contract.UnpackLog(event, "AllowedAllPaymentTokens", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAllowedAllPaymentTokens is a log parse operation binding the contract event 0xe92fe5c07ca76c17786aacc94ab29d26d505cd07ec755c8f8595944587e41588.
//
// Solidity: event AllowedAllPaymentTokens(bool arg0)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) ParseAllowedAllPaymentTokens(log types.Log) (*Marketplacegatewayv2AllowedAllPaymentTokens, error) {
	event := new(Marketplacegatewayv2AllowedAllPaymentTokens)
	if err := _Marketplacegatewayv2.contract.UnpackLog(event, "AllowedAllPaymentTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Marketplacegatewayv2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2InitializedIterator struct {
	Event *Marketplacegatewayv2Initialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Marketplacegatewayv2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Marketplacegatewayv2Initialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Marketplacegatewayv2Initialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Marketplacegatewayv2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Marketplacegatewayv2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Marketplacegatewayv2Initialized represents a Initialized event raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) FilterInitialized(opts *bind.FilterOpts) (*Marketplacegatewayv2InitializedIterator, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Marketplacegatewayv2InitializedIterator{contract: _Marketplacegatewayv2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Marketplacegatewayv2Initialized) (event.Subscription, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Marketplacegatewayv2Initialized)
				if err := _Marketplacegatewayv2.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) ParseInitialized(log types.Log) (*Marketplacegatewayv2Initialized, error) {
	event := new(Marketplacegatewayv2Initialized)
	if err := _Marketplacegatewayv2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Marketplacegatewayv2InterfacesUpdatedIterator is returned from FilterInterfacesUpdated and is used to iterate over the raw logs and unpacked data for InterfacesUpdated events raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2InterfacesUpdatedIterator struct {
	Event *Marketplacegatewayv2InterfacesUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Marketplacegatewayv2InterfacesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Marketplacegatewayv2InterfacesUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Marketplacegatewayv2InterfacesUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Marketplacegatewayv2InterfacesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Marketplacegatewayv2InterfacesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Marketplacegatewayv2InterfacesUpdated represents a InterfacesUpdated event raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2InterfacesUpdated struct {
	Interfaces []string
	Arg1       []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInterfacesUpdated is a free log retrieval operation binding the contract event 0x89c7fdf0dd4fd92de69739c751056ed471a48cad8266f623d1ae546ddb962a18.
//
// Solidity: event InterfacesUpdated(string[] _interfaces, address[] arg1)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) FilterInterfacesUpdated(opts *bind.FilterOpts) (*Marketplacegatewayv2InterfacesUpdatedIterator, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.FilterLogs(opts, "InterfacesUpdated")
	if err != nil {
		return nil, err
	}
	return &Marketplacegatewayv2InterfacesUpdatedIterator{contract: _Marketplacegatewayv2.contract, event: "InterfacesUpdated", logs: logs, sub: sub}, nil
}

// WatchInterfacesUpdated is a free log subscription operation binding the contract event 0x89c7fdf0dd4fd92de69739c751056ed471a48cad8266f623d1ae546ddb962a18.
//
// Solidity: event InterfacesUpdated(string[] _interfaces, address[] arg1)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) WatchInterfacesUpdated(opts *bind.WatchOpts, sink chan<- *Marketplacegatewayv2InterfacesUpdated) (event.Subscription, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.WatchLogs(opts, "InterfacesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Marketplacegatewayv2InterfacesUpdated)
				if err := _Marketplacegatewayv2.contract.UnpackLog(event, "InterfacesUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInterfacesUpdated is a log parse operation binding the contract event 0x89c7fdf0dd4fd92de69739c751056ed471a48cad8266f623d1ae546ddb962a18.
//
// Solidity: event InterfacesUpdated(string[] _interfaces, address[] arg1)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) ParseInterfacesUpdated(log types.Log) (*Marketplacegatewayv2InterfacesUpdated, error) {
	event := new(Marketplacegatewayv2InterfacesUpdated)
	if err := _Marketplacegatewayv2.contract.UnpackLog(event, "InterfacesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Marketplacegatewayv2KatanaRouterUpdatedIterator is returned from FilterKatanaRouterUpdated and is used to iterate over the raw logs and unpacked data for KatanaRouterUpdated events raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2KatanaRouterUpdatedIterator struct {
	Event *Marketplacegatewayv2KatanaRouterUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Marketplacegatewayv2KatanaRouterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Marketplacegatewayv2KatanaRouterUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Marketplacegatewayv2KatanaRouterUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Marketplacegatewayv2KatanaRouterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Marketplacegatewayv2KatanaRouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Marketplacegatewayv2KatanaRouterUpdated represents a KatanaRouterUpdated event raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2KatanaRouterUpdated struct {
	KatanaRouterContract common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterKatanaRouterUpdated is a free log retrieval operation binding the contract event 0x1277ddb88d00e054b2831a04b75bda24dffee5621cecbb7c52405b89cc358174.
//
// Solidity: event KatanaRouterUpdated(address KatanaRouterContract)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) FilterKatanaRouterUpdated(opts *bind.FilterOpts) (*Marketplacegatewayv2KatanaRouterUpdatedIterator, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.FilterLogs(opts, "KatanaRouterUpdated")
	if err != nil {
		return nil, err
	}
	return &Marketplacegatewayv2KatanaRouterUpdatedIterator{contract: _Marketplacegatewayv2.contract, event: "KatanaRouterUpdated", logs: logs, sub: sub}, nil
}

// WatchKatanaRouterUpdated is a free log subscription operation binding the contract event 0x1277ddb88d00e054b2831a04b75bda24dffee5621cecbb7c52405b89cc358174.
//
// Solidity: event KatanaRouterUpdated(address KatanaRouterContract)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) WatchKatanaRouterUpdated(opts *bind.WatchOpts, sink chan<- *Marketplacegatewayv2KatanaRouterUpdated) (event.Subscription, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.WatchLogs(opts, "KatanaRouterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Marketplacegatewayv2KatanaRouterUpdated)
				if err := _Marketplacegatewayv2.contract.UnpackLog(event, "KatanaRouterUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseKatanaRouterUpdated is a log parse operation binding the contract event 0x1277ddb88d00e054b2831a04b75bda24dffee5621cecbb7c52405b89cc358174.
//
// Solidity: event KatanaRouterUpdated(address KatanaRouterContract)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) ParseKatanaRouterUpdated(log types.Log) (*Marketplacegatewayv2KatanaRouterUpdated, error) {
	event := new(Marketplacegatewayv2KatanaRouterUpdated)
	if err := _Marketplacegatewayv2.contract.UnpackLog(event, "KatanaRouterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Marketplacegatewayv2MakerNonceUpdatedIterator is returned from FilterMakerNonceUpdated and is used to iterate over the raw logs and unpacked data for MakerNonceUpdated events raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2MakerNonceUpdatedIterator struct {
	Event *Marketplacegatewayv2MakerNonceUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Marketplacegatewayv2MakerNonceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Marketplacegatewayv2MakerNonceUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Marketplacegatewayv2MakerNonceUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Marketplacegatewayv2MakerNonceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Marketplacegatewayv2MakerNonceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Marketplacegatewayv2MakerNonceUpdated represents a MakerNonceUpdated event raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2MakerNonceUpdated struct {
	Maker common.Address
	Nonce *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMakerNonceUpdated is a free log retrieval operation binding the contract event 0x7b341b46777282de7f0b54a96a8543a00ab3256e5006760164438e610e566a28.
//
// Solidity: event MakerNonceUpdated(address maker, uint256 nonce)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) FilterMakerNonceUpdated(opts *bind.FilterOpts) (*Marketplacegatewayv2MakerNonceUpdatedIterator, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.FilterLogs(opts, "MakerNonceUpdated")
	if err != nil {
		return nil, err
	}
	return &Marketplacegatewayv2MakerNonceUpdatedIterator{contract: _Marketplacegatewayv2.contract, event: "MakerNonceUpdated", logs: logs, sub: sub}, nil
}

// WatchMakerNonceUpdated is a free log subscription operation binding the contract event 0x7b341b46777282de7f0b54a96a8543a00ab3256e5006760164438e610e566a28.
//
// Solidity: event MakerNonceUpdated(address maker, uint256 nonce)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) WatchMakerNonceUpdated(opts *bind.WatchOpts, sink chan<- *Marketplacegatewayv2MakerNonceUpdated) (event.Subscription, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.WatchLogs(opts, "MakerNonceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Marketplacegatewayv2MakerNonceUpdated)
				if err := _Marketplacegatewayv2.contract.UnpackLog(event, "MakerNonceUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMakerNonceUpdated is a log parse operation binding the contract event 0x7b341b46777282de7f0b54a96a8543a00ab3256e5006760164438e610e566a28.
//
// Solidity: event MakerNonceUpdated(address maker, uint256 nonce)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) ParseMakerNonceUpdated(log types.Log) (*Marketplacegatewayv2MakerNonceUpdated, error) {
	event := new(Marketplacegatewayv2MakerNonceUpdated)
	if err := _Marketplacegatewayv2.contract.UnpackLog(event, "MakerNonceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Marketplacegatewayv2MinMarketFeePercentageUpdatedIterator is returned from FilterMinMarketFeePercentageUpdated and is used to iterate over the raw logs and unpacked data for MinMarketFeePercentageUpdated events raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2MinMarketFeePercentageUpdatedIterator struct {
	Event *Marketplacegatewayv2MinMarketFeePercentageUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Marketplacegatewayv2MinMarketFeePercentageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Marketplacegatewayv2MinMarketFeePercentageUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Marketplacegatewayv2MinMarketFeePercentageUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Marketplacegatewayv2MinMarketFeePercentageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Marketplacegatewayv2MinMarketFeePercentageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Marketplacegatewayv2MinMarketFeePercentageUpdated represents a MinMarketFeePercentageUpdated event raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2MinMarketFeePercentageUpdated struct {
	MinMarketFeePercentage *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterMinMarketFeePercentageUpdated is a free log retrieval operation binding the contract event 0xd53de4b06b8357354b91283dc549b54c8953899f1173e89814e6dad281940a80.
//
// Solidity: event MinMarketFeePercentageUpdated(uint256 minMarketFeePercentage)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) FilterMinMarketFeePercentageUpdated(opts *bind.FilterOpts) (*Marketplacegatewayv2MinMarketFeePercentageUpdatedIterator, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.FilterLogs(opts, "MinMarketFeePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return &Marketplacegatewayv2MinMarketFeePercentageUpdatedIterator{contract: _Marketplacegatewayv2.contract, event: "MinMarketFeePercentageUpdated", logs: logs, sub: sub}, nil
}

// WatchMinMarketFeePercentageUpdated is a free log subscription operation binding the contract event 0xd53de4b06b8357354b91283dc549b54c8953899f1173e89814e6dad281940a80.
//
// Solidity: event MinMarketFeePercentageUpdated(uint256 minMarketFeePercentage)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) WatchMinMarketFeePercentageUpdated(opts *bind.WatchOpts, sink chan<- *Marketplacegatewayv2MinMarketFeePercentageUpdated) (event.Subscription, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.WatchLogs(opts, "MinMarketFeePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Marketplacegatewayv2MinMarketFeePercentageUpdated)
				if err := _Marketplacegatewayv2.contract.UnpackLog(event, "MinMarketFeePercentageUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinMarketFeePercentageUpdated is a log parse operation binding the contract event 0xd53de4b06b8357354b91283dc549b54c8953899f1173e89814e6dad281940a80.
//
// Solidity: event MinMarketFeePercentageUpdated(uint256 minMarketFeePercentage)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) ParseMinMarketFeePercentageUpdated(log types.Log) (*Marketplacegatewayv2MinMarketFeePercentageUpdated, error) {
	event := new(Marketplacegatewayv2MinMarketFeePercentageUpdated)
	if err := _Marketplacegatewayv2.contract.UnpackLog(event, "MinMarketFeePercentageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Marketplacegatewayv2OrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2OrderCancelledIterator struct {
	Event *Marketplacegatewayv2OrderCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Marketplacegatewayv2OrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Marketplacegatewayv2OrderCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Marketplacegatewayv2OrderCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Marketplacegatewayv2OrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Marketplacegatewayv2OrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Marketplacegatewayv2OrderCancelled represents a OrderCancelled event raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2OrderCancelled struct {
	Arg0      [32]byte
	Requester common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0xa6eb7cdc219e1518ced964e9a34e61d68a94e4f1569db3e84256ba981ba52753.
//
// Solidity: event OrderCancelled(bytes32 arg0, address requester)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) FilterOrderCancelled(opts *bind.FilterOpts) (*Marketplacegatewayv2OrderCancelledIterator, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.FilterLogs(opts, "OrderCancelled")
	if err != nil {
		return nil, err
	}
	return &Marketplacegatewayv2OrderCancelledIterator{contract: _Marketplacegatewayv2.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0xa6eb7cdc219e1518ced964e9a34e61d68a94e4f1569db3e84256ba981ba52753.
//
// Solidity: event OrderCancelled(bytes32 arg0, address requester)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *Marketplacegatewayv2OrderCancelled) (event.Subscription, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.WatchLogs(opts, "OrderCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Marketplacegatewayv2OrderCancelled)
				if err := _Marketplacegatewayv2.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrderCancelled is a log parse operation binding the contract event 0xa6eb7cdc219e1518ced964e9a34e61d68a94e4f1569db3e84256ba981ba52753.
//
// Solidity: event OrderCancelled(bytes32 arg0, address requester)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) ParseOrderCancelled(log types.Log) (*Marketplacegatewayv2OrderCancelled, error) {
	event := new(Marketplacegatewayv2OrderCancelled)
	if err := _Marketplacegatewayv2.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Marketplacegatewayv2OrderMatchedIterator is returned from FilterOrderMatched and is used to iterate over the raw logs and unpacked data for OrderMatched events raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2OrderMatchedIterator struct {
	Event *Marketplacegatewayv2OrderMatched // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Marketplacegatewayv2OrderMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Marketplacegatewayv2OrderMatched)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Marketplacegatewayv2OrderMatched)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Marketplacegatewayv2OrderMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Marketplacegatewayv2OrderMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Marketplacegatewayv2OrderMatched represents a OrderMatched event raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2OrderMatched struct {
	Hash                [32]byte
	Maker               common.Address
	Matcher             common.Address
	Kind                uint8
	BidToken            common.Address
	BidPrice            *big.Int
	PaymentToken        common.Address
	SettlePrice         *big.Int
	SellerReceived      *big.Int
	MarketFeePercentage *big.Int
	MarketFeeTaken      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterOrderMatched is a free log retrieval operation binding the contract event 0xafa0d706792fa5d4e9aaf5e456e08e2a833b1e64a201710b782f29172f6d7a3a.
//
// Solidity: event OrderMatched(bytes32 hash, address maker, address matcher, uint8 kind, address bidToken, uint256 bidPrice, address paymentToken, uint256 settlePrice, uint256 sellerReceived, uint256 marketFeePercentage, uint256 marketFeeTaken)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) FilterOrderMatched(opts *bind.FilterOpts) (*Marketplacegatewayv2OrderMatchedIterator, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.FilterLogs(opts, "OrderMatched")
	if err != nil {
		return nil, err
	}
	return &Marketplacegatewayv2OrderMatchedIterator{contract: _Marketplacegatewayv2.contract, event: "OrderMatched", logs: logs, sub: sub}, nil
}

// WatchOrderMatched is a free log subscription operation binding the contract event 0xafa0d706792fa5d4e9aaf5e456e08e2a833b1e64a201710b782f29172f6d7a3a.
//
// Solidity: event OrderMatched(bytes32 hash, address maker, address matcher, uint8 kind, address bidToken, uint256 bidPrice, address paymentToken, uint256 settlePrice, uint256 sellerReceived, uint256 marketFeePercentage, uint256 marketFeeTaken)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) WatchOrderMatched(opts *bind.WatchOpts, sink chan<- *Marketplacegatewayv2OrderMatched) (event.Subscription, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.WatchLogs(opts, "OrderMatched")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Marketplacegatewayv2OrderMatched)
				if err := _Marketplacegatewayv2.contract.UnpackLog(event, "OrderMatched", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrderMatched is a log parse operation binding the contract event 0xafa0d706792fa5d4e9aaf5e456e08e2a833b1e64a201710b782f29172f6d7a3a.
//
// Solidity: event OrderMatched(bytes32 hash, address maker, address matcher, uint8 kind, address bidToken, uint256 bidPrice, address paymentToken, uint256 settlePrice, uint256 sellerReceived, uint256 marketFeePercentage, uint256 marketFeeTaken)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) ParseOrderMatched(log types.Log) (*Marketplacegatewayv2OrderMatched, error) {
	event := new(Marketplacegatewayv2OrderMatched)
	if err := _Marketplacegatewayv2.contract.UnpackLog(event, "OrderMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Marketplacegatewayv2PaymentTokensAllowedIterator is returned from FilterPaymentTokensAllowed and is used to iterate over the raw logs and unpacked data for PaymentTokensAllowed events raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2PaymentTokensAllowedIterator struct {
	Event *Marketplacegatewayv2PaymentTokensAllowed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Marketplacegatewayv2PaymentTokensAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Marketplacegatewayv2PaymentTokensAllowed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Marketplacegatewayv2PaymentTokensAllowed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Marketplacegatewayv2PaymentTokensAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Marketplacegatewayv2PaymentTokensAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Marketplacegatewayv2PaymentTokensAllowed represents a PaymentTokensAllowed event raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2PaymentTokensAllowed struct {
	Arg0 []common.Address
	Arg1 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPaymentTokensAllowed is a free log retrieval operation binding the contract event 0xd7a151a9547201814adf6087a8e2e3eadf6b09c7cd8bc583c11f58ed90f7e9f7.
//
// Solidity: event PaymentTokensAllowed(address[] arg0, bool arg1)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) FilterPaymentTokensAllowed(opts *bind.FilterOpts) (*Marketplacegatewayv2PaymentTokensAllowedIterator, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.FilterLogs(opts, "PaymentTokensAllowed")
	if err != nil {
		return nil, err
	}
	return &Marketplacegatewayv2PaymentTokensAllowedIterator{contract: _Marketplacegatewayv2.contract, event: "PaymentTokensAllowed", logs: logs, sub: sub}, nil
}

// WatchPaymentTokensAllowed is a free log subscription operation binding the contract event 0xd7a151a9547201814adf6087a8e2e3eadf6b09c7cd8bc583c11f58ed90f7e9f7.
//
// Solidity: event PaymentTokensAllowed(address[] arg0, bool arg1)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) WatchPaymentTokensAllowed(opts *bind.WatchOpts, sink chan<- *Marketplacegatewayv2PaymentTokensAllowed) (event.Subscription, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.WatchLogs(opts, "PaymentTokensAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Marketplacegatewayv2PaymentTokensAllowed)
				if err := _Marketplacegatewayv2.contract.UnpackLog(event, "PaymentTokensAllowed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaymentTokensAllowed is a log parse operation binding the contract event 0xd7a151a9547201814adf6087a8e2e3eadf6b09c7cd8bc583c11f58ed90f7e9f7.
//
// Solidity: event PaymentTokensAllowed(address[] arg0, bool arg1)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) ParsePaymentTokensAllowed(log types.Log) (*Marketplacegatewayv2PaymentTokensAllowed, error) {
	event := new(Marketplacegatewayv2PaymentTokensAllowed)
	if err := _Marketplacegatewayv2.contract.UnpackLog(event, "PaymentTokensAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Marketplacegatewayv2ReferralConfigUpdatedIterator is returned from FilterReferralConfigUpdated and is used to iterate over the raw logs and unpacked data for ReferralConfigUpdated events raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2ReferralConfigUpdatedIterator struct {
	Event *Marketplacegatewayv2ReferralConfigUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Marketplacegatewayv2ReferralConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Marketplacegatewayv2ReferralConfigUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Marketplacegatewayv2ReferralConfigUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Marketplacegatewayv2ReferralConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Marketplacegatewayv2ReferralConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Marketplacegatewayv2ReferralConfigUpdated represents a ReferralConfigUpdated event raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2ReferralConfigUpdated struct {
	ReferralFeature            bool
	ReferralContract           common.Address
	AutoTransferReferralReward bool
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterReferralConfigUpdated is a free log retrieval operation binding the contract event 0x0e08f6ed4dce4aa7304adf70dbcefee1dd50ba4076fe970b18287eb83242b0e1.
//
// Solidity: event ReferralConfigUpdated(bool referralFeature, address referralContract, bool autoTransferReferralReward)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) FilterReferralConfigUpdated(opts *bind.FilterOpts) (*Marketplacegatewayv2ReferralConfigUpdatedIterator, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.FilterLogs(opts, "ReferralConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &Marketplacegatewayv2ReferralConfigUpdatedIterator{contract: _Marketplacegatewayv2.contract, event: "ReferralConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchReferralConfigUpdated is a free log subscription operation binding the contract event 0x0e08f6ed4dce4aa7304adf70dbcefee1dd50ba4076fe970b18287eb83242b0e1.
//
// Solidity: event ReferralConfigUpdated(bool referralFeature, address referralContract, bool autoTransferReferralReward)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) WatchReferralConfigUpdated(opts *bind.WatchOpts, sink chan<- *Marketplacegatewayv2ReferralConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.WatchLogs(opts, "ReferralConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Marketplacegatewayv2ReferralConfigUpdated)
				if err := _Marketplacegatewayv2.contract.UnpackLog(event, "ReferralConfigUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReferralConfigUpdated is a log parse operation binding the contract event 0x0e08f6ed4dce4aa7304adf70dbcefee1dd50ba4076fe970b18287eb83242b0e1.
//
// Solidity: event ReferralConfigUpdated(bool referralFeature, address referralContract, bool autoTransferReferralReward)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) ParseReferralConfigUpdated(log types.Log) (*Marketplacegatewayv2ReferralConfigUpdated, error) {
	event := new(Marketplacegatewayv2ReferralConfigUpdated)
	if err := _Marketplacegatewayv2.contract.UnpackLog(event, "ReferralConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Marketplacegatewayv2ReferralRewardTransferredIterator is returned from FilterReferralRewardTransferred and is used to iterate over the raw logs and unpacked data for ReferralRewardTransferred events raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2ReferralRewardTransferredIterator struct {
	Event *Marketplacegatewayv2ReferralRewardTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Marketplacegatewayv2ReferralRewardTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Marketplacegatewayv2ReferralRewardTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Marketplacegatewayv2ReferralRewardTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Marketplacegatewayv2ReferralRewardTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Marketplacegatewayv2ReferralRewardTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Marketplacegatewayv2ReferralRewardTransferred represents a ReferralRewardTransferred event raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2ReferralRewardTransferred struct {
	Token        common.Address
	ReferralAddr common.Address
	Amount       *big.Int
	Receiver     common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterReferralRewardTransferred is a free log retrieval operation binding the contract event 0x9d2f142870db1003bfb87daca9acdbd2209cb4bc11143042196639c2e464d0b5.
//
// Solidity: event ReferralRewardTransferred(address token, address referralAddr, uint256 amount, address receiver)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) FilterReferralRewardTransferred(opts *bind.FilterOpts) (*Marketplacegatewayv2ReferralRewardTransferredIterator, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.FilterLogs(opts, "ReferralRewardTransferred")
	if err != nil {
		return nil, err
	}
	return &Marketplacegatewayv2ReferralRewardTransferredIterator{contract: _Marketplacegatewayv2.contract, event: "ReferralRewardTransferred", logs: logs, sub: sub}, nil
}

// WatchReferralRewardTransferred is a free log subscription operation binding the contract event 0x9d2f142870db1003bfb87daca9acdbd2209cb4bc11143042196639c2e464d0b5.
//
// Solidity: event ReferralRewardTransferred(address token, address referralAddr, uint256 amount, address receiver)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) WatchReferralRewardTransferred(opts *bind.WatchOpts, sink chan<- *Marketplacegatewayv2ReferralRewardTransferred) (event.Subscription, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.WatchLogs(opts, "ReferralRewardTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Marketplacegatewayv2ReferralRewardTransferred)
				if err := _Marketplacegatewayv2.contract.UnpackLog(event, "ReferralRewardTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReferralRewardTransferred is a log parse operation binding the contract event 0x9d2f142870db1003bfb87daca9acdbd2209cb4bc11143042196639c2e464d0b5.
//
// Solidity: event ReferralRewardTransferred(address token, address referralAddr, uint256 amount, address receiver)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) ParseReferralRewardTransferred(log types.Log) (*Marketplacegatewayv2ReferralRewardTransferred, error) {
	event := new(Marketplacegatewayv2ReferralRewardTransferred)
	if err := _Marketplacegatewayv2.contract.UnpackLog(event, "ReferralRewardTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Marketplacegatewayv2RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2RoleAdminChangedIterator struct {
	Event *Marketplacegatewayv2RoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Marketplacegatewayv2RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Marketplacegatewayv2RoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Marketplacegatewayv2RoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Marketplacegatewayv2RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Marketplacegatewayv2RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Marketplacegatewayv2RoleAdminChanged represents a RoleAdminChanged event raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*Marketplacegatewayv2RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Marketplacegatewayv2.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &Marketplacegatewayv2RoleAdminChangedIterator{contract: _Marketplacegatewayv2.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *Marketplacegatewayv2RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Marketplacegatewayv2.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Marketplacegatewayv2RoleAdminChanged)
				if err := _Marketplacegatewayv2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) ParseRoleAdminChanged(log types.Log) (*Marketplacegatewayv2RoleAdminChanged, error) {
	event := new(Marketplacegatewayv2RoleAdminChanged)
	if err := _Marketplacegatewayv2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Marketplacegatewayv2RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2RoleGrantedIterator struct {
	Event *Marketplacegatewayv2RoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Marketplacegatewayv2RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Marketplacegatewayv2RoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Marketplacegatewayv2RoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Marketplacegatewayv2RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Marketplacegatewayv2RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Marketplacegatewayv2RoleGranted represents a RoleGranted event raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Marketplacegatewayv2RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Marketplacegatewayv2.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Marketplacegatewayv2RoleGrantedIterator{contract: _Marketplacegatewayv2.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *Marketplacegatewayv2RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Marketplacegatewayv2.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Marketplacegatewayv2RoleGranted)
				if err := _Marketplacegatewayv2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) ParseRoleGranted(log types.Log) (*Marketplacegatewayv2RoleGranted, error) {
	event := new(Marketplacegatewayv2RoleGranted)
	if err := _Marketplacegatewayv2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Marketplacegatewayv2RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2RoleRevokedIterator struct {
	Event *Marketplacegatewayv2RoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Marketplacegatewayv2RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Marketplacegatewayv2RoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Marketplacegatewayv2RoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Marketplacegatewayv2RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Marketplacegatewayv2RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Marketplacegatewayv2RoleRevoked represents a RoleRevoked event raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Marketplacegatewayv2RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Marketplacegatewayv2.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Marketplacegatewayv2RoleRevokedIterator{contract: _Marketplacegatewayv2.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *Marketplacegatewayv2RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Marketplacegatewayv2.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Marketplacegatewayv2RoleRevoked)
				if err := _Marketplacegatewayv2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) ParseRoleRevoked(log types.Log) (*Marketplacegatewayv2RoleRevoked, error) {
	event := new(Marketplacegatewayv2RoleRevoked)
	if err := _Marketplacegatewayv2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Marketplacegatewayv2TreasuryUpdatedIterator is returned from FilterTreasuryUpdated and is used to iterate over the raw logs and unpacked data for TreasuryUpdated events raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2TreasuryUpdatedIterator struct {
	Event *Marketplacegatewayv2TreasuryUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Marketplacegatewayv2TreasuryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Marketplacegatewayv2TreasuryUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Marketplacegatewayv2TreasuryUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Marketplacegatewayv2TreasuryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Marketplacegatewayv2TreasuryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Marketplacegatewayv2TreasuryUpdated represents a TreasuryUpdated event raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2TreasuryUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTreasuryUpdated is a free log retrieval operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address arg0)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) FilterTreasuryUpdated(opts *bind.FilterOpts) (*Marketplacegatewayv2TreasuryUpdatedIterator, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.FilterLogs(opts, "TreasuryUpdated")
	if err != nil {
		return nil, err
	}
	return &Marketplacegatewayv2TreasuryUpdatedIterator{contract: _Marketplacegatewayv2.contract, event: "TreasuryUpdated", logs: logs, sub: sub}, nil
}

// WatchTreasuryUpdated is a free log subscription operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address arg0)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) WatchTreasuryUpdated(opts *bind.WatchOpts, sink chan<- *Marketplacegatewayv2TreasuryUpdated) (event.Subscription, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.WatchLogs(opts, "TreasuryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Marketplacegatewayv2TreasuryUpdated)
				if err := _Marketplacegatewayv2.contract.UnpackLog(event, "TreasuryUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTreasuryUpdated is a log parse operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address arg0)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) ParseTreasuryUpdated(log types.Log) (*Marketplacegatewayv2TreasuryUpdated, error) {
	event := new(Marketplacegatewayv2TreasuryUpdated)
	if err := _Marketplacegatewayv2.contract.UnpackLog(event, "TreasuryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Marketplacegatewayv2WRONUpdatedIterator is returned from FilterWRONUpdated and is used to iterate over the raw logs and unpacked data for WRONUpdated events raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2WRONUpdatedIterator struct {
	Event *Marketplacegatewayv2WRONUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Marketplacegatewayv2WRONUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Marketplacegatewayv2WRONUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Marketplacegatewayv2WRONUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Marketplacegatewayv2WRONUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Marketplacegatewayv2WRONUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Marketplacegatewayv2WRONUpdated represents a WRONUpdated event raised by the Marketplacegatewayv2 contract.
type Marketplacegatewayv2WRONUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWRONUpdated is a free log retrieval operation binding the contract event 0x1278746c7c2b06b6a4a9bff00c1e79353a94b44ffd6e6ec1e25c45764b63aba6.
//
// Solidity: event WRONUpdated(address arg0)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) FilterWRONUpdated(opts *bind.FilterOpts) (*Marketplacegatewayv2WRONUpdatedIterator, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.FilterLogs(opts, "WRONUpdated")
	if err != nil {
		return nil, err
	}
	return &Marketplacegatewayv2WRONUpdatedIterator{contract: _Marketplacegatewayv2.contract, event: "WRONUpdated", logs: logs, sub: sub}, nil
}

// WatchWRONUpdated is a free log subscription operation binding the contract event 0x1278746c7c2b06b6a4a9bff00c1e79353a94b44ffd6e6ec1e25c45764b63aba6.
//
// Solidity: event WRONUpdated(address arg0)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) WatchWRONUpdated(opts *bind.WatchOpts, sink chan<- *Marketplacegatewayv2WRONUpdated) (event.Subscription, error) {

	logs, sub, err := _Marketplacegatewayv2.contract.WatchLogs(opts, "WRONUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Marketplacegatewayv2WRONUpdated)
				if err := _Marketplacegatewayv2.contract.UnpackLog(event, "WRONUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWRONUpdated is a log parse operation binding the contract event 0x1278746c7c2b06b6a4a9bff00c1e79353a94b44ffd6e6ec1e25c45764b63aba6.
//
// Solidity: event WRONUpdated(address arg0)
func (_Marketplacegatewayv2 *Marketplacegatewayv2Filterer) ParseWRONUpdated(log types.Log) (*Marketplacegatewayv2WRONUpdated, error) {
	event := new(Marketplacegatewayv2WRONUpdated)
	if err := _Marketplacegatewayv2.contract.UnpackLog(event, "WRONUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
