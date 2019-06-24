package moceansdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type Options struct {
	BaseUrl    string
	Version    string
	HttpClient *http.Client
}

type Mocean struct {
	Options    *Options
	apiKey     string
	apiSecret  string
	sdkVersion string
}

type abstractResponse struct {
	Status      interface{} `json:"status"`
	rawResponse string
}

func (res *abstractResponse) String() string {
	return res.rawResponse
}

type ErrorResponse struct {
	abstractResponse
	ErrorMsg interface{} `json:"err_msg"`
}

func NewMoceanClient(apiKey, apiSecret string) *Mocean {
	return &Mocean{
		Options: &Options{
			BaseUrl: "https://rest.moceanapi.com",
			Version: "2",
			HttpClient: &http.Client{
				Timeout: time.Second * 30,
			},
		},
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		sdkVersion: "2.0.0",
	}
}

func (mocean *Mocean) SdkVersion() string {
	return mocean.sdkVersion
}

func (mocean *Mocean) post(url string, formData url.Values) ([]byte, error) {
	formData = mocean.setAuth(formData)
	req, err := http.NewRequest("POST", mocean.Options.BaseUrl+"/rest/"+mocean.Options.Version+url, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := mocean.Options.HttpClient.Do(req);
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusAccepted {
		errRes := new(ErrorResponse)
		err = json.Unmarshal(responseBody, errRes)

		return nil, errors.New(errRes.ErrorMsg)
	}

	return responseBody, nil
}

func (mocean *Mocean) get(url string, formData url.Values) ([]byte, error) {
	formData = mocean.setAuth(formData)
	req, err := http.NewRequest("GET", mocean.Options.BaseUrl+"/rest/"+mocean.Options.Version+url+"?"+formData.Encode(), nil)
	if err != nil {
		return nil, err
	}

	res, err := mocean.Options.HttpClient.Do(req);
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusAccepted {
		errRes := new(ErrorResponse)
		err = json.Unmarshal(responseBody, errRes)

		return nil, errors.New(errRes.ErrorMsg)
	}

	return responseBody, nil
}

func (mocean *Mocean) setAuth(data url.Values) url.Values {
	data.Set("mocean-api-key", mocean.apiKey)
	data.Set("mocean-api-secret", mocean.apiSecret)
	data.Set("mocean-resp-format", "JSON")
	data.Set("mocean-medium", "GO-SDK")

	return data
}

func (mocean *Mocean) structToMap(i interface{}, inputValues url.Values) url.Values {
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		if fmt.Sprint(iVal.Field(i)) != "" {
			inputValues.Set(typ.Field(i).Name, fmt.Sprint(iVal.Field(i)))
		}
	}
	return inputValues
}
