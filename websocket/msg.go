package websocket

import (
	"encoding/json"
	"time"
)

type Msg struct {
	Addr string //请求IP和端口

	Fin uint8 //FIN，1 bit，0=不是消息的最后一个分片；1=这是消息的最后一个分片；
	//RSV1、RSV2、RSV3，各 1 bit，这里不处理。
	Opcode       uint8   //OPCODE，4 bit
	Mask         uint8   //MASK，1 bit，0=没有 Masking-key；1=有 Masking-key；
	payloadLen8  uint8   //Payload len，7 bit
	payloadLen16 uint16  //Extended payload length，16 bit，if payload len==126
	payloadLen64 uint64  //Extended payload length，64 bit，if payload len==127
	maskingKey   [4]byte //Masking-key，4 byte

	headerLen  int //请求头长度
	payloadLen int //请求体长度

	MsgLen int    //消息长度
	Msg    []byte //消息体

	Payload string //请求体
}

func NewMaskTestMsg() *Msg {
	req := &Msg{
		Fin:    fin1,
		Opcode: opcodeText,
		Mask:   musk1,
	}

	timestamp := time.Now().Second()
	req.maskingKey[0] = byte(timestamp)
	req.maskingKey[1] = byte(timestamp - 1)
	req.maskingKey[2] = byte(timestamp - 2)
	req.maskingKey[3] = byte(timestamp - 3)

	return req
}

func NewUnmaskTextMsg() *Msg {
	return &Msg{
		Fin:    fin1,
		Opcode: opcodeText,
		Mask:   musk0,
	}
}

func NewPingMsg() *Msg {
	return &Msg{
		Fin:    fin1,
		Opcode: opcodePing,
		Mask:   musk0,
	}
}

func NewPongMsg() *Msg {
	return &Msg{
		Fin:    fin1,
		Opcode: opcodePong,
		Mask:   musk0,
	}
}

func (t *Msg) encode() ([]byte, error) {
	t.payloadLen = len(t.Payload)
	var msg []byte
	if t.Mask == musk1 {
		i, j := 0, 0

		if t.payloadLen <= 125 {
			msg = make([]byte, 2+4+t.payloadLen)
			msg[0] = t.Fin | t.Opcode
			msg[1] = t.Mask | (0b01111111 & byte(t.payloadLen))
			msg[2] = t.maskingKey[0]
			msg[3] = t.maskingKey[1]
			msg[4] = t.maskingKey[2]
			msg[5] = t.maskingKey[3]
			i = 6
		} else if t.payloadLen <= 65535 {
			msg = make([]byte, 2+2+4+t.payloadLen)
			msg[0] = t.Fin | t.Opcode
			msg[1] = t.Mask | 126
			msg[2] = uint8(t.payloadLen16 >> 8)
			msg[3] = uint8(t.payloadLen16)
			msg[4] = t.maskingKey[0]
			msg[5] = t.maskingKey[1]
			msg[6] = t.maskingKey[2]
			msg[7] = t.maskingKey[3]
			i = 8
		} else {
			msg = make([]byte, 2+8+4+t.payloadLen)
			msg[0] = t.Fin | t.Opcode
			msg[1] = t.Mask | 127
			msg[2] = uint8(t.payloadLen64 >> 56)
			msg[3] = uint8(t.payloadLen64 >> 48)
			msg[4] = uint8(t.payloadLen64 >> 40)
			msg[5] = uint8(t.payloadLen64 >> 32)
			msg[6] = uint8(t.payloadLen64 >> 24)
			msg[7] = uint8(t.payloadLen64 >> 16)
			msg[8] = uint8(t.payloadLen64 >> 8)
			msg[9] = uint8(t.payloadLen64)
			msg[10] = t.maskingKey[0]
			msg[11] = t.maskingKey[1]
			msg[12] = t.maskingKey[2]
			msg[13] = t.maskingKey[3]
			i = 14
		}
		for j < t.payloadLen {
			msg[i] = t.Payload[j] ^ t.maskingKey[j%4]
			i++
			j++
		}
	} else {
		i, j := 0, 0
		if t.payloadLen <= 125 {
			msg = make([]byte, 2+t.payloadLen)
			msg[0] = t.Fin | t.Opcode
			msg[1] = t.Mask | (0b01111111 & byte(t.payloadLen))
			i = 2
		} else if t.payloadLen <= 65535 {
			msg = make([]byte, 2+2+t.payloadLen)
			msg[0] = t.Fin | t.Opcode
			msg[1] = t.Mask | 126
			msg[2] = uint8(t.payloadLen16 >> 8)
			msg[3] = uint8(t.payloadLen16)
			i = 4
		} else {
			msg = make([]byte, 2+8+t.payloadLen)
			msg[0] = t.Fin | t.Opcode
			msg[1] = t.Mask | 127
			msg[2] = uint8(t.payloadLen64 >> 56)
			msg[3] = uint8(t.payloadLen64 >> 48)
			msg[4] = uint8(t.payloadLen64 >> 40)
			msg[5] = uint8(t.payloadLen64 >> 32)
			msg[6] = uint8(t.payloadLen64 >> 24)
			msg[7] = uint8(t.payloadLen64 >> 16)
			msg[8] = uint8(t.payloadLen64 >> 8)
			msg[9] = uint8(t.payloadLen64)
			i = 10
		}
		for j < t.payloadLen {
			msg[i] = t.Payload[j]
			i++
			j++
		}
	}
	return msg, nil
}

func (t *Msg) decode(buffer []byte, bufferLen int) error {

	if bufferLen < 2 {
		return ErrParseFailed //至少 2 个字节才能解析
	}

	t.Fin = buffer[0] & 0b10000000 //取 FIN，第 1 个字节的第 1 位

	t.Opcode = buffer[0] & 0b00001111 //取 OPCODE，第 1 个字节的后 4 位

	t.headerLen = 2 //请求头至少 2 个字节

	t.Mask = buffer[1] & 0b10000000 //取 MASK，第 2 个字节的第 1 位
	if t.Mask == musk1 {
		t.headerLen += 4 //有 Masking-key，请求头再加 4 个字节
	}

	t.payloadLen8 = buffer[1] & 0b01111111 //取 Payload len，第 2 个字节的后 7 位
	if t.payloadLen8 == 126 {
		t.headerLen += 2 //如果 Payload len==126，请求头再加 2 个字节
	} else if t.payloadLen8 == 127 {
		t.headerLen += 8 //如果 Payload len==127，请求头再加 8 个字节
	}

	if bufferLen < t.headerLen {
		return ErrParseFailed //请求头长度不够
	}

	if t.payloadLen8 == 126 {
		//取 Extended payload length，第 3、4 字节
		t.payloadLen16 = 0
		t.payloadLen16 |= uint16(buffer[2]) << 8
		t.payloadLen16 |= uint16(buffer[3])
		t.payloadLen = int(t.payloadLen16)
	} else if t.payloadLen8 == 127 {
		//取 Extended payload length，第 3、4、5、6、7、8、9、10 字节
		t.payloadLen64 = 0
		t.payloadLen64 |= uint64(buffer[2]) << 56
		t.payloadLen64 |= uint64(buffer[3]) << 48
		t.payloadLen64 |= uint64(buffer[4]) << 40
		t.payloadLen64 |= uint64(buffer[5]) << 32
		t.payloadLen64 |= uint64(buffer[6]) << 24
		t.payloadLen64 |= uint64(buffer[7]) << 16
		t.payloadLen64 |= uint64(buffer[8]) << 8
		t.payloadLen64 |= uint64(buffer[9])
		t.payloadLen = int(t.payloadLen64)
	} else {
		t.payloadLen = int(t.payloadLen8)
	}
	t.MsgLen = t.headerLen + t.payloadLen

	if bufferLen < t.MsgLen {
		return ErrParseFailed //消息长度不够
	}

	t.Msg = buffer[:t.MsgLen]

	if t.Mask == musk1 {
		//取 Masking-key，请求头倒数第 4、3、2、1 字节
		t.maskingKey[0] = buffer[t.headerLen-4]
		t.maskingKey[1] = buffer[t.headerLen-3]
		t.maskingKey[2] = buffer[t.headerLen-2]
		t.maskingKey[3] = buffer[t.headerLen-1]

		// 用 Masking-key 解析 Payload Data
		msgUnMask := make([]byte, t.MsgLen)
		copy(msgUnMask, t.Msg)
		i, j := 0, t.headerLen
		for j < t.MsgLen {
			msgUnMask[j] = msgUnMask[j] ^ t.maskingKey[i%4]
			i++
			j++
		}
		t.Payload = string(msgUnMask[t.headerLen:])
	} else {
		t.Payload = string(t.Msg[t.headerLen:])
	}

	return nil
}

func (t *Msg) parseJson() (map[string]any, error) {
	ret := make(map[string]any)
	err := json.Unmarshal([]byte(t.Payload), &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
