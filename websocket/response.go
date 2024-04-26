package websocket

type Response struct {
	MsgBody
}

func NewResponse() *Response {
	return &Response{
		MsgBody{
			Fin:    fin1,
			Opcode: opcodeText,
			Mask:   musk0,
		},
	}
}

func (t *Response) encode() ([]byte, error) {
	return encode(&t.MsgBody)
}

func (t *Response) decode(buffer []byte, bufferLen int) error {
	msgBody, err := decode(buffer, bufferLen)
	if err != nil {
		return err
	}

	t.MsgBody = *msgBody

	return nil
}
