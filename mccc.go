package moceansdk

import (
	"encoding/json"
	"reflect"
)

type McccBridge struct {
	Action string `json:"action"`
	To     string `json:"to"`
}

type McccCollect struct {
	Action      string `json:"action"`
	EventUrl    string `json:"event-url"`
	Min         int    `json:"min"`
	Max         int    `json:"max"`
	Terminators string `json:"terminators"`
	Timeout     int    `json:"timeout"`
}

type McccPlay struct {
	Action  string `json:"action"`
	File    string `json:"file"`
	BargeIn bool   `json:"barge-in"`
}

type McccSay struct {
	Action   string `json:"action"`
	Language string `json:"language"`
	Text     string `json:"text"`
	BargeIn  bool   `json:"barge-in"`
}

type McccSleep struct {
	Action   string `json:"action"`
	Duration int    `json:"duration"`
	BargeIn  bool   `json:"barge-in"`
}

// simple interface to make mccc
func MakeMcccBridge(to string) *McccBridge {
	return &McccBridge{
		"dial",
		to,
	}
}

func MakeMcccCollect(eventUrl string) *McccCollect {
	return &McccCollect{
		"collect",
		eventUrl,
		1,
		10,
		"#",
		5000,
	}
}

func MakeMcccPlay(file string) *McccPlay {
	return &McccPlay{
		Action: "play",
		File:   file,
	}
}

func MakeMcccSay(text string) *McccSay {
	return &McccSay{
		Action:   "say",
		Text:     text,
		Language: "en-US",
	}
}

func MakeMcccSleep(duration int) *McccSleep {
	return &McccSleep{
		Action:   "sleep",
		Duration: duration,
	}
}

// builder
type mcccBuilderService struct {
	mcccObjects []interface{}
}

func NewMcccBuilder() *mcccBuilderService {
	return &mcccBuilderService{}
}

func (s *mcccBuilderService) Add(mccc interface{}) *mcccBuilderService {
	mcccType := reflect.TypeOf(mccc)
	if mcccType == reflect.TypeOf(&McccBridge{}) {
		mcccBridge := mccc.(*McccBridge)
		mcccBridge.Action = "dial"
		s.mcccObjects = append(s.mcccObjects, mcccBridge)
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

func (s *mcccBuilderService) Build() (string, error) {
	converted, err := json.Marshal(s.mcccObjects)
	return string(converted), err
}
