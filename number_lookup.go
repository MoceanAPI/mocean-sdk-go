package moceansdk

import (
	"encoding/json"
	"net/url"
)

type NumberLookupService struct {
	client          *Mocean
	numberLookupURL string
}

//Number Lookup Constructor
func (m *Mocean) NumberLookup() *NumberLookupService {
	return &NumberLookupService{
		m,
		"/nl",
	}
}

type carrier struct {
	Country     interface{} `json:"country"`
	Name        interface{} `json:"name"`
	NetworkCode interface{} `json:"network_code"`
	Mcc         interface{} `json:"mcc"`
	Mnc         interface{} `json:"mnc"`
}

type numberLookupResponse struct {
	abstractResponse
	Msgid           interface{} `json:"msgid"`
	To              interface{} `json:"to"`
	CurrentCarrier  *carrier    `json:"current_carrier"`
	OriginalCarrier *carrier    `json:"original_carrier"`
	Ported          interface{} `json:"ported"`
	Reachable       interface{} `json:"reachable"`
}

//Request Number Lookup
//For more info, see docs: https://moceanapi.com/docs/#request-number-lookup
func (s *NumberLookupService) Inquiry(params url.Values) (response *numberLookupResponse, err error) {
	res, err := s.client.post(s.numberLookupURL, params)
	if err != nil {
		return response, err
	}

	response = new(numberLookupResponse)
	err = json.Unmarshal(res, response)

	response.rawResponse = string(res)
	return response, err
}
