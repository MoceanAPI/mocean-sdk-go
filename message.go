package moceansdk

import (
	"encoding/json"
	"net/url"
)

type MessageService struct {
	client           *Mocean
	smsUrl           string
	messageStatusUrl string
}

//Message constructor
func (mocean *Mocean) Message() *MessageService {
	return &MessageService{
		mocean,
		"/sms",
		"/report/message",
	}
}

type SmsResponse struct {
	abstractResponse
	Messages []struct {
		Status   interface{} `json:"status"`
		Receiver interface{} `json:"receiver"`
		Msgid    interface{} `json:"msgid"`
	} `json:"messages"`
}

//Send SMS
//For more info, see docs: https://moceanapi.com/docs/#send-sms
func (s *MessageService) Send(params url.Values) (smsResponse *SmsResponse, err error) {
	res, err := s.client.post(s.smsUrl, params)
	if err != nil {
		return smsResponse, err
	}

	smsResponse = new(SmsResponse)
	err = json.Unmarshal(res, smsResponse)

	smsResponse.rawResponse = string(res)
	return smsResponse, err
}

type MsgStatusResponse struct {
	abstractResponse
	MessageStatus  interface{} `json:"message_status"`
	Msgid          interface{} `json:"msgid"`
	CreditDeducted interface{} `json:"credit_deducted"`
}

//Get Message Status
//For more info, see docs: https://moceanapi.com/docs/#message-status
func (s *MessageService) GetMessageStatus(params url.Values) (msgStatusResponse *MsgStatusResponse, err error) {
	res, err := s.client.get(s.messageStatusUrl, params)
	if err != nil {
		return msgStatusResponse, err
	}

	msgStatusResponse = new(MsgStatusResponse)
	err = json.Unmarshal(res, msgStatusResponse)

	msgStatusResponse.rawResponse = string(res)
	return msgStatusResponse, err
}
