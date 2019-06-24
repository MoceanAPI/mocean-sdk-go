package moceansdk

import (
	"encoding/json"
	"net/url"
)

type NumberLookupService struct {
	client          *Mocean
	numberLookupUrl string
}

//Verify Constructor
func (mocean *Mocean) NumberLookup() *NumberLookupService {
	return &NumberLookupService{
		mocean,
		"/nl",
	}
}

type carrier struct {
	Country     string `json:"country"`
	Name        string `json:"name"`
	NetworkCode int    `json:"network_code"`
	Mcc         string `json:"mcc"`
	Mnc         string `json:"mnc"`
}

type NumberLookupResponse struct {
	abstractResponse
	Msgid           string   `json:"msgid"`
	To              string   `json:"to"`
	CurrentCarrier  *carrier `json:"current_carrier"`
	OriginalCarrier *carrier `json:"original_carrier"`
	Ported          string   `json:"ported"`
	Reachable       string   `json:"reachable"`
}

//Send verify code
//For more info, see docs: https://moceanapi.com/docs/#send-code
func (s *NumberLookupService) Inquiry(params url.Values) (numberLookupResponse *NumberLookupResponse, err error) {
	res, err := s.client.post(s.numberLookupUrl, params)
	if err != nil {
		return numberLookupResponse, err
	}

	numberLookupResponse = new(NumberLookupResponse)
	err = json.Unmarshal(res, numberLookupResponse)

	numberLookupResponse.rawResponse = string(res)
	return numberLookupResponse, err
}
