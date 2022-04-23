package websocket

// Response 响应
type Response struct {
}

func NewResponse() *Response {
  return &Response{}
}

// MakeMsg 构造响应报文
func (p1this *Response) MakeResponse(body string) []byte {
  sli1body := []byte(body)
  bodyLen := len(sli1body)

  var msg []byte
  if bodyLen <= 125 {
    msg = make([]byte, 2)
    msg[0] = 0b10000000 | opcodeText
    msg[1] = 0b01111111 & uint8(bodyLen)
    msg = append(msg, sli1body...)
  } else if bodyLen <= 65535 {
    msg = make([]byte, 4)
    msg[0] = 0b10000000 | opcodeText
    msg[1] = 126
    msg[2] = uint8(bodyLen >> 8)
    msg[3] = uint8(bodyLen >> 0)
    msg = append(msg, sli1body...)
  } else {
    msg = make([]byte, 10)
    msg[0] = 0b10000000 | opcodeText
    msg[1] = 127
    msg[2] = uint8(bodyLen >> 56)
    msg[3] = uint8(bodyLen >> 48)
    msg[4] = uint8(bodyLen >> 40)
    msg[5] = uint8(bodyLen >> 32)
    msg[6] = uint8(bodyLen >> 24)
    msg[7] = uint8(bodyLen >> 16)
    msg[8] = uint8(bodyLen >> 8)
    msg[9] = uint8(bodyLen >> 0)
    msg = append(msg, sli1body...)
  }

  return msg
}
