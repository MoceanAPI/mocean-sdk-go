package moceansdk

import (
	"encoding/json"
	"net/url"
)

type AccountService struct {
	client     *Mocean
	balanceUrl string
	pricingUrl string
}

//Account Constructor
func (mocean *Mocean) Account() *AccountService {
	return &AccountService{
		mocean,
		"/account/balance",
		"/account/pricing",
	}
}

type BalanceResponse struct {
	abstractResponse
	Balance interface{} `json:"value"`
}

//Get Account Balance
//For more info, see docs: https://moceanapi.com/docs/#get-balance
func (s *AccountService) GetBalance(params url.Values) (balanceResponse *BalanceResponse, err error) {
	res, err := s.client.get(s.balanceUrl, params)
	if err != nil {
		return balanceResponse, err
	}

	balanceResponse = new(BalanceResponse)
	err = json.Unmarshal(res, balanceResponse)

	balanceResponse.rawResponse = string(res)
	return balanceResponse, err
}

type PricingResponse struct {
	abstractResponse
	Destinations []struct {
		Country  string  `json:"country"`
		Operator string  `json:"operator"`
		Mcc      string  `json:"mcc"`
		Mnc      string  `json:"mnc"`
		Price    float64 `json:"price"`
		Currency string  `json:"currency"`
	} `json:"destinations"`
}

//Get Account Pricing
//For more info, see docs: https://moceanapi.com/docs/#account-pricing
func (s *AccountService) GetPricing(params url.Values) (pricingResponse *PricingResponse, err error) {
	res, err := s.client.get(s.pricingUrl, params)
	if err != nil {
		return pricingResponse, err
	}

	pricingResponse = new(PricingResponse)
	err = json.Unmarshal(res, pricingResponse)

	pricingResponse.rawResponse = string(res)
	return pricingResponse, err
}
