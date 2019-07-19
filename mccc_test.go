package moceansdk

import (
	"encoding/json"
	"testing"
)

func TestMakeMcccBridge(t *testing.T) {
	expectedBridge := &McccBridge{
		"dial",
		"test to",
	}

	AssertEqual(t, expectedBridge, MakeMcccBridge("test to"))
}

func TestMakeMcccCollect(t *testing.T) {
	expectedCollect := &McccCollect{
		"collect",
		"http://test.com",
		1,
		10,
		"#",
		5000,
	}

	AssertEqual(t, expectedCollect, MakeMcccCollect("http://test.com"))
}

func TestMakeMcccPlay(t *testing.T) {
	expectedPlay := &McccPlay{
		"play",
		"http://test.com",
		false,
	}

	AssertEqual(t, expectedPlay, MakeMcccPlay("http://test.com"))
}

func TestMakeMcccSay(t *testing.T) {
	expectedSay := &McccSay{
		"say",
		"en-US",
		"testing text",
		false,
	}

	AssertEqual(t, expectedSay, MakeMcccSay("testing text"))
}

func TestMakeMcccSleep(t *testing.T) {
	expectedSleep := &McccSleep{
		"sleep",
		5000,
		false,
	}

	AssertEqual(t, expectedSleep, MakeMcccSleep(5000))
}

func TestMcccBuilderService(t *testing.T) {
	mcccBuilder := NewMcccBuilder()

	expected := []interface{}{MakeMcccBridge("test to")}
	mcccBuilder.Add(MakeMcccBridge("test to"))
	expectedBridge, err := json.Marshal(expected)
	AssertNoError(t, err)
	actualBridge, err := mcccBuilder.Build()
	AssertNoError(t, err)
	AssertEqual(t, string(expectedBridge), actualBridge)

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
