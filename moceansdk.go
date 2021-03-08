package moceansdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//SdkVersion
func SdkVersion() string {
	return "2.1.1-rc"
}

//Options
type options struct {
	BaseURL    string
	Version    string
	HTTPClient *http.Client
}

//Mocean
type mocean struct {
	Options   *options
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
		Options: &options{
			BaseURL: "https://rest.moceanapi.com",
			Version: "2",
			HTTPClient: &http.Client{
				Timeout: time.Second * 30,
			},
		},
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}

func (m *mocean) post(url string, formData url.Values) ([]byte, error) {
	return m.makeRequest("POST", url, formData)
}

func (m *mocean) get(url string, formData url.Values) ([]byte, error) {
	return m.makeRequest("GET", url, formData)
}

func (m *mocean) makeRequest(method string, url string, formData url.Values) ([]byte, error) {
	formData = m.setAuth(formData)

	var req *http.Request
	var newRequestErr error
	if method == "GET" {
		req, newRequestErr = http.NewRequest(method, m.Options.BaseURL+"/rest/"+m.Options.Version+url+"?"+formData.Encode(), nil)
	} else {
		req, newRequestErr = http.NewRequest(method, m.Options.BaseURL+"/rest/"+m.Options.Version+url, strings.NewReader(formData.Encode()))
	}
	if newRequestErr != nil {
		return nil, newRequestErr
	}

	if method == "POST" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	res, err := m.Options.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= http.StatusOK && res.StatusCode < 300 {
		return responseBody, nil
	}

	//error response
	errRes := new(errorResponse)
	err = json.Unmarshal(responseBody, errRes)

	return nil, fmt.Errorf("%v", errRes.ErrorMsg)
}

func (m *mocean) setAuth(data url.Values) url.Values {
	data.Set("mocean-api-key", m.apiKey)
	data.Set("mocean-api-secret", m.apiSecret)
	data.Set("mocean-resp-format", "JSON")
	data.Set("mocean-medium", "GO-SDK")

	return data
}
