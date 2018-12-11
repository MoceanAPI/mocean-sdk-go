package moceango

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type account struct {
	*Mocean
	GetBalanceUrl string
	GetPricingUrl string
}

//Account Constructor
func (mocean *Mocean) Account() *account {
	return &account{
		mocean,
		mocean.BaseUrl + "/account/balance",
		mocean.BaseUrl + "/account/pricing",
	}
}

type BalanceResponse struct {
	Status  int     `json:"status"`
	Balance float64 `json:"value"`
}

//Get Account Balance
//For more info, see docs: https://moceanapi.com/docs/#get-balance
func (account *account) getBalance() (balanceResponse *BalanceResponse, err error) {
	formData := account.makeFormData(account.ApiKey, account.ApiSecret);
	res, err := account.get(account.GetBalanceUrl + "?" + formData.Encode())
	if err != nil {
		return balanceResponse, err
	}

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return balanceResponse, err
	}

	if res.StatusCode != http.StatusAccepted {
		errRes := new(ErrorResponse)
		err = json.Unmarshal(responseBody, errRes)

		return balanceResponse, errors.New(errRes.ErrorMsg)
	}

	balanceResponse = new(BalanceResponse)
	err = json.Unmarshal(responseBody, balanceResponse)

	return balanceResponse, err
}

type PricingResponse struct {
	Status       int `json:"status"`
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
func (account *account) getPricing() (pricingResponse *PricingResponse, err error) {
	formData := account.makeFormData(account.ApiKey, account.ApiSecret);
	res, err := account.get(account.GetPricingUrl + "?" + formData.Encode())
	if err != nil {
		return pricingResponse, err
	}

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return pricingResponse, err
	}

	if res.StatusCode != http.StatusAccepted {
		errRes := new(ErrorResponse)
		err = json.Unmarshal(responseBody, errRes)

		return pricingResponse, errors.New(errRes.ErrorMsg)
	}

	pricingResponse = new(PricingResponse)
	err = json.Unmarshal(responseBody, pricingResponse)

	return pricingResponse, err
}
