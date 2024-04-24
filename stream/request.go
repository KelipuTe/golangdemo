package stream

import (
	"encoding/binary"
	"encoding/json"
	"errors"
)

var (
	ErrParseFailed = errors.New("解析失败")
)

type Request struct {
	MsgLen int    //消息长度
	Msg    string //消息体

	Addr string //请求IP和端口
	Body string //请求体
}

func NewRequest() *Request {
	return &Request{}
}

func (t *Request) encode() ([]byte, error) {
	return encode(t.Body)
}

func (t *Request) decode(buffer []byte, bufferLen int) error {
	msgLen, err := decode(buffer, bufferLen)
	if err != nil {
		return err
	}

	t.MsgLen = msgLen
	t.Msg = string(buffer[0:msgLen])
	t.Body = string(buffer[4:msgLen])

	return nil
}

func (t *Request) parseJson() (map[string]any, error) {
	ret := make(map[string]any)
	err := json.Unmarshal([]byte(t.Body), &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func encode(msgStr string) ([]byte, error) {
	bodyLen := len(msgStr)
	if bodyLen <= 0 {
		return nil, ErrParseFailed
	}
	msgSlice := make([]byte, 0, 4+bodyLen)

	//把消息长度转换成大端字节序格式
	msgLenBig := make([]byte, 4)
	binary.BigEndian.PutUint32(msgLenBig, uint32(4+bodyLen))

	msgSlice = append(msgSlice, msgLenBig...)      //最前面的4个字节是消息长度
	msgSlice = append(msgSlice, []byte(msgStr)...) //从第5个字节开始，后面都是请求体

	return msgSlice, nil
}

func decode(buffer []byte, bufferLen int) (int, error) {
	if bufferLen <= 4 {
		return 0, ErrParseFailed
	}

	//把大端字节序格式的消息长度转换回来
	msgLenBig := buffer[0:4]
	msgLen := int(binary.BigEndian.Uint32(msgLenBig))
	if bufferLen < msgLen {
		return 0, ErrParseFailed
	}

	return msgLen, nil
}
