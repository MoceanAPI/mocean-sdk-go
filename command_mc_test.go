package moceansdk

import (
	"encoding/json"
	"testing"
)

func TestMakeTgSendText (t *testing.T) {
	exceptTgSendText := &TgSendText{
		Action: "send-telegram"
		From: &TgFrom {
			Type: "bot_username"
			ID: "test ID"
		}
		To: &TgTo {
			Type: "chat_id"
			ID: "test ID"
		}
		Content: &TgContent {
			Type: "text"
			Text: "test text"
		}
		
	}
	AssertEqual(t, exceptTgSendText, MakeTgSendText("test ID","test ID","test text"))
}


func TestMakeTgSendAudio (t *testing.T) {
	var exceptTgSendAudio := &MakeTgSendAudio{
		Action:	"send-telegram"
		From: &TgFrom{
			Type: "bot_username"
			ID: "test ID"
		}
		To: &TgTo {
			Type: "chat_id"
			ID: "test ID"
		}
		Content: &TgContent {
			Type: "audio"
			RichMediaUrl: "test url"
			Text: "test text"
		}
	}
	
	AssertEqual(t,exceptTgSendAudio,MakeTgSendAudio("test ID", "test ID","test url","test text"))
}

func TestMakeTgSendAnimation (t *testing.T) 
{
	var exceptTgSendAnimation := &MakeTgSendAnimation{
		Action:	"send-telegram"
		From: &TgFrom{
			Type: "bot_username"
			ID: "test ID"
		}
		To: &TgTo {
			Type: "chat_id"
			ID: "test ID"
		}
		Content: &TgContent {
			Type: "animation"
			RichMediaUrl: "test url"
			Text: "test text"
		}
	}
	
	AssertEqual(t,exceptTgSendAnimation,MakeTgSendAudio("test ID", "test ID","test url","test text"))
}

func TestTgSendDocument (t *testing.T) {
	var exceptTgSendDocument := &MakeTgSendDocument{
		Action:	"send-telegram"
		From: &TgFrom{
			Type: "bot_username"
			ID: "test ID"
		}
		To: &TgTo {
			Type: "chat_id"
			ID: "test ID"
		}
		Content: &TgContent {
			Type: "document"
			RichMediaUrl: "test url"
			Text: "test text"
		}
	}
	
	AssertEqual(t,exceptTgSendAudio,MakeTgSendDocument("test ID", "test ID","test url","test text"))
}

func TestTgSendVideo (t *testing.T) {
	var exceptTgSendVideo := &MakeTgSendVideo{
		Action:	"send-telegram"
		From: &TgFrom{
			Type: "bot_username"
			ID: "test ID"
		}
		To: &TgTo {
			Type: "chat_id"
			ID: "test ID"
		}
		Content: &TgContent {
			Type: "video"
			RichMediaUrl: "test url"
			Text: "test text"
		}
	}
	
	AssertEqual(t,exceptTgSendVideo,MakeTgSendVideo("test ID", "test ID","test url","test text"))
}

func TestTgRequestContact (t *testing.T) {
	var exceptTgRequestContact := &MakeTgRequestContact{
		Action:	"send-telegram"
		From: &TgFrom{
			Type: "bot_username"
			ID: "test ID"
		}
		To: &TgTo {
			Type: "chat_id"
			ID: "test ID"
		}
		Content: &TgContent {
			Type: "text"
			Text: "test text"
		}
		TgKeyboard: &TgKeyboard{
			ButtonText: "test button text"
			ButtonRequest: "contact"
		}

	}
	
	AssertEqual(t,exceptTgRequestContact,MakeTgRequestContact("test ID", "test ID","test url","test text", "Share contact"))
}