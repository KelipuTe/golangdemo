package orm_result

import (
	"errors"
	"fmt"
)

func NewErrUnknownColumn(col string) error {
	return fmt.Errorf("orm: 未知列 %s", col)
}

var ErrNoRows = errors.New("orm: 未找到数据")
