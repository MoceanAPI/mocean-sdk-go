package moceansdk

import (
	"encoding/json"
	"net/url"
)

type messageService struct {
	client           *Mocean
	smsURL           string
	messageStatusURL string
}

//Message constructor
func (m *Mocean) Message() *messageService {
	return &messageService{
		m,
		"/sms",
		"/report/message",
	}
}

type smsResponse struct {
	abstractResponse
	Messages []struct {
		Status   interface{} `json:"status"`
		Receiver interface{} `json:"receiver"`
		Msgid    interface{} `json:"msgid"`
	} `json:"messages"`
}

//Send SMS
//For more info, see docs: https://moceanapi.com/docs/#send-sms
func (s *messageService) Send(params url.Values) (response *smsResponse, err error) {
	res, err := s.client.post(s.smsURL, params)
	if err != nil {
		return response, err
	}

	response = new(smsResponse)
	err = json.Unmarshal(res, response)

	response.rawResponse = string(res)
	return response, err
}

type msgStatusResponse struct {
	abstractResponse
	MessageStatus  interface{} `json:"message_status"`
	Msgid          interface{} `json:"msgid"`
	CreditDeducted interface{} `json:"credit_deducted"`
}

//Get Message Status
//For more info, see docs: https://moceanapi.com/docs/#message-status
func (s *messageService) GetMessageStatus(params url.Values) (response *msgStatusResponse, err error) {
	res, err := s.client.get(s.messageStatusURL, params)
	if err != nil {
		return response, err
	}

	response = new(msgStatusResponse)
	err = json.Unmarshal(res, response)

	response.rawResponse = string(res)
	return response, err
}
