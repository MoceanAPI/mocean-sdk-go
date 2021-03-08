package moceansdk

import (
	"testing"
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

