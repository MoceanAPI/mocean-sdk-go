package moceansdk

import (
	"testing"
	"encoding/json"
)

func TestMakeTgSendText(t *testing.T) {
	exceptTgSendText := &TgSendText {
		Action: "send-telegram",
		From: TgFrom{
			Type: "bot_username",
			ID:   "bot ID",
		},
		To: TgTo{
			Type: "chat_id",
			ID:   "chat ID",
		},
		Content: TgContent{
			Type: "text",
			Text: "test text",
		},
	}
	AssertEqual(t, exceptTgSendText, MakeTgSendText("bot ID", "chat ID", "test text"))
}

func TestMakeTgSendAudio(t *testing.T) {
	exceptTgSendAudio := &TgSendAudio {
		Action: "send-telegram",
		From: TgFrom{
			Type: "bot_username",
			ID:   "bot ID",
		},
		To: TgTo{
			Type: "chat_id",
			ID:   "chat ID",
		},
		Content: TgContent{
			Type:         "audio",
			RichMediaUrl: "test url",
			Text:         "test text",
		},
	}

	AssertEqual(t, exceptTgSendAudio, MakeTgSendAudio("bot ID", "chat ID", "test url", "test text"))
}

func TestMakeTgSendAnimation(t *testing.T) {
	exceptTgSendAnimation := &TgSendAnimation{
		Action: "send-telegram",
		From: TgFrom{
			Type: "bot_username",
			ID:   "bot ID",
		},
		To: TgTo{
			Type: "chat_id",
			ID:   "chat ID",
		},
		Content: TgContent{
			Type:         "animation",
			RichMediaUrl: "test url",
			Text:         "test text",
		},
	}

	AssertEqual(t, exceptTgSendAnimation, MakeTgSendAnimation("bot ID", "chat ID", "test url", "test text"))
}

func TestTgSendDocument(t *testing.T) {
	exceptTgSendDocument := &TgSendDocument{
		Action: "send-telegram",
		From: TgFrom{
			Type: "bot_username",
			ID:   "bot ID",
		},
		To: TgTo{
			Type: "chat_id",
			ID:   "chat ID",
		},
		Content: TgContent{
			Type:         "document",
			RichMediaUrl: "test url",
			Text:         "test text",
		},
	}

	AssertEqual(t, exceptTgSendDocument, MakeTgSendDocument("bot ID", "chat ID", "test url", "test text"))
}


func TestMakeTgSendPhoto(t *testing.T) {
	exceptTgSendPhoto := &TgSendPhoto{
		Action: "send-telegram",
		From: TgFrom{
			Type: "bot_username",
			ID:   "bot ID",
		},
		To: TgTo{
			Type: "chat_id",
			ID:   "chat ID",
		},
		Content: TgContent{
			Type:         "photo",
			RichMediaUrl: "test url",
			Text:         "test text",
		},
	}

	AssertEqual(t, exceptTgSendPhoto, MakeTgSendPhoto("bot ID", "chat ID", "test url", "test text"))
}

func TestTgSendVideo(t *testing.T) {
	exceptTgSendVideo := &TgSendVideo{
		Action: "send-telegram",
		From: TgFrom{
			Type: "bot_username",
			ID:   "bot ID",
		},
		To: TgTo{
			Type: "chat_id",
			ID:   "chat ID",
		},
		Content: TgContent{
			Type:         "video",
			RichMediaUrl: "test url",
			Text:         "test text",
		},
	}

	AssertEqual(t, exceptTgSendVideo, MakeTgSendVideo("bot ID", "chat ID", "test url", "test text"))
}


func TestMakeSendSMS(t *testing.T) {
	exceptTgSendSMS := &SendSMS{
		Action: "send-sms",
		From: TgFrom{
			Type: "phone_num",
			ID:   "from num",
		},
		To: TgTo{
			Type: "phone_num",
			ID:   "to num",
		},
		Content: TgContent{
			Type:  "text",
			Text:  "test text",
		},
	}

	AssertEqual(t, exceptTgSendSMS, MakeSendSMS("from num", "to num", "test text"))
}

func TestTgRequestContact(t *testing.T) {
	exceptTgRequestContact := &TgRequestContact{
		Action: "send-telegram",
		From: TgFrom{
			Type: "bot_username",
			ID:   "bot ID",
		},
		To: TgTo{
			Type: "chat_id",
			ID:   "chat ID",
		},
		Content: TgContent{
			Type: "text",
			Text: "test text",
		},
		TgKeyboard: TgKeyboard{
			ButtonText:    "test button text",
			ButtonRequest: "contact",
		},
	}

	AssertEqual(t, exceptTgRequestContact, MakeTgRequestContact("bot ID", "chat ID", "test text", "test button text", "Share contact"))
}

func TestCommandMcBuilderService(t *testing.T) {
	mcBuilder := CommandNewMcBuilder()

	expected := []interface{}{MakeTgSendText("bot ID", "chat ID", "test text")}
	mcBuilder.Add(MakeTgSendText("bot ID", "chat ID", "test text"))
	expectedTgSendText, err := json.Marshal(expected)
	AssertNoError(t, err)
	actualTgSendText, err := mcBuilder.Build()
	AssertNoError(t, err)
	AssertEqual(t, string(expectedTgSendText), actualTgSendText)

	expected = append(expected, MakeTgSendAudio("bot ID", "chat ID", "test url", "test text"))
	mcBuilder.Add(MakeTgSendAudio("bot ID", "chat ID", "test url", "test text"))
	expectedTgSendAudio, err := json.Marshal(expected)
	AssertNoError(t, err)
	actualTgSendAudio, err := mcBuilder.Build()
	AssertNoError(t, err)
	AssertEqual(t, string(expectedTgSendAudio), actualTgSendAudio)

	expected = append(expected, MakeTgSendVideo("bot ID", "chat ID", "test url", "test text"))
	mcBuilder.Add(MakeTgSendVideo("bot ID", "chat ID", "test url", "test text"))
	expectedTgSendVideo, err := json.Marshal(expected)
	AssertNoError(t, err)
	actualTgSendvideo, err := mcBuilder.Build()
	AssertNoError(t, err)
	AssertEqual(t, string(expectedTgSendVideo), actualTgSendvideo)
}

