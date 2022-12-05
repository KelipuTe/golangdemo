package mapkn

import (
	"fmt"
	"testing"
)

func TestMapStruct(p7s6t *testing.T) {
	t4m3 := f8MapStruct()
	fmt.Printf("map,%p\r\n", t4m3)
}

func TestMapStructPointer(p7s6t *testing.T) {
	t4m3 := f8MapStructPointer()
	fmt.Printf("map,%p\r\n", t4m3)
}
