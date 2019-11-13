package moceansdk

import (
	"bytes"
	"encoding/json"
	"net/url"
)

type voiceService struct {
	client       *mocean
	voiceURL     string
	hangupURL    string
	recordingURL string
}

//Voice Constructor
func (m *mocean) Voice() *voiceService {
	return &voiceService{
		m,
		"/voice/dial",
		"/voice/hangup",
		"/voice/rec",
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
//For more info, see docs: https://moceanapi.com/docs/#make-an-outbound-call
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

type hangupResponse struct {
	abstractResponse
}

//Hangup
//for more info, see docs: https://moceanapi.com/docs/#hangup-a-call
func (s *voiceService) Hangup(callUuid string) (response *hangupResponse, err error) {
	res, err := s.client.post(s.hangupURL, url.Values{"mocean-call-uuid": {callUuid}})
	if err != nil {
		return response, err
	}

	response = new(hangupResponse)
	err = json.Unmarshal(res, response)

	response.rawResponse = string(res)
	return response, err
}

type recordingResponse struct {
	Filename        string
	RecordingBuffer *bytes.Reader
}

//Recording
//for more info, see docs: https://moceanapi.com/docs/#download-a-recording
func (s *voiceService) Recording(callUuid string) (response *recordingResponse, err error) {
	res, err := s.client.get(s.recordingURL, url.Values{"mocean-call-uuid": {callUuid}})
	if err != nil {
		return response, err
	}

	return &recordingResponse{
		Filename:        callUuid + ".mp3",
		RecordingBuffer: bytes.NewReader(res),
	}, err
}
