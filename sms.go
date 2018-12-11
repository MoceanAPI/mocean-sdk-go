package moceango

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type sms struct {
	*Mocean
	SendSMSUrl          string
	GetMessageStatusUrl string
}

type Message struct {
	From     string
	To       string
	Text     string
	Udh      string
	Coding   string
	Dlrmask  string
	Dlrurl   string
	Schedule string
	Mclass   string
	Altdcs   string
	Charset  string
	Validity string
}

type SendSMSResponse struct {
	Messages []struct {
		Status   int    `json:"status"`
		Receiver string `json:"receiver"`
		Msgid    string `json:"msgid"`
	} `json:"messages"`
}

//SMS constructor
func (mocean *Mocean) Sms() *sms {
	return &sms{
		mocean,
		mocean.BaseUrl + "/sms",
		mocean.BaseUrl + "/report/message",
	}
}

//Send SMS
//For more info, see docs: https://moceanapi.com/docs/#send-sms
func (sms *sms) Send(message *Message) (sendSMSResponse *SendSMSResponse, err error) {
	formData := sms.makeFormData(sms.ApiKey, sms.ApiSecret);
	formData = sms.setParams(message, formData)

	res, err := sms.post(sms.SendSMSUrl, formData)
	if err != nil {
		return sendSMSResponse, err
	}

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return sendSMSResponse, err
	}

	if res.StatusCode != http.StatusAccepted {
		errRes := new(ErrorResponse)
		err = json.Unmarshal(responseBody, errRes)

		return sendSMSResponse, errors.New(errRes.ErrorMsg)
	}

	sendSMSResponse = new(SendSMSResponse)
	err = json.Unmarshal(responseBody, &sendSMSResponse)

	return sendSMSResponse, err
}

type MsgStatusResponse struct {
	Status         int     `json:"status"`
	MessageStatus  int     `json:"message_status"`
	Msgid          string  `json:"msgid"`
	CreditDeducted float64 `json:"credit_deducted"`
}

//Get Message Status
//For more info, see docs: https://moceanapi.com/docs/#message-status
func (sms *sms) getMsgStatus(msgid string) (msgStatusResponse *MsgStatusResponse, err error) {
	formData := sms.makeFormData(sms.ApiKey, sms.ApiSecret);
	formData.Set("mocean-msgid", msgid)
	res, err := sms.get(sms.GetMessageStatusUrl + "?" + formData.Encode())
	if err != nil {
		return msgStatusResponse, err
	}

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return msgStatusResponse, err
	}

	if res.StatusCode != http.StatusAccepted {
		errRes := new(ErrorResponse)
		err = json.Unmarshal(responseBody, errRes)

		return msgStatusResponse, errors.New(errRes.ErrorMsg)
	}

	msgStatusResponse = new(MsgStatusResponse)
	err = json.Unmarshal(responseBody, msgStatusResponse)

	return msgStatusResponse, err
}

func (sms *sms) setParams(message *Message, inputParams url.Values) url.Values {
	if message.From != "" {
		inputParams.Set("mocean-from", message.From)
	}
	if message.To != "" {
		inputParams.Set("mocean-to", message.To)
	}
	if message.Text != "" {
		inputParams.Set("mocean-text", message.Text)
	}
	if message.Udh != "" {
		inputParams.Set("mocean-udh", message.Udh)
	}
	if message.Coding != "" {
		inputParams.Set("mocean-coding", message.Coding)
	}
	if message.Dlrmask != "" {
		inputParams.Set("mocean-dlr-mask", message.Dlrmask)
	}
	if message.Dlrurl != "" {
		inputParams.Set("mocean-dlr-url", message.Dlrurl)
	}
	if message.Schedule != "" {
		inputParams.Set("mocean-schedule", message.Schedule)
	}
	if message.Mclass != "" {
		inputParams.Set("mocean-mclass", message.Mclass)
	}
	if message.Altdcs != "" {
		inputParams.Set("mocean-alt-dcs", message.Altdcs)
	}
	if message.Charset != "" {
		inputParams.Set("mocean-charset", message.Charset)
	}
	if message.Validity != "" {
		inputParams.Set("mocean-validity", message.Validity)
	}
	return inputParams
}
