package moceango

import (
	"testing"
)

func TestSendSMS(t *testing.T) {
	mocean := NewMoceanClient(testParams["apiKey"], testParams["apiSecret"])

	message := &Message{
		From: "Mocean",
		To:   "60123456789",
		Text: "testing",
	}

	res, err := mocean.Sms().Send(message)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Msgid: " + res.Messages[0].Msgid)

	testMsgStatusRes, err := mocean.Sms().getMsgStatus(res.Messages[0].Msgid)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("MsgStatus: %d", testMsgStatusRes.MessageStatus)
}
