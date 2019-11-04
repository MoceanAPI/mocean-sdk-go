package moceansdk

import (
	"errors"
	"github.com/google/go-cmp/cmp"
	"github.com/jarcoal/httpmock"
	"io/ioutil"
	"net/url"
	"os"
	"testing"
)

var _mocean *mocean

func TestMain(m *testing.M) {
	os.Exit(func() int {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		_mocean = NewMoceanClient("test api key", "test api secret")
		return m.Run()
	}())
}

func ReadResourceFile(fileName string) string {
	file, _ := os.Open("resources/" + fileName)
	defer file.Close()

	fileContent, _ := ioutil.ReadAll(file)
	return string(fileContent)
}

func AssertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if !cmp.Equal(expected, actual) {
		t.Errorf("failed to assert that two things are equals.\nExpected : %v\nActual   : %v", expected, actual)
	}
}

func AssertTrue(t *testing.T, object bool) {
	if object != true {
		t.Errorf("failed to assert that object is true.\nObject : %v", object)
	}
}

func AssertNotNil(t *testing.T, object interface{}) {
	if object == nil {
		t.Errorf("failed to assert that object is not nil.\nObject : %v", object)
	}
}

func AssertError(t *testing.T, err error) {
	if err == nil {
		t.Error("failed to assert that function call has error.")
	}
}

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("failed to assert that function call has no errors.\nError : %v", err)
	}
}

func TestKeySecretAuth(t *testing.T) {
	AssertEqual(t, "test api key", _mocean.apiKey)
	AssertEqual(t, "test api secret", _mocean.apiSecret)
}

func TestSdkVersion(t *testing.T) {
	AssertNotNil(t, SdkVersion())
}

func TestInvalidMethod(t *testing.T) {
	_, err := _mocean.makeRequest("]", "/test", url.Values{})
	AssertError(t, err)
}

func TestHttpClientError(t *testing.T) {
	httpmock.RegisterResponder("GET", _mocean.Options.BaseURL+"/rest/"+_mocean.Options.Version+"/test",
		httpmock.NewErrorResponder(errors.New("timeout")))

	_, err := _mocean.makeRequest("GET", "/test", url.Values{})
	AssertError(t, err)
}
