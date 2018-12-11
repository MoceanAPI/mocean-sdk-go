package moceango

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type Mocean struct {
	ApiKey    string
	ApiSecret string

	BaseUrl    string
	HTTPClient *http.Client
}

type ErrorResponse struct {
	Status   int    `json:"status"`
	ErrorMsg string `json:"err_msg"`
}

const (
	baseURL       = "https://rest.moceanapi.com/rest/1"
	clientTimeout = time.Second * 30
)

func NewMoceanClient(apiKey, apiSecret string) *Mocean {
	client := &http.Client{
		Timeout: clientTimeout,
	}

	return &Mocean{
		ApiKey:     apiKey,
		ApiSecret:  apiSecret,
		BaseUrl:    baseURL,
		HTTPClient: client,
	}
}

func (mocean *Mocean) post(url string, formData url.Values) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return mocean.do(req);
}

func (mocean *Mocean) get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return mocean.do(req)
}

func (mocean *Mocean) do(req *http.Request) (*http.Response, error) {
	return mocean.HTTPClient.Do(req)
}

func (mocean *Mocean) makeFormData(apiKey, apiSecret string) url.Values {
	data := url.Values{}

	data.Set("mocean-api-key", apiKey)
	data.Set("mocean-api-secret", apiSecret)
	data.Set("mocean-resp-format", "JSON")

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
