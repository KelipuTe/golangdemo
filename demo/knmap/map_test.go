package knmap

import (
	"fmt"
	"testing"
)

func TestMapStruct(t *testing.T) {
	t4map := MapStruct()
	fmt.Printf("map,%p\r\n", t4map)
}

func TestMapStructPointer(t *testing.T) {
	t4map := MapStructPointer()
	fmt.Printf("map,%p\r\n", t4map)
}
