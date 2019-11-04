package moceansdk

import (
	"encoding/json"
	"net/url"
	"strings"
)

type VerifyService struct {
	client        *Mocean
	sendCodeURL   string
	verifyCodeURL string
	channel       string
	isResend      bool
}

func (s *VerifyService) SendAs(channel string) *VerifyService {
	s.channel = channel
	return s
}

//Verify Constructor
func (m *Mocean) Verify() *VerifyService {
	return &VerifyService{
		m,
		"/verify",
		"/verify/check",
		"AUTO",
		false,
	}
}

type sendCodeResponse struct {
	abstractResponse
	Reqid        interface{} `json:"reqid"`
	To           interface{} `json:"to"`
	ResendNumber interface{} `json:"resend_number"`
}

//Send verify code
//For more info, see docs: https://moceanapi.com/docs/#send-code
func (s *VerifyService) SendCode(params url.Values) (response *sendCodeResponse, err error) {
	sendCodeURL := s.sendCodeURL

	if s.isResend == true {
		sendCodeURL += "/resend"
	} else {
		sendCodeURL += "/req"
	}

	if strings.EqualFold(s.channel, "sms") {
		sendCodeURL += "/sms"
	}

	res, err := s.client.post(sendCodeURL, params)
	if err != nil {
		return response, err
	}

	response = new(sendCodeResponse)
	err = json.Unmarshal(res, response)

	response.rawResponse = string(res)
	return response, err
}

type verifyCodeResponse struct {
	abstractResponse
	Reqid interface{} `json:"reqid"`
}

//Verify code
//For more info, see docs: https://moceanapi.com/docs/#verify-code
func (s *VerifyService) VerifyCode(params url.Values) (response *verifyCodeResponse, err error) {
	res, err := s.client.post(s.verifyCodeURL, params)
	if err != nil {
		return response, err
	}

	response = new(verifyCodeResponse)
	err = json.Unmarshal(res, response)

	response.rawResponse = string(res)
	return response, err
}

func (s *VerifyService) Resend(params url.Values) (response *sendCodeResponse, err error) {
	s.SendAs("SMS")
	s.isResend = true
	return s.SendCode(params)
}
