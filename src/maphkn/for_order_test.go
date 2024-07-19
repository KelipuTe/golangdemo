package maphkn

import (
	"log"
	"testing"
)

// 使用 for range 遍历 map 的时候，输出元素的顺序，每次都是不一样的

func TestFor5Item(p7t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}

	for {
		log.Printf("map;")
		for k, v := range m {
			log.Printf("k=%d;v=%d;", k, v)
		}
	}
}

func TestFor9Item(p7t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}

	for {
		log.Printf("map;")
		for k, v := range m {
			log.Printf("k=%d;v=%d;", k, v)
		}
	}
}
