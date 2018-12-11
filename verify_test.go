package moceango

import (
	"testing"
)

func TestVerifySendCode(t *testing.T) {
	mocean := NewMoceanClient(testParams["apiKey"], testParams["apiSecret"])

	verifyParams := &VerifyReqParams{
		Brand: "Mocean",
		To:    "60123456789",
	}

	res, err := mocean.Verify().sendCode(verifyParams)

	if err != nil {
		t.Fatal(err)
	}

	t.Log("Send Code Req Id: " + res.Reqid)
}

func TestVerifyCode(t *testing.T) {
	mocean := NewMoceanClient(testParams["apiKey"], testParams["apiSecret"])

	verifyParams := &VerifyCodeParams{
		Reqid: "CPASS_restapi_C1211111252001230.0002",
		Code:  "192165",
	}

	res, err := mocean.Verify().verifyCode(verifyParams)

	if err != nil {
		t.Fatal(err)
	}

	t.Log("Verify Code Status: " + res.Status)
}
