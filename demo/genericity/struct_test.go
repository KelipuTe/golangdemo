package genericity

// 泛型类型和泛型方法示例

import (
	"log"
	"testing"
)

type AnyStruct[T any] struct {
	value T
}

func TestAnyStruct(t *testing.T) {
	intStruct := AnyStruct[int]{1}
	log.Println(intStruct.value)

	strStruct := AnyStruct[string]{"1"}
	log.Println(strStruct.value)
}

func (t *AnyStruct[T]) SetAny(value T) {
	t.value = value
}

func TestAnyStructFunc(t *testing.T) {
	intStruct := AnyStruct[int]{1}
	log.Println(intStruct.value)
	intStruct.SetAny(2)
	log.Println(intStruct.value)

	strStruct := AnyStruct[string]{"1"}
	log.Println(strStruct.value)
	strStruct.SetAny("2")
	log.Println(strStruct.value)
}
