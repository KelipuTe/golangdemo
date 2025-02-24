package serialize

// SerializeI9 序列化接口
type SerializeI9 interface {
	F8GetCode() uint8
	F8Encode(anyInput any) ([]byte, error)
	F8Decode(s5input []byte, anyOutput any) error
}
