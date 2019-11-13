package moceansdk

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"net/url"
	"testing"
)

func TestNumberLookupService_Inquiry(t *testing.T) {
	numberLookupRes := ReadResourceFile("number_lookup.json")
	httpmock.RegisterResponder("POST", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.NumberLookup().numberLookupURL,
		httpmock.NewStringResponder(http.StatusAccepted, numberLookupRes))

	res, err := _mocean.NumberLookup().Inquiry(url.Values{})
	AssertNoError(t, err)
	AssertEqual(t, res.String(), numberLookupRes)
}

func TestNumberLookupService_InquiryError(t *testing.T) {
	errorRes := ReadResourceFile("error_response.json")
	httpmock.RegisterResponder("POST", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.NumberLookup().numberLookupURL,
		httpmock.NewStringResponder(http.StatusBadRequest, errorRes))

	_, err := _mocean.NumberLookup().Inquiry(url.Values{})
	AssertError(t, err)
	AssertEqual(t, err.Error(), "Authorization failed")
}
