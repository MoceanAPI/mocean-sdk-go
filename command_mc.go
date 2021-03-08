package moceansdk

import (
	"encoding/json"
)

//TgSendText
type TgSendText struct {
	Action  string    `json:"action"`
	From    TgFrom    `json:"from"`
	To      TgTo      `json:"to"`
	Content TgContent `json:"content"`
}

type TgSendAudio struct {
	Action  string    `json:"action"`
	From    TgFrom    `json:"from"`
	To      TgTo      `json:"to"`
	Content TgContent `json:"content"`
}

type TgSendAnimation struct {
	Action  string    `json:"action"`
	From    TgFrom    `json:"from"`
	To      TgTo      `json:"to"`
	Content TgContent `json:"content"`
}

type TgSendPhoto struct {
	Action  string    `json:"action"`
	From    TgFrom    `json:"from"`
	To      TgTo      `json:"to"`
	Content TgContent `json:"content"`
}

type TgSendDocument struct {
	Action  string    `json:"action"`
	From    TgFrom    `json:"from"`
	To      TgTo      `json:"to"`
	Content TgContent `json:"content"`
}

type TgSendVideo struct {
	Action  string    `json:"action"`
	From    TgFrom    `json:"from"`
	To      TgTo      `json:"to"`
	Content TgContent `json:"content"`
}

type SendSMS struct {
	Action  string    `json:"action"`
	From    TgFrom    `json:"from"`
	To      TgTo      `json:"to"`
	Content TgContent `json:"content"`
}

type TgRequestContact struct {
	Action     string     `json:"action"`
	From       TgFrom     `json:"from"`
	To         TgTo       `json:"to"`
	Content    TgContent  `json:"content"`
	TgKeyboard TgKeyboard `json:tg_keyboard`
}

type TgFrom struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type TgTo struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type TgContent struct {
	Type         string      `json:"type"`
	Text         interface{} `json:"text"`
	RichMediaUrl interface{} `json:"rich_media_url,omitempty"`
}

type TgKeyboard struct {
	ButtonRequest string `json:"button_request"`
	ButtonText    string `json:"button_text"`
}

//MakeTgSendText
func MakeTgSendText(from string, to string, text string) *TgSendText {
	return &TgSendText{
		Action:  "send-telegram",
		From:    TgFrom{Type: "bot_username", ID: from},
		To:      TgTo{Type: "chat_id", ID: to},
		Content: TgContent{Type: "text", Text: text},
	}
}

//MakeTgSendAudio
func MakeTgSendAudio(from string, to string, url string, opt ...string) *TgSendAudio {

	text := ""

	if len(opt) > 0 {
		text = opt[0]
	}

	return &TgSendAudio{
		Action:  "send-telegram",
		From:    TgFrom{Type: "bot_username", ID: from},
		To:      TgTo{Type: "chat_id", ID: to},
		Content: TgContent{Type: "audio", RichMediaUrl: url, Text: text},
	}
}

//MakeTgSendAnimation
func MakeTgSendAnimation(from string, to string, url string, opt ...string) *TgSendAnimation {

	text := ""

	if len(opt) > 0 {
		text = opt[0]
	}

	return &TgSendAnimation{
		Action:  "send-telegram",
		From:    TgFrom{Type: "bot_username", ID: from},
		To:      TgTo{Type: "chat_id", ID: to},
		Content: TgContent{Type: "animation", RichMediaUrl: url, Text: text},
	}
}

//MakeTgSendDocument
func MakeTgSendDocument(from string, to string, url string, opt ...string) *TgSendDocument {

	text := ""

	if len(opt) > 0 {
		text = opt[0]
	}

	return &TgSendDocument{
		Action:  "send-telegram",
		From:    TgFrom{Type: "bot_username", ID: from},
		To:      TgTo{Type: "chat_id", ID: to},
		Content: TgContent{Type: "document", RichMediaUrl: url, Text: text},
	}
}

//MakeTgSendPhoto
func MakeTgSendPhoto(from string, to string, url string, opt ...string) *TgSendPhoto {

	text := ""

	if len(opt) > 0 {
		text = opt[0]
	}

	return &TgSendPhoto{
		Action:  "send-telegram",
		From:    TgFrom{Type: "bot_username", ID: from},
		To:      TgTo{Type: "chat_id", ID: to},
		Content: TgContent{Type: "photo", RichMediaUrl: url, Text: text},
	}
}

//MakeTgSendVideo
func MakeTgSendVideo(from string, to string, url string, opt ...string) *TgSendVideo {

	text := ""

	if len(opt) > 0 {
		text = opt[0]
	}

	return &TgSendVideo{
		Action:  "send-telegram",
		From:    TgFrom{Type: "bot_username", ID: from},
		To:      TgTo{Type: "chat_id", ID: to},
		Content: TgContent{Type: "video", RichMediaUrl: url, Text: text},
	}
}

func MakeSendSMS(from string, to string, text string) *SendSMS {

	return &SendSMS{
		Action:  "send-sms",
		From:    TgFrom{Type: "phone_num", ID: from},
		To:      TgTo{Type: "phone_num", ID: to},
		Content: TgContent{Type: "text", Text: text},
	}
}

//MakeTgRequestContact
func MakeTgRequestContact(from string, to string, text string, opt ...string) *TgRequestContact {

	buttonText := "Share Contact"

	if len(opt) > 0 {
		buttonText = opt[0]
	}

	return &TgRequestContact{
		Action:     "send-telegram",
		From:       TgFrom{Type: "bot_username", ID: from},
		To:         TgTo{Type: "chat_id", ID: to},
		Content:    TgContent{Type: "text", Text: text},
		TgKeyboard: TgKeyboard{ButtonRequest: "contact", ButtonText: buttonText},
	}
}

// func jsonEscape(i string) string {
// 	b, err := json.Marshal(i)
// 	if err != nil {
// 		panic(err)
// 	}
// 	s := string(b)
// 	return s[1 : len(s)-1]
// }

// builder
type CommandMcBuilderService struct {
	mcObjects []interface{}
}

//NewMcBuilder
func CommandNewMcBuilder() *CommandMcBuilderService {
	return &CommandMcBuilderService{}
}

//Add
func (s *CommandMcBuilderService) Add(mc interface{}) *CommandMcBuilderService {
	s.mcObjects = append(s.mcObjects, mc)
	return s
}

func (s *CommandMcBuilderService) Build() (string, error) {
	converted, err := json.Marshal(s.mcObjects)
	return string(converted), err
}
