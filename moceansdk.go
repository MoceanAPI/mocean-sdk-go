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

func SdkVersion() string {
	return "2.0.2"
}

type Options struct {
	BaseUrl    string
	Version    string
	HttpClient *http.Client
}

type mocean struct {
	Options   *Options
	apiKey    string
	apiSecret string
}

type abstractResponse struct {
	Status      interface{} `json:"status"`
	rawResponse string
}

func (res *abstractResponse) String() string {
	return res.rawResponse
}

type errorResponse struct {
	abstractResponse
	ErrorMsg interface{} `json:"err_msg"`
}

func NewMoceanClient(apiKey, apiSecret string) *mocean {
	return &mocean{
		Options: &Options{
			BaseUrl: "https://rest.moceanapi.com",
			Version: "2",
			HttpClient: &http.Client{
				Timeout: time.Second * 30,
			},
		},
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}

func (m *mocean) post(url string, formData url.Values) ([]byte, error) {
	formData = m.setAuth(formData)
	req, err := http.NewRequest("POST", m.Options.BaseUrl+"/rest/"+m.Options.Version+url, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := m.Options.HttpClient.Do(req);
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusAccepted {
		errRes := new(errorResponse)
		err = json.Unmarshal(responseBody, errRes)

		return nil, errors.New(fmt.Sprintf("%v", errRes.ErrorMsg))
	}

	return responseBody, nil
}

func (m *mocean) get(url string, formData url.Values) ([]byte, error) {
	formData = m.setAuth(formData)
	req, err := http.NewRequest("GET", m.Options.BaseUrl+"/rest/"+m.Options.Version+url+"?"+formData.Encode(), nil)
	if err != nil {
		return nil, err
	}

	res, err := m.Options.HttpClient.Do(req);
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusAccepted {
		errRes := new(errorResponse)
		err = json.Unmarshal(responseBody, errRes)

		return nil, errors.New(fmt.Sprintf("%v", errRes.ErrorMsg))
	}

	return responseBody, nil
}

func (m *mocean) setAuth(data url.Values) url.Values {
	data.Set("mocean-api-key", m.apiKey)
	data.Set("mocean-api-secret", m.apiSecret)
	data.Set("mocean-resp-format", "JSON")
	data.Set("mocean-medium", "GO-SDK")

	return data
}

func (m *mocean) structToMap(i interface{}, inputValues url.Values) url.Values {
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		if fmt.Sprint(iVal.Field(i)) != "" {
			inputValues.Set(typ.Field(i).Name, fmt.Sprint(iVal.Field(i)))
		}
	}
	return inputValues
}
