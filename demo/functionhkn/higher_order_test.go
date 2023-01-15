package functionhkn

import (
	"fmt"
	"testing"
)

func TestHigherOrder(p7s6t *testing.T) {
	var f8oe f8Operate

	f8oe = f8GetAddOperate()
	fmt.Println(f8DoCalculate(f8oe, 3, 4))

	f8oe = f8GetMultiplyOperate()
	fmt.Println(f8DoCalculate(f8oe, 3, 4))
}

func TestClosure(p7s6t *testing.T) {
	f8EditOutside()
	f8Remember()
}
