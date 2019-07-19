package moceansdk

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"net/url"
	"testing"
)

func TestMessageService_GetMessageStatus(t *testing.T) {
	msgStatusRes := ReadResourceFile("message_status.json")
	httpmock.RegisterResponder("GET", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.Message().messageStatusURL,
		httpmock.NewStringResponder(http.StatusAccepted, msgStatusRes))

	res, err := _mocean.Message().GetMessageStatus(url.Values{})
	AssertNoError(t, err)
	AssertEqual(t, res.String(), msgStatusRes)
}

func TestMessageService_GetMessageStatusError(t *testing.T) {
	errorRes := ReadResourceFile("error_response.json")
	httpmock.RegisterResponder("GET", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.Message().messageStatusURL,
		httpmock.NewStringResponder(http.StatusBadRequest, errorRes))

	_, err := _mocean.Message().GetMessageStatus(url.Values{})
	AssertError(t, err)
	AssertEqual(t, err.Error(), "Authorization failed")
}

func TestMessageService_Send(t *testing.T) {
	msgRes := ReadResourceFile("message.json")
	httpmock.RegisterResponder("POST", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.Message().smsURL,
		httpmock.NewStringResponder(http.StatusAccepted, msgRes))

	res, err := _mocean.Message().Send(url.Values{})
	AssertNoError(t, err)
	AssertEqual(t, res.String(), msgRes)
}

func TestMessageService_SendError(t *testing.T) {
	errorRes := ReadResourceFile("error_response.json")
	httpmock.RegisterResponder("POST", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.Message().smsURL,
		httpmock.NewStringResponder(http.StatusBadRequest, errorRes))

	_, err := _mocean.Message().Send(url.Values{})
	AssertError(t, err)
	AssertEqual(t, err.Error(), "Authorization failed")
}
