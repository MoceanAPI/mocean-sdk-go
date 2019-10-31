package moceansdk

import (
	"encoding/json"
	"testing"
)

func TestMakeMcccDial(t *testing.T) {
	expectedDial := &McccDial{
		Action: "dial",
		To:     "test to",
	}

	AssertEqual(t, expectedDial, MakeMcccDial("test to"))
}

func TestMakeMcccCollect(t *testing.T) {
	expectedCollect := &McccCollect{
		Action:   "collect",
		EventURL: "http://test.com",
	}

	AssertEqual(t, expectedCollect, MakeMcccCollect("http://test.com"))
}

func TestMakeMcccPlay(t *testing.T) {
	expectedPlay := &McccPlay{
		Action: "play",
		File:   "http://test.com",
	}

	AssertEqual(t, expectedPlay, MakeMcccPlay("http://test.com"))
}

func TestMakeMcccSay(t *testing.T) {
	expectedSay := &McccSay{
		Action:   "say",
		Language: "en-US",
		Text:     "testing text",
	}

	AssertEqual(t, expectedSay, MakeMcccSay("testing text"))
}

func TestMakeMcccSleep(t *testing.T) {
	expectedSleep := &McccSleep{
		Action:   "sleep",
		Duration: 5000,
	}

	AssertEqual(t, expectedSleep, MakeMcccSleep(5000))
}

func TestMcccBuilderService(t *testing.T) {
	mcccBuilder := NewMcccBuilder()

	expected := []interface{}{MakeMcccDial("test to")}
	mcccBuilder.Add(MakeMcccDial("test to"))
	expectedDial, err := json.Marshal(expected)
	AssertNoError(t, err)
	actualDial, err := mcccBuilder.Build()
	AssertNoError(t, err)
	AssertEqual(t, string(expectedDial), actualDial)

	expected = append(expected, MakeMcccCollect("http://test.com"))
	mcccBuilder.Add(MakeMcccCollect("http://test.com"))
	expectedCollect, err := json.Marshal(expected)
	AssertNoError(t, err)
	actualCollect, err := mcccBuilder.Build()
	AssertNoError(t, err)
	AssertEqual(t, string(expectedCollect), actualCollect)

	expected = append(expected, MakeMcccPlay("http://test.com"))
	mcccBuilder.Add(MakeMcccPlay("http://test.com"))
	expectedPlay, err := json.Marshal(expected)
	AssertNoError(t, err)
	actualPlay, err := mcccBuilder.Build()
	AssertNoError(t, err)
	AssertEqual(t, string(expectedPlay), actualPlay)

	expected = append(expected, MakeMcccSay("testing text"))
	mcccBuilder.Add(MakeMcccSay("testing text"))
	expectedSay, err := json.Marshal(expected)
	AssertNoError(t, err)
	actualSay, err := mcccBuilder.Build()
	AssertNoError(t, err)
	AssertEqual(t, string(expectedSay), actualSay)

	expected = append(expected, MakeMcccSleep(5000))
	mcccBuilder.Add(MakeMcccSleep(5000))
	expectedSleep, err := json.Marshal(expected)
	AssertNoError(t, err)
	actualSleep, err := mcccBuilder.Build()
	AssertNoError(t, err)
	AssertEqual(t, string(expectedSleep), actualSleep)
}
