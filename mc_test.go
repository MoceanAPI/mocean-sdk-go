package moceansdk

import (
	"encoding/json"
	"testing"
)

func TestMakeMcDial(t *testing.T) {
	expectedDial := &McDial{
		Action: "dial",
		To:     "test to",
	}

	AssertEqual(t, expectedDial, MakeMcDial("test to"))
}

func TestMakeMcCollect(t *testing.T) {
	expectedCollect := &McCollect{
		Action:   "collect",
		EventURL: "http://test.com",
	}

	AssertEqual(t, expectedCollect, MakeMcCollect("http://test.com"))
}

func TestMakeMcPlay(t *testing.T) {
	expectedPlay := &McPlay{
		Action: "play",
		File:   "http://test.com",
	}

	AssertEqual(t, expectedPlay, MakeMcPlay("http://test.com"))
}

func TestMakeMcSay(t *testing.T) {
	expectedSay := &McSay{
		Action:   "say",
		Language: "en-US",
		Text:     "testing text",
	}

	AssertEqual(t, expectedSay, MakeMcSay("testing text"))
}

func TestMakeMcSleep(t *testing.T) {
	expectedSleep := &McSleep{
		Action:   "sleep",
		Duration: 5000,
	}

	AssertEqual(t, expectedSleep, MakeMcSleep(5000))
}

func TestMakeMcRecord(t *testing.T) {
	expectedRecord := &McRecord{
		Action: "record",
	}

	AssertEqual(t, expectedRecord, MakeMcRecord())
}

func TestMcBuilderService(t *testing.T) {
	mcBuilder := NewMcBuilder()

	expected := []interface{}{MakeMcDial("test to")}
	mcBuilder.Add(MakeMcDial("test to"))
	expectedDial, err := json.Marshal(expected)
	AssertNoError(t, err)
	actualDial, err := mcBuilder.Build()
	AssertNoError(t, err)
	AssertEqual(t, string(expectedDial), actualDial)

	expected = append(expected, MakeMcCollect("http://test.com"))
	mcBuilder.Add(MakeMcCollect("http://test.com"))
	expectedCollect, err := json.Marshal(expected)
	AssertNoError(t, err)
	actualCollect, err := mcBuilder.Build()
	AssertNoError(t, err)
	AssertEqual(t, string(expectedCollect), actualCollect)

	expected = append(expected, MakeMcPlay("http://test.com"))
	mcBuilder.Add(MakeMcPlay("http://test.com"))
	expectedPlay, err := json.Marshal(expected)
	AssertNoError(t, err)
	actualPlay, err := mcBuilder.Build()
	AssertNoError(t, err)
	AssertEqual(t, string(expectedPlay), actualPlay)

	expected = append(expected, MakeMcSay("testing text"))
	mcBuilder.Add(MakeMcSay("testing text"))
	expectedSay, err := json.Marshal(expected)
	AssertNoError(t, err)
	actualSay, err := mcBuilder.Build()
	AssertNoError(t, err)
	AssertEqual(t, string(expectedSay), actualSay)

	expected = append(expected, MakeMcSleep(5000))
	mcBuilder.Add(MakeMcSleep(5000))
	expectedSleep, err := json.Marshal(expected)
	AssertNoError(t, err)
	actualSleep, err := mcBuilder.Build()
	AssertNoError(t, err)
	AssertEqual(t, string(expectedSleep), actualSleep)

	expected = append(expected, MakeMcRecord())
	mcBuilder.Add(MakeMcRecord())
	expectedRecord, err := json.Marshal(expected)
	AssertNoError(t, err)
	actualRecord, err := mcBuilder.Build()
	AssertNoError(t, err)
	AssertEqual(t, string(expectedRecord), actualRecord)
}
