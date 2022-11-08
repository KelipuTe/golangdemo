package metadata

import (
	"errors"
	"fmt"
)

func F8NewErrInputOnlyStructPointer() error {
	return errors.New("orm: 只支持一级结构体指针作为输入\r\n")
}

func F8NewErrInvalidTagContent(tag string) error {
	return fmt.Errorf("orm: 标签 [%s] 格式错误\r\n", tag)
}
