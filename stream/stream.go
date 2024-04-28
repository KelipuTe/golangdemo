package stream

import (
	"encoding/binary"
)

type Msg struct {
	MsgLen  int    //消息长度
	Msg     []byte //消息体
	PayLoad string //有效载荷
}

func NewMsg() *Msg {
	return &Msg{}
}

func (t *Msg) encode() ([]byte, error) {
	payLoadLen := len(t.PayLoad)
	if payLoadLen <= 0 {
		return nil, ErrParseFailed
	}
	msgSlice := make([]byte, 0, 4+payLoadLen)

	//把消息长度转换成大端字节序格式
	msgLenBig := make([]byte, 4)
	binary.BigEndian.PutUint32(msgLenBig, uint32(4+payLoadLen))

	msgSlice = append(msgSlice, msgLenBig...)         //最前面的4个字节是消息长度
	msgSlice = append(msgSlice, []byte(t.PayLoad)...) //从第5个字节开始，后面都是请求体

	return msgSlice, nil
}

func (t *Msg) decode(buffer []byte, bufferLen int) error {
	if bufferLen <= 4 {
		return ErrParseFailed
	}

	//把大端字节序格式的消息长度转换回来
	msgLenBig := buffer[0:4]
	msgLen := int(binary.BigEndian.Uint32(msgLenBig))
	if bufferLen < msgLen {
		return ErrParseFailed
	}

	t.MsgLen = msgLen
	t.Msg = buffer[0:msgLen]
	t.PayLoad = string(buffer[4:msgLen])

	return nil
}
