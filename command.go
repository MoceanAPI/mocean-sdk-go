package moceansdk

import (
	"encoding/json"
	"net/url"
)

type commandService struct {
	client     *mocean
	commandURL string
}

//Voice Constructor
func (m *mocean) Command() *commandService {
	return &commandService{
		m,
		"/send-message",
	}
}

type commandResponse struct {
	abstractResponse
	SessionUUID       string              `json:"session_uuid"`
	MoceanCommandResp []moceanCommandResp `json:"mocean_command_resp"`
}

type moceanCommandResp struct {
	Action               interface{} `json:"action"`
	MessageID            interface{} `json:"message_id"`
	McPosition           interface{} `json:"mc_position"`
	TotalMessageSegments interface{} `json:"total_message_segments"`
}

func (s *commandService) Execute(params url.Values) (response *commandResponse, err error) {

	res, err := s.client.post(s.commandURL, params)
	if err != nil {
		return response, err
	}

	response = new(commandResponse)
	err = json.Unmarshal(res, response)

	response.rawResponse = string(res)
	return response, err
}
