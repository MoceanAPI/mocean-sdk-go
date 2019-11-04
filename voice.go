package moceansdk

import (
	"encoding/json"
	"net/url"
)

type voiceService struct {
	client   *mocean
	voiceURL string
}

//Voice Constructor
func (m *mocean) Voice() *voiceService {
	return &voiceService{
		m,
		"/voice/dial",
	}
}

type voiceResponse struct {
	abstractResponse
	Calls []struct {
		Status      interface{} `json:"status"`
		Receiver    interface{} `json:"receiver"`
		SessionUUID interface{} `json:"session-uuid"`
		CallUUID    interface{} `json:"call-uuid"`
	} `json:"calls"`
}

//Voice
//For more info, see docs: https://moceanapi.com/docs/#voice
func (s *voiceService) Call(params url.Values) (response *voiceResponse, err error) {
	res, err := s.client.post(s.voiceURL, params)
	if err != nil {
		return response, err
	}

	response = new(voiceResponse)
	err = json.Unmarshal(res, response)

	response.rawResponse = string(res)
	return response, err
}
