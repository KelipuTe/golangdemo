package protocol

import (
  "demo_golang/net_service/tool"
  "encoding/binary"
  "errors"
)

type Stream struct {
  ParseStatus int    // 解析状态
  BodyLength  int    // 数据长度
  BodyStr     string // 数据
}

func (p1this *Stream) DataLength(data []byte) (dataLength int, err error) {
  totalLen := len(data)
  if 0 == totalLen {
    err = errors.New("STREAM_STATUS_NO_DATA")
    return
  }
  if totalLen < 4 {
    err = errors.New("STREAM_STATUS_NOT_FINISH")
    return
  }
  tool.DebugPrintln("Stream.DataLength()", data[0:4])

  dataLength = int(binary.BigEndian.Uint32(data[0:4]))
  if totalLen < 4+dataLength {
    err = errors.New("STREAM_STATUS_NOT_FINISH")
    return
  }
  p1this.BodyLength = dataLength
  dataLength = 4 + dataLength

  return
}

func (p1this *Stream) DataDecode(data []byte) (decodeData []byte, err error) {
  decodeData = data[4:]
  p1this.BodyStr = string(decodeData)
  return
}

func (p1this *Stream) DataEncode(data []byte) (encodeData []byte, err error) {
  dataLength := len(data)
  if 0 == dataLength {
    err = errors.New("STREAM_STATUS_NO_DATA")
    return
  }
  tempData := make([]byte, 4+dataLength)
  binary.BigEndian.PutUint32(tempData[0:4], uint32(dataLength))
  tempData = append(tempData[0:4], data...)
  encodeData = tempData

  return
}
