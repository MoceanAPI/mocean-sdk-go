package moceansdk

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"net/url"
	"testing"
)

func TestVoiceService_Call(t *testing.T) {
	voiceRes := ReadResourceFile("voice.json")
	httpmock.RegisterResponder("POST", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.Voice().voiceURL,
		httpmock.NewStringResponder(http.StatusAccepted, voiceRes))

	res, err := _mocean.Voice().Call(url.Values{})
	AssertNoError(t, err)
	AssertEqual(t, res.String(), voiceRes)
}

func TestVoiceService_CallError(t *testing.T) {
	voiceRes := ReadResourceFile("error_response.json")
	httpmock.RegisterResponder("POST", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.Voice().voiceURL,
		httpmock.NewStringResponder(http.StatusBadRequest, voiceRes))

	_, err := _mocean.Voice().Call(url.Values{})
	AssertError(t, err)
	AssertEqual(t, err.Error(), "Authorization failed")
}

func TestVoiceService_Hangup(t *testing.T) {
	hangupRes := ReadResourceFile("hangup.json")
	httpmock.RegisterResponder("POST", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.Voice().hangupURL+"/xxx-xxx-xxx-xxx",
		httpmock.NewStringResponder(http.StatusAccepted, hangupRes))

	res, err := _mocean.Voice().Hangup("xxx-xxx-xxx-xxx")
	AssertNoError(t, err)
	AssertEqual(t, res.String(), hangupRes)
}

func TestVoiceService_HangupError(t *testing.T) {
	voiceRes := ReadResourceFile("error_response.json")
	httpmock.RegisterResponder("POST", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.Voice().hangupURL+"/xxx-xxx-xxx-xxx",
		httpmock.NewStringResponder(http.StatusBadRequest, voiceRes))

	_, err := _mocean.Voice().Hangup("xxx-xxx-xxx-xxx")
	AssertError(t, err)
	AssertEqual(t, err.Error(), "Authorization failed")
}

func TestVoiceService_Recording(t *testing.T) {
	recordingRes := ReadResourceFile("recording.json")
	httpmock.RegisterResponder("GET", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.Voice().recordingURL,
		httpmock.NewStringResponder(http.StatusAccepted, recordingRes))

	res, err := _mocean.Voice().Recording("xxx-xxx-xxx-xxx")
	AssertNoError(t, err)
	AssertEqual(t, res.Filename, "xxx-xxx-xxx-xxx.mp3")
	AssertNotNil(t, res.RecordingBuffer)
}

func TestVoiceService_RecordingError(t *testing.T) {
	recordingRes := ReadResourceFile("error_response.json")
	httpmock.RegisterResponder("GET", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.Voice().recordingURL,
		httpmock.NewStringResponder(http.StatusBadRequest, recordingRes))

	_, err := _mocean.Voice().Recording("xxx-xxx-xxx-xxx")
	AssertError(t, err)
	AssertEqual(t, err.Error(), "Authorization failed")
}
