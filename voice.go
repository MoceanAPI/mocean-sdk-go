package moceansdk

import (
	"encoding/json"
	"net/url"
)

type VoiceService struct {
	client   *Mocean
	voiceURL string
}

//Voice Constructor
func (m *Mocean) Voice() *VoiceService {
	return &VoiceService{
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
func (s *VoiceService) Call(params url.Values) (response *voiceResponse, err error) {
	res, err := s.client.post(s.voiceURL, params)
	if err != nil {
		return response, err
	}

	response = new(voiceResponse)
	err = json.Unmarshal(res, response)

	response.rawResponse = string(res)
	return response, err
}
