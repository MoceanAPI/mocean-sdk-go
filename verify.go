package moceango

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type verify struct {
	*Mocean
	SendCodeUrl   string
	VerifyCodeUrl string
}

type VerifyReqParams struct {
	To            string
	Brand         string
	From          string
	Codelength    string
	PinValidity   string
	NextEventWait string
}

type VerifyReqResponse struct {
	Status string
	Reqid  string
}

//Verify Constructor
func (mocean *Mocean) Verify() *verify {
	return &verify{
		mocean,
		mocean.BaseUrl + "/verify/req",
		mocean.BaseUrl + "/verify/check",
	}
}

//Send verify code
//For more info, see docs: https://moceanapi.com/docs/#send-code
func (verify *verify) sendCode(params *VerifyReqParams) (verifyReqResponse *VerifyReqResponse, err error) {
	formData := verify.makeFormData(verify.ApiKey, verify.ApiSecret);
	formData = verify.setVerifyReqParams(params, formData)
	res, err := verify.post(verify.SendCodeUrl, formData)
	if err != nil {
		return verifyReqResponse, err
	}

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return verifyReqResponse, err
	}

	if res.StatusCode != http.StatusAccepted {
		errRes := new(ErrorResponse)
		err = json.Unmarshal(responseBody, errRes)

		return verifyReqResponse, errors.New(errRes.ErrorMsg)
	}

	verifyReqResponse = new(VerifyReqResponse)
	err = json.Unmarshal(responseBody, verifyReqResponse)

	return verifyReqResponse, err
}

type VerifyCodeParams struct {
	Reqid string
	Code  string
}

type VerifyCodeResponse struct {
	Reqid    string
	Status   string
	Price    string
	Currency string
}

//Verify code
//For more info, see docs: https://moceanapi.com/docs/#verify-code
func (verify *verify) verifyCode(params *VerifyCodeParams) (verifyCodeResponse *VerifyCodeResponse, err error) {
	formData := verify.makeFormData(verify.ApiKey, verify.ApiSecret);
	formData.Set("mocean-reqid", params.Reqid)
	formData.Set("mocean-code", params.Code)

	res, err := verify.post(verify.VerifyCodeUrl, formData)
	if err != nil {
		return verifyCodeResponse, err
	}

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return verifyCodeResponse, err
	}

	if res.StatusCode != http.StatusAccepted {
		errRes := new(ErrorResponse)
		err = json.Unmarshal(responseBody, errRes)

		return verifyCodeResponse, errors.New(errRes.ErrorMsg)
	}

	verifyCodeResponse = new(VerifyCodeResponse)
	err = json.Unmarshal(responseBody, verifyCodeResponse)

	return verifyCodeResponse, err
}

func (verify *verify) setVerifyReqParams(params *VerifyReqParams, inputParams url.Values) url.Values {
	if params.From != "" {
		inputParams.Set("mocean-from", params.From)
	}
	if params.To != "" {
		inputParams.Set("mocean-to", params.To)
	}
	if params.Brand != "" {
		inputParams.Set("mocean-brand", params.Brand)
	}
	if params.Codelength != "" {
		inputParams.Set("mocean-code-length", params.Codelength)
	}
	if params.PinValidity != "" {
		inputParams.Set("mocean-pin-validity", params.PinValidity)
	}
	if params.NextEventWait != "" {
		inputParams.Set("mocean-next-event-wait", params.NextEventWait)
	}
	return inputParams
}
