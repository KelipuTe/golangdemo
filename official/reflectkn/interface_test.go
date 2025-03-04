package reflectkn

import (
	"log"
	"reflect"
	"testing"
)

//判断某个结构体是否实现了某个接口

type CaseI9 interface {
	CaseFunc()
}

type CaseStruct struct{}

func (t CaseStruct) CaseFunc() {}

func TestIsImplement(t *testing.T) {
	s6 := CaseStruct{}
	s6rt := reflect.TypeOf(s6)
	log.Println("s6rt kind", s6rt.Kind())

	//这里nil只能转接口指针
	//因为结构体有值，不可能把没值的nil转成结构体
	i9 := (*CaseI9)(nil)
	i9rt := reflect.TypeOf(i9)
	log.Println("i9rt kind", i9rt.Kind())

	//因为前面是指针，所以这里要取指针指向的结构体
	i9rte := i9rt.Elem()

	ok := s6rt.Implements(i9rte)
	log.Println("s6rt implements i9rt", ok)
}
