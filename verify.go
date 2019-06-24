package moceansdk

import (
	"encoding/json"
	"net/url"
	"strings"
)

type VerifyService struct {
	client        *Mocean
	sendCodeUrl   string
	verifyCodeUrl string
	channel       string
	isResend      bool
}

func (s *VerifyService) SendAs(channel string) *VerifyService {
	s.channel = channel
	return s
}

//Verify Constructor
func (mocean *Mocean) Verify() *VerifyService {
	return &VerifyService{
		mocean,
		"/verify",
		"/verify/check",
		"AUTO",
		false,
	}
}

type SendCodeResponse struct {
	abstractResponse
	Reqid        interface{} `json:"reqid"`
	To           interface{} `json:"to"`
	ResendNumber interface{} `json:"resend_number"`
}

//Send verify code
//For more info, see docs: https://moceanapi.com/docs/#send-code
func (s *VerifyService) SendCode(params url.Values) (sendCodeResponse *SendCodeResponse, err error) {
	sendCodeUrl := s.sendCodeUrl

	if s.isResend == true {
		sendCodeUrl += "/resend"
	} else {
		sendCodeUrl += "/req"
	}

	if strings.EqualFold(s.channel, "sms") {
		sendCodeUrl += "/sms"
	}

	res, err := s.client.post(sendCodeUrl, params)
	if err != nil {
		return sendCodeResponse, err
	}

	sendCodeResponse = new(SendCodeResponse)
	err = json.Unmarshal(res, sendCodeResponse)

	sendCodeResponse.rawResponse = string(res)
	return sendCodeResponse, err
}

type VerifyCodeResponse struct {
	abstractResponse
	Reqid interface{} `json:"reqid"`
}

//Verify code
//For more info, see docs: https://moceanapi.com/docs/#verify-code
func (s *VerifyService) VerifyCode(params url.Values) (verifyCodeResponse *VerifyCodeResponse, err error) {
	res, err := s.client.post(s.verifyCodeUrl, params)
	if err != nil {
		return verifyCodeResponse, err
	}

	verifyCodeResponse = new(VerifyCodeResponse)
	err = json.Unmarshal(res, verifyCodeResponse)

	verifyCodeResponse.rawResponse = string(res)
	return verifyCodeResponse, err
}

func (s *VerifyService) Resend(params url.Values) (sendCodeResponse *SendCodeResponse, err error) {
	s.SendAs("SMS")
	s.isResend = true
	return s.SendCode(params)
}
