package addresses

import goex "github.com/nntaoli-project/GoEx"

const (
	BITFINEX_BTC_IAN = "1JWKguEVBU48za86QPBwqc2FCsQzijPkXD"
	BITFINEX_BCH_IAN = "1NhWLT9H4RkayXxBjLuGKYpWLn3NgDvWNb"
	BITFINEX_ETH_IAN = "0x248c8cf7b7f93a8806efaf69eb0004ed53e80061"
	BITFINEX_ETC_IAN = "0x248c8cf7b7f93a8806efaf69eb0004ed53e80061"
	BITFINEX_LTC_IAN = "LKbmZK4uW1GLfLTwUWujYjUuxEm8ZJteTe"

	KRAKEN_BTC_IAN = "3AsFf7NfnSiqVeQBPfkfqbqP8jeR5CwMv4"
	KRAKEN_BCH_IAN = "1F3U4DHPCGYyVd4RxFaPj4K2DZNNdrPz8h"
	KRAKEN_ETH_IAN = "0x03F86Bc0329653A982c156e67a070328CB426aFC"
	KRAKEN_ETC_IAN = "0x03f86bc0329653a982c156e67a070328cb426afc"
	KRAKEN_LTC_IAN = "LUtZwuCKWpLFxjKDcnEW9P8tHj6UFXPxSV"

	OKCOIN_BTC_IAN = "36svBm2JSsvSNdSEfoaRWD1Vj5SN2bh2wW"
	OKCOIN_BCH_IAN = "13VudiWsAz7rBefKUeUD3V8QWKuba3FHjX"
	OKCOIN_ETH_IAN = "0xbd9fb3d2d40274a0e3e16a62cbcd7cf139621b13"
	OKCOIN_ETC_IAN = "0x882b717b2b31f8747c0af462e27cacdbf5b307b0"
	OKCOIN_LTC_IAN = "36D7EGf7XEJLFBpDKKK2QJxf87azMs84Ts"
)

// CryptoAddress stores the address hash and related information.
type CryptoAddress struct {
	currency     goex.Currency
	address      string
	tag          string
	exchangeName string
}

// Currency returns the goex currency
func (c *CryptoAddress) Currency() goex.Currency {
	return c.currency
}

// Address returns the address hash as a string
func (c *CryptoAddress) Address() string {
	return c.address
}

// Tag returns the tag as a string
func (c *CryptoAddress) Tag() string {
	return c.tag
}

// ExchangeName returns the exchange name as a string
func (c *CryptoAddress) ExchangeName() string {
	return c.exchangeName
}

var IAN_BITFINEX_EXCHANGE_BTC = CryptoAddress{
	currency:     goex.BTC,
	address:      BITFINEX_BTC_IAN,
	tag:          "I_B_BTC",
	exchangeName: "bitfinex.com",
}

var IAN_BITFINEX_EXCHANGE_BCH = CryptoAddress{
	currency:     goex.BCH,
	address:      BITFINEX_BCH_IAN,
	tag:          "I_B_BCH",
	exchangeName: "bitfinex.com",
}

var IAN_BITFINEX_EXCHANGE_ETC = CryptoAddress{
	currency:     goex.ETC,
	address:      BITFINEX_ETC_IAN,
	tag:          "I_B_ETC",
	exchangeName: "bitfinex.com",
}

var IAN_BITFINEX_EXCHANGE_ETH = CryptoAddress{
	currency:     goex.ETH,
	address:      BITFINEX_ETH_IAN,
	tag:          "I_B_ETH",
	exchangeName: "bitfinex.com",
}

var IAN_BITFINEX_EXCHANGE_LTC = CryptoAddress{
	currency:     goex.LTC,
	address:      BITFINEX_LTC_IAN,
	tag:          "I_B_ETH",
	exchangeName: "bitfinex.com",
}

var IAN_KRAKEN_BTC = CryptoAddress{
	currency:     goex.BTC,
	address:      KRAKEN_BTC_IAN,
	tag:          "I_K_BTC",
	exchangeName: "kraken.com",
}

var IAN_KRAKEN_BCH = CryptoAddress{
	currency:     goex.BCH,
	address:      KRAKEN_BCH_IAN,
	tag:          "I_K_BCH",
	exchangeName: "kraken.com",
}

var IAN_KRAKEN_ETC = CryptoAddress{
	currency:     goex.ETC,
	address:      KRAKEN_ETC_IAN,
	tag:          "I_K_ETC",
	exchangeName: "kraken.com",
}

var IAN_KRAKEN_ETH = CryptoAddress{
	currency:     goex.ETH,
	address:      KRAKEN_ETH_IAN,
	tag:          "I_K_ETH",
	exchangeName: "kraken.com",
}

var IAN_KRAKEN_LTC = CryptoAddress{
	currency:     goex.LTC,
	address:      KRAKEN_LTC_IAN,
	tag:          "I_K_ETH",
	exchangeName: "kraken.com",
}

var IAN_OKCOIN_BTC = CryptoAddress{
	currency:     goex.BTC,
	address:      OKCOIN_BTC_IAN,
	tag:          "I_O_BTC",
	exchangeName: "okcoin.com",
}

var IAN_OKCOIN_BCH = CryptoAddress{
	currency:     goex.BCH,
	address:      OKCOIN_BCH_IAN,
	tag:          "I_O_BCH",
	exchangeName: "okcoin.com",
}

var IAN_OKCOIN_ETC = CryptoAddress{
	currency:     goex.ETC,
	address:      OKCOIN_ETC_IAN,
	tag:          "I_O_ETC",
	exchangeName: "okcoin.com",
}

var IAN_OKCOIN_ETH = CryptoAddress{
	currency:     goex.ETH,
	address:      OKCOIN_ETH_IAN,
	tag:          "I_O_ETH",
	exchangeName: "okcoin.com",
}

var IAN_OKCOIN_LTC = CryptoAddress{
	currency:     goex.LTC,
	address:      OKCOIN_LTC_IAN,
	tag:          "I_O_ETH",
	exchangeName: "okcoin.com",
}

//var Addresses map[string]map[goex.Currency]CryptoAddress
//var BitfinexAddresses map[goex.Currency]CryptoAddress
//Addresses = make(map[string]map[goex.Currency]CryptoAddress)

var BitfinexAddresses = map[goex.Currency]CryptoAddress{
	goex.BTC: IAN_BITFINEX_EXCHANGE_BTC,
	goex.BCH: IAN_BITFINEX_EXCHANGE_BCH,
	goex.ETC: IAN_BITFINEX_EXCHANGE_ETC,
	goex.ETH: IAN_BITFINEX_EXCHANGE_ETH,
	goex.LTC: IAN_BITFINEX_EXCHANGE_LTC,
}

var KrakenAddresses = map[goex.Currency]CryptoAddress{
	goex.BTC: IAN_KRAKEN_BTC,
	goex.BCH: IAN_KRAKEN_BCH,
	goex.ETC: IAN_KRAKEN_ETC,
	goex.ETH: IAN_KRAKEN_ETH,
	goex.LTC: IAN_KRAKEN_LTC,
}

var OkcoinAddresses = map[goex.Currency]CryptoAddress{
	goex.BTC: IAN_OKCOIN_BTC,
	goex.BCH: IAN_OKCOIN_BCH,
	goex.ETC: IAN_OKCOIN_ETC,
	goex.ETH: IAN_OKCOIN_ETH,
	goex.LTC: IAN_OKCOIN_LTC,
}

var All = map[string]map[goex.Currency]CryptoAddress{
	"bitfinex.com": BitfinexAddresses,
	"kraken.com":   KrakenAddresses,
	"okcoin.com":   OkcoinAddresses,
}
