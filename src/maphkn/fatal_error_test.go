package maphkn

import (
	"log"
	"testing"
	"time"
)

// 并发写 map 会导致 fatal error: concurrent map writes
func TestCW(t *testing.T) {
	m := map[int]int{1: 1, 2: 2}

	go func() {
		for {
			m[1] = 1
		}
	}()

	go func() {
		for {
			m[2] = 2
		}
	}()

	for {
		log.Println("sleep")
		time.Sleep(time.Second)
	}
}

// 并发读写 map 会导致 fatal error: concurrent map read and map write
func TestCRW(t *testing.T) {
	m := map[int]int{1: 1, 2: 2}

	go func() {
		for {
			m[1] = 1
		}
	}()

	go func() {
		for {
			a := m[2]
			log.Println(a)
		}
	}()

	for {
		log.Println("sleep")
		time.Sleep(time.Second)
	}
}

// 并发读 map 不会导致 fatal error
func TestCR(t *testing.T) {
	m := map[int]int{1: 1, 2: 2}

	go func() {
		for {
			a := m[1]
			log.Println(a)
		}
	}()

	go func() {
		for {
			b := m[2]
			log.Println(b)
		}
	}()

	for {
		log.Println("sleep")
		time.Sleep(time.Second)
	}
}
