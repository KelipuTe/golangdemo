package stream

import (
	"encoding/json"
)

type Request struct {
	Addr string //请求IP和端口
	MsgBody
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
	t.Msg = buffer[0:msgLen]
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
