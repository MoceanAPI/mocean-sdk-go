package moceansdk

import (
	"encoding/json"
	"net/url"
)

type accountService struct {
	client     *mocean
	balanceUrl string
	pricingUrl string
}

//Account Constructor
func (m *mocean) Account() *accountService {
	return &accountService{
		m,
		"/account/balance",
		"/account/pricing",
	}
}

type balanceResponse struct {
	abstractResponse
	Value interface{} `json:"value"`
}

//Get Account Balance
//For more info, see docs: https://moceanapi.com/docs/#get-balance
func (s *accountService) GetBalance(params url.Values) (response *balanceResponse, err error) {
	res, err := s.client.get(s.balanceUrl, params)
	if err != nil {
		return response, err
	}

	response = new(balanceResponse)
	err = json.Unmarshal(res, response)

	response.rawResponse = string(res)
	return response, err
}

type pricingResponse struct {
	abstractResponse
	Destinations []struct {
		Country  interface{} `json:"country"`
		Operator interface{} `json:"operator"`
		Mcc      interface{} `json:"mcc"`
		Mnc      interface{} `json:"mnc"`
		Price    interface{} `json:"price"`
		Currency interface{} `json:"currency"`
	} `json:"destinations"`
}

//Get Account Pricing
//For more info, see docs: https://moceanapi.com/docs/#account-pricing
func (s *accountService) GetPricing(params url.Values) (response *pricingResponse, err error) {
	res, err := s.client.get(s.pricingUrl, params)
	if err != nil {
		return response, err
	}

	response = new(pricingResponse)
	err = json.Unmarshal(res, response)

	response.rawResponse = string(res)
	return response, err
}
