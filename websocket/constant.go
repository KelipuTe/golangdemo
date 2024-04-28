package websocket

const (
	connPoolNumMax = 1024 //最大tcp连接数
)

const (
	readBufferMaxLen = 1048576 // 1048576 == 2^20 == 1MB。
)

const (
	fin0 = 0b00000000 //Fin=0
	fin1 = 0b10000000 //Fin=1
)

const (
	musk0 = 0b00000000 //musk=0
	musk1 = 0b10000000 //musk=1
)

const (
	opcodeText   uint8 = 0x01 //文本帧
	opcodeBinary uint8 = 0x02 //二进制帧
	opcodeClose  uint8 = 0x08 //连接断开
	opcodePing   uint8 = 0x09 //ping
	opcodePong   uint8 = 0x0A //pong
)
