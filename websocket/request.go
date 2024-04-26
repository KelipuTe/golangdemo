package websocket

import "encoding/json"

type Request struct {
	Addr string //请求IP和端口
	MsgBody
}

func NewRequest() *Request {
	return &Request{
		MsgBody: MsgBody{
			Fin:    fin1,
			Opcode: opcodeText,
			Mask:   musk1,
		},
	}
}

func (t *Request) encode() ([]byte, error) {
	return encode(&t.MsgBody)
}

func (t *Request) decode(buffer []byte, bufferLen int) error {
	msgBody, err := decode(buffer, bufferLen)
	if err != nil {
		return err
	}

	t.MsgBody = *msgBody

	return nil
}

func (t *Request) parseJson() (map[string]any, error) {
	ret := make(map[string]any)
	err := json.Unmarshal([]byte(t.Payload), &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (t *Request) makeHandshake(data map[string]any) error {
	ret, err := json.Marshal(data)
	if err != nil {
		return err
	}

	t.Payload = string(ret)

	return nil
}
