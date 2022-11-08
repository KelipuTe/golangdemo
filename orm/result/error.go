package result

import (
	"errors"
	"fmt"
)

var ErrNoRows = errors.New("orm: 未找到数据")
var ErrTooManyReturnedColumns = errors.New("orm: 返回的列过多")

func F8NewErrUnknownColumn(col string) error {
	return fmt.Errorf("orm: 未知列 %s", col)
}
