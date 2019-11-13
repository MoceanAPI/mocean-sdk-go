package moceansdk

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"net/url"
	"testing"
)

func TestVerifyService_SendCode(t *testing.T) {
	sendCodeRes := ReadResourceFile("send_code.json")
	httpmock.RegisterResponder("POST", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.Verify().sendCodeURL+"/req",
		httpmock.NewStringResponder(http.StatusAccepted, sendCodeRes))

	res, err := _mocean.Verify().SendCode(url.Values{})
	AssertNoError(t, err)
	AssertEqual(t, res.String(), sendCodeRes)
}

func TestVerifyService_SendAsSmsChannel(t *testing.T) {
	verifySv := _mocean.Verify().SendAs("sms")
	AssertEqual(t, verifySv.channel, "sms")

	sendCodeRes := ReadResourceFile("send_code.json")
	httpmock.RegisterResponder("POST", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.Verify().sendCodeURL+"/req/sms",
		httpmock.NewStringResponder(http.StatusAccepted, sendCodeRes))

	res, err := verifySv.SendCode(url.Values{})
	AssertNoError(t, err)
	AssertEqual(t, res.String(), sendCodeRes)
}

func TestVerifyService_Resend(t *testing.T) {
	resendCodeRes := ReadResourceFile("resend_code.json")
	httpmock.RegisterResponder("POST", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.Verify().sendCodeURL+"/resend/sms",
		httpmock.NewStringResponder(http.StatusAccepted, resendCodeRes))

	res, err := _mocean.Verify().Resend(url.Values{})
	AssertNoError(t, err)
	AssertEqual(t, res.String(), resendCodeRes)
}

func TestVerifyService_SendCodeError(t *testing.T) {
	errorRes := ReadResourceFile("error_response.json")
	httpmock.RegisterResponder("POST", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.Verify().sendCodeURL+"/req",
		httpmock.NewStringResponder(http.StatusBadRequest, errorRes))

	_, err := _mocean.Verify().SendCode(url.Values{})
	AssertError(t, err)
	AssertEqual(t, err.Error(), "Authorization failed")
}

func TestVerifyService_VerifyCode(t *testing.T) {
	verifyCodeRes := ReadResourceFile("verify_code.json")
	httpmock.RegisterResponder("POST", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.Verify().verifyCodeURL,
		httpmock.NewStringResponder(http.StatusAccepted, verifyCodeRes))

	res, err := _mocean.Verify().VerifyCode(url.Values{})
	AssertNoError(t, err)
	AssertEqual(t, res.String(), verifyCodeRes)
}

func TestVerifyService_VerifyCodeError(t *testing.T) {
	errorRes := ReadResourceFile("error_response.json")
	httpmock.RegisterResponder("POST", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+_mocean.Verify().verifyCodeURL,
		httpmock.NewStringResponder(http.StatusBadRequest, errorRes))

	_, err := _mocean.Verify().VerifyCode(url.Values{})
	AssertError(t, err)
	AssertEqual(t, err.Error(), "Authorization failed")
}
