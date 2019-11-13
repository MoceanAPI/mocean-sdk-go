package moceansdk

import (
	"encoding/json"
	"reflect"
)

//McDial
type McDial struct {
	Action           string `json:"action"`
	To               string `json:"to"`
	From             string `json:"from"`
	DialSequentially bool   `json:"dial-sequentially"`
}

//McCollect
type McCollect struct {
	Action      string `json:"action"`
	EventURL    string `json:"event-url"`
	Min         int    `json:"min"`
	Max         int    `json:"max"`
	Terminators string `json:"terminators"`
	Timeout     int    `json:"timeout"`
}

//McPlay
type McPlay struct {
	Action          string `json:"action"`
	File            string `json:"file"`
	BargeIn         bool   `json:"barge-in"`
	ClearDigitCache bool   `json:"clear-digit-cache"`
}

//McSay
type McSay struct {
	Action          string `json:"action"`
	Language        string `json:"language"`
	Text            string `json:"text"`
	BargeIn         bool   `json:"barge-in"`
	ClearDigitCache bool   `json:"clear-digit-cache"`
}

//McSleep
type McSleep struct {
	Action   string `json:"action"`
	Duration int    `json:"duration"`
}

//McRecord
type McRecord struct {
	Action string `json:"action"`
}

// simple interface to make mc
//MakeMcDial
func MakeMcDial(to string) *McDial {
	return &McDial{
		Action: "dial",
		To:     to,
	}
}

//MakeMcCollect
func MakeMcCollect(eventURL string) *McCollect {
	return &McCollect{
		Action:   "collect",
		EventURL: eventURL,
	}
}

//MakeMcPlay
func MakeMcPlay(file string) *McPlay {
	return &McPlay{
		Action: "play",
		File:   file,
	}
}

//MakeMcSay
func MakeMcSay(text string) *McSay {
	return &McSay{
		Action:   "say",
		Text:     text,
		Language: "en-US",
	}
}

//MakeMcSleep
func MakeMcSleep(duration int) *McSleep {
	return &McSleep{
		Action:   "sleep",
		Duration: duration,
	}
}

//MakeMcRecord
func MakeMcRecord() *McRecord {
	return &McRecord{
		Action: "record",
	}
}

// builder
type McBuilderService struct {
	mcObjects []interface{}
}

//NewMcBuilder
func NewMcBuilder() *McBuilderService {
	return &McBuilderService{}
}

func (s *McBuilderService) Add(mc interface{}) *McBuilderService {
	mcType := reflect.TypeOf(mc)
	if mcType == reflect.TypeOf(&McDial{}) {
		mcDial := mc.(*McDial)
		mcDial.Action = "dial"
		s.mcObjects = append(s.mcObjects, mcDial)
	} else if mcType == reflect.TypeOf(&McCollect{}) {
		mcCollect := mc.(*McCollect)
		mcCollect.Action = "collect"
		s.mcObjects = append(s.mcObjects, mcCollect)
	} else if mcType == reflect.TypeOf(&McPlay{}) {
		mcPlay := mc.(*McPlay)
		mcPlay.Action = "play"
		s.mcObjects = append(s.mcObjects, mcPlay)
	} else if mcType == reflect.TypeOf(&McSay{}) {
		mcSay := mc.(*McSay)
		mcSay.Action = "say"
		s.mcObjects = append(s.mcObjects, mcSay)
	} else if mcType == reflect.TypeOf(&McSleep{}) {
		mcSleep := mc.(*McSleep)
		mcSleep.Action = "sleep"
		s.mcObjects = append(s.mcObjects, mcSleep)
	} else if mcType == reflect.TypeOf(&McRecord{}) {
		mcRecord := mc.(*McRecord)
		mcRecord.Action = "record"
		s.mcObjects = append(s.mcObjects, mcRecord)
	}
	return s
}

func (s *McBuilderService) Build() (string, error) {
	converted, err := json.Marshal(s.mcObjects)
	return string(converted), err
}
