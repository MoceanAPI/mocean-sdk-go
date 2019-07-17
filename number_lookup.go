package moceansdk

import (
	"encoding/json"
	"net/url"
)

type NumberLookupService struct {
	client          *Mocean
	numberLookupUrl string
}

//Number Lookup Constructor
func (mocean *Mocean) NumberLookup() *NumberLookupService {
	return &NumberLookupService{
		mocean,
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

type NumberLookupResponse struct {
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
