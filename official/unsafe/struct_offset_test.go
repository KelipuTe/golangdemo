package unsafe

import (
	"testing"
)

func TestPrintStructFieldOffset(t *testing.T) {
	PrintStructFieldOffset(UserV2{})
	// type:  UserV2, size:64
	// field:    Name,type:  string, size:16, offset: 0
	// field:     Age,type:   int32, size: 4, offset:16
	// field:   Alias,type:        , size:24, offset:24
	// field: Address,type:  string, size:16, offset:48

	PrintStructFieldOffset(UserV4{})
	// type:  UserV4, size:64
	// field:    Name,type:  string, size:16, offset: 0
	// field:     Age,type:   int32, size: 4, offset:16
	// field:   AgeV2,type:   int32, size: 4, offset:20
	// field:   Alias,type:        , size:24, offset:24
	// field: Address,type:  string, size:16, offset:48

	// 可以注意到 UserV2 和 UserV4 虽然差了一个 AgeV2 int32 的字段，但是大小是一样的。
	// 这是因为，在分配内存的时候，会按照字长对齐。32 位机器是 4 字节，64 位机器是 8 字节。
	// 所以在 64 位的机器上，UserV2 给 Age 字段分配内存的时候，分配的是 8 个字节中的 4 个字节。
	// 下一个字段 Alias 的长度又超过 4 个字节，所以需要相当于另起一行。但是 UserV4 的 AgeV2 正好可以填进去。

	PrintStructFieldOffset(UserV6{})
	PrintStructFieldOffset(UserV8{})
}
