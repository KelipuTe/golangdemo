package stream

type Response struct {
	MsgBody
}

func NewResponse() *Response {
	return &Response{}
}

func (t *Response) encode() ([]byte, error) {
	return encode(t.Body)
}

func (t *Response) decode(buffer []byte, bufferLen int) error {
	msgLen, err := decode(buffer, bufferLen)
	if err != nil {
		return err
	}

	t.MsgLen = msgLen
	t.MsgBody.Msg = buffer[0:msgLen]
	t.Body = string(buffer[4:msgLen])

	return nil
}
