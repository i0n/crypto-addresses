package addresses_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	addresses "github.com/i0n/crypto-addresses"
	goex "github.com/nntaoli-project/GoEx"
)

func TestAddressesAll(t *testing.T) {
	destinationAPIName := "okcoin.com"
	currencyPair := goex.ETH_USD
	a := addresses.All[destinationAPIName][currencyPair.CurrencyA].Tag()
	assert.Equal(t, "I_O_ETH", a)
}
