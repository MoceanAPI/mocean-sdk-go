package moceansdk

import (
	"encoding/json"
	"net/url"
	"strings"
)

type verifyService struct {
	client        *mocean
	sendCodeUrl   string
	verifyCodeUrl string
	channel       string
	isResend      bool
}

func (s *verifyService) SendAs(channel string) *verifyService {
	s.channel = channel
	return s
}

//Verify Constructor
func (m *mocean) Verify() *verifyService {
	return &verifyService{
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
func (s *verifyService) SendCode(params url.Values) (response *sendCodeResponse, err error) {
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
func (s *verifyService) VerifyCode(params url.Values) (response *verifyCodeResponse, err error) {
	res, err := s.client.post(s.verifyCodeUrl, params)
	if err != nil {
		return response, err
	}

	response = new(verifyCodeResponse)
	err = json.Unmarshal(res, response)

	response.rawResponse = string(res)
	return response, err
}

func (s *verifyService) Resend(params url.Values) (response *sendCodeResponse, err error) {
	s.SendAs("SMS")
	s.isResend = true
	return s.SendCode(params)
}
