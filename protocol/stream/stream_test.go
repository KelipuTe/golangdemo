package stream

import (
	"log"
	"testing"
)

func Test_Encode(t *testing.T) {
	msg := NewMsg()
	msg.PayLoad = "123123"
	msgEncode, err := msg.encode()
	if err != nil {
		t.Error(err)
	}
	log.Println(msg.MsgLen)
	log.Println(msgEncode)
}

func Test_Decode(t *testing.T) {
	msg := NewMsg()
	msg.MsgLen = 10
	msg.Msg = []byte{0, 0, 0, 10, 49, 50, 51, 49, 50, 51}
	err := msg.decode(msg.Msg, 10)
	if err != nil {
		t.Error(err)
	}
	log.Println(msg.PayLoad)
}
