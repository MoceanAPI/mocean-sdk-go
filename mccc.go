package moceansdk

import (
	"encoding/json"
	"reflect"
)

//McccDial
type McccDial struct {
	Action           string `json:"action"`
	To               string `json:"to"`
	From             string `json:"from"`
	DialSequentially bool   `json:"dial-sequentially"`
}

//McccCollect
type McccCollect struct {
	Action      string `json:"action"`
	EventURL    string `json:"event-url"`
	Min         int    `json:"min"`
	Max         int    `json:"max"`
	Terminators string `json:"terminators"`
	Timeout     int    `json:"timeout"`
}

//McccPlay
type McccPlay struct {
	Action          string `json:"action"`
	File            string `json:"file"`
	BargeIn         bool   `json:"barge-in"`
	ClearDigitCache bool   `json:"clear-digit-cache"`
}

//McccSay
type McccSay struct {
	Action          string `json:"action"`
	Language        string `json:"language"`
	Text            string `json:"text"`
	BargeIn         bool   `json:"barge-in"`
	ClearDigitCache bool   `json:"clear-digit-cache"`
}

//McccSleep
type McccSleep struct {
	Action   string `json:"action"`
	Duration int    `json:"duration"`
}

//McccRecord
type McccRecord struct {
	Action string `json:"action"`
}

// simple interface to make mccc
//MakeMcccDial
func MakeMcccDial(to string) *McccDial {
	return &McccDial{
		Action: "dial",
		To:     to,
	}
}

//MakeMcccCollect
func MakeMcccCollect(eventURL string) *McccCollect {
	return &McccCollect{
		Action:   "collect",
		EventURL: eventURL,
	}
}

//MakeMcccPlay
func MakeMcccPlay(file string) *McccPlay {
	return &McccPlay{
		Action: "play",
		File:   file,
	}
}

//MakeMcccSay
func MakeMcccSay(text string) *McccSay {
	return &McccSay{
		Action:   "say",
		Text:     text,
		Language: "en-US",
	}
}

//MakeMcccSleep
func MakeMcccSleep(duration int) *McccSleep {
	return &McccSleep{
		Action:   "sleep",
		Duration: duration,
	}
}

//MakeMcccRecord
func MakeMcccRecord() *McccRecord {
	return &McccRecord{
		Action: "record",
	}
}

// builder
type McccBuilderService struct {
	mcccObjects []interface{}
}

//NewMcccBuilder
func NewMcccBuilder() *McccBuilderService {
	return &McccBuilderService{}
}

func (s *McccBuilderService) Add(mccc interface{}) *McccBuilderService {
	mcccType := reflect.TypeOf(mccc)
	if mcccType == reflect.TypeOf(&McccDial{}) {
		mcccDial := mccc.(*McccDial)
		mcccDial.Action = "dial"
		s.mcccObjects = append(s.mcccObjects, mcccDial)
	} else if mcccType == reflect.TypeOf(&McccCollect{}) {
		mcccCollect := mccc.(*McccCollect)
		mcccCollect.Action = "collect"
		s.mcccObjects = append(s.mcccObjects, mcccCollect)
	} else if mcccType == reflect.TypeOf(&McccPlay{}) {
		mcccPlay := mccc.(*McccPlay)
		mcccPlay.Action = "play"
		s.mcccObjects = append(s.mcccObjects, mcccPlay)
	} else if mcccType == reflect.TypeOf(&McccSay{}) {
		mcccSay := mccc.(*McccSay)
		mcccSay.Action = "say"
		s.mcccObjects = append(s.mcccObjects, mcccSay)
	} else if mcccType == reflect.TypeOf(&McccSleep{}) {
		mcccSleep := mccc.(*McccSleep)
		mcccSleep.Action = "sleep"
		s.mcccObjects = append(s.mcccObjects, mcccSleep)
	}
	return s
}

func (s *McccBuilderService) Build() (string, error) {
	converted, err := json.Marshal(s.mcccObjects)
	return string(converted), err
}
