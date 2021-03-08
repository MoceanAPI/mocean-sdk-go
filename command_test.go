package moceansdk

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"net/url"
	"testing"
)

func TestTgSendMessage(t *testing.T) {
	singleTgRes := ReadResourceFile("send_single_tg.json")
	httpmock.RegisterResponder("POST", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.Command().commandURL,
		httpmock.NewStringResponder(http.StatusAccepted, singleTgRes))

	res, err := _mocean.Command().Execute(url.Values{})
	AssertNoError(t, err)
	AssertEqual(t, res.String(), singleTgRes)
}

func TestTgSendMessageError(t *testing.T) {
	tgErrorRes := ReadResourceFile("error_response.json")
	httpmock.RegisterResponder("POST", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.Command().commandURL,
		httpmock.NewStringResponder(http.StatusBadRequest, tgErrorRes))

	_, err := _mocean.Command().Execute(url.Values{})
	AssertError(t, err)
	AssertEqual(t, err.Error(), "Authorization failed")
}
