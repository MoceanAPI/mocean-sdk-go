package moceansdk

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"net/url"
	"testing"
)

func TestVoiceService_Call(t *testing.T) {
	voiceRes := ReadResourceFile("voice.json")
	httpmock.RegisterResponder("GET", _mocean.Options.BaseUrl+"/rest/"+_mocean.Options.Version+_mocean.Voice().voiceUrl,
		httpmock.NewStringResponder(http.StatusAccepted, voiceRes))

	res, err := _mocean.Voice().Call(url.Values{})
	AssertNoError(t, err)
	AssertEqual(t, res.String(), voiceRes)
}

func TestVoiceService_CallError(t *testing.T) {
	voiceRes := ReadResourceFile("error_response.json")
	httpmock.RegisterResponder("GET", _mocean.Options.BaseUrl+"/rest/"+_mocean.Options.Version+_mocean.Voice().voiceUrl,
		httpmock.NewStringResponder(http.StatusBadRequest, voiceRes))

	_, err := _mocean.Voice().Call(url.Values{})
	AssertError(t, err)
	AssertEqual(t, err.Error(), "Authorization failed")
}
