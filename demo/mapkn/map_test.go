package interfacehkn

import (
	"log"
	"testing"
)

func TestFunction(p7t *testing.T) {
	log.Println(4 << (^uintptr(0) >> 63))

	m := map[int]int{
		11: 11,
		22: 22,
		33: 33,
		44: 44,
		55: 55,
	}

	for {
		for k, v := range m {
			log.Printf("k=%d;v=%d;", k, v)
		}
	}
}

func TestFunction2(p7t *testing.T) {
	m := map[int]int{}

	m[11] = 11
	m[22] = 22
	m[33] = 33
	m[44] = 44
	m[55] = 55
	m[66] = 66
	m[77] = 77
	m[88] = 88
	m[99] = 99

	for {
		log.Printf("map;")
		for k, v := range m {
			log.Printf("k=%d;v=%d;", k, v)
		}
	}
}
