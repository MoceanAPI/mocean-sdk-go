package moceango

import (
	"testing"
)

func TestBalance(t *testing.T) {
	mocean := NewMoceanClient(testParams["apiKey"], testParams["apiSecret"])

	res, err := mocean.Account().getBalance()

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Balance: %f", res.Balance)
}

func TestPricing(t *testing.T) {
	mocean := NewMoceanClient(testParams["apiKey"], testParams["apiSecret"])

	res, err := mocean.Account().getPricing()

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Pricing: %f", res.Destinations[0].Price)
}
