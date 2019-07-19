package moceansdk

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"net/url"
	"testing"
)

func TestAccountService_GetBalance(t *testing.T) {
	balRes := ReadResourceFile("price.json")
	httpmock.RegisterResponder("GET", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.Account().balanceURL,
		httpmock.NewStringResponder(http.StatusAccepted, balRes))

	res, err := _mocean.Account().GetBalance(url.Values{})
	AssertNoError(t, err)
	AssertEqual(t, res.String(), balRes)
}

func TestAccountService_GetBalanceError(t *testing.T) {
	errorRes := ReadResourceFile("error_response.json")
	httpmock.RegisterResponder("GET", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.Account().balanceURL,
		httpmock.NewStringResponder(http.StatusBadRequest, errorRes))

	_, err := _mocean.Account().GetBalance(url.Values{})
	AssertError(t, err)
	AssertEqual(t, err.Error(), "Authorization failed")
}

func TestAccountService_GetPricing(t *testing.T) {
	priceRes := ReadResourceFile("price.json")
	httpmock.RegisterResponder("GET", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.Account().pricingURL,
		httpmock.NewStringResponder(http.StatusAccepted, priceRes))

	res, err := _mocean.Account().GetPricing(url.Values{})
	AssertNoError(t, err)
	AssertEqual(t, res.String(), priceRes)
}

func TestAccountService_GetPricingError(t *testing.T) {
	errorRes := ReadResourceFile("error_response.json")
	httpmock.RegisterResponder("GET", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.Account().pricingURL,
		httpmock.NewStringResponder(http.StatusBadRequest, errorRes))

	_, err := _mocean.Account().GetPricing(url.Values{})
	AssertError(t, err)
	AssertEqual(t, err.Error(), "Authorization failed")
}
