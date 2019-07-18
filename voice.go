package moceansdk

import (
	"encoding/json"
	"net/url"
)

type voiceService struct {
	client   *mocean
	voiceUrl string
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
	SessionUuid interface{} `json:"session-uuid"`
	CallUuid    interface{} `json:"call-uuid"`
}

//Voice
//For more info, see docs: https://moceanapi.com/docs/#voice
func (s *voiceService) Call(params url.Values) (response *voiceResponse, err error) {
	res, err := s.client.get(s.voiceUrl, params)
	if err != nil {
		return response, err
	}

	response = new(voiceResponse)
	err = json.Unmarshal(res, response)

	response.rawResponse = string(res)
	return response, err
}
