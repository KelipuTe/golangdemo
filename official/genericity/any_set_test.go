package genericity

//泛型集合示例
//集合需要保证元素唯一，这里就涉及到两个泛型变量的比较问题。
//两个泛型变量是不能直接进行比较的。需要先通过断言或者反射拿到类型信息。
//拿到类型信息之后就可以进行后续的比较操作了，但是可以预见会很繁琐。

//对泛型变量直接使用类型断言是不可以的。
//直接断言和 switch type 断言两种方法都会报错。
//需要先使用空接口 interface{} 做泛型约束。

//可以用 comparable 处理比较问题，但是这种的就用不了 any 了。

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

type AnySet[T any] struct {
	storage []T
}

func (t *AnySet[T]) Init() {
	if t.storage != nil {
		return
	}
	t.storage = make([]T, 0, 8)
}

func (t *AnySet[T]) Add(value T) {
	if t.storage == nil {
		return
	}
	t.storage = append(t.storage, value)
}

func (t *AnySet[T]) Contains(value T) bool {
	b1 := t.typeContain(value)
	b2 := t.reflectContain(value)
	return b1 && b2
}

// 通过断言比较（这里只处理了 int 和 string）
func (t *AnySet[T]) typeContain(value T) bool {
	for _, v := range t.storage {
		vType := fmt.Sprintf("%T", v)
		valueType := fmt.Sprintf("%T", value)
		if vType != valueType {
			continue
		}
		//上面已经判断了类型相同，这里 value 就可以直接断言了。
		var vi interface{} = v
		var valuei interface{} = value
		switch vi2 := vi.(type) {
		case int:
			if vi2 == valuei.(int) {
				return true
			}
		case string:
			if vi2 == valuei.(string) {
				return true
			}
		}
	}

	return false
}

// 通过反射比较（这里只处理了 int 和 string）
func (t *AnySet[T]) reflectContain(value T) bool {
	for _, v := range t.storage {
		vr := reflect.ValueOf(v)
		valuer := reflect.ValueOf(value)

		if vr.Type() != valuer.Type() {
			continue //类型不同
		}

		switch vr.Kind() {
		case reflect.Int:
			if vr.Int() == valuer.Int() {
				return true
			}
		case reflect.String:
			if vr.String() == valuer.String() {
				return true
			}
		}
	}

	return false
}

func TestAnySet(t *testing.T) {
	var intSet AnySet[int]
	intSet.Init()
	intSet.Add(10)
	intSet.Add(20)
	log.Println(intSet)

	var strSet AnySet[string]
	strSet.Init()
	strSet.Add("aa")
	strSet.Add("bb")
	log.Println(strSet)
}

func TestAnySetContains(t *testing.T) {
	var anySet AnySet[any]
	anySet.Init()
	anySet.Add(10)
	anySet.Add("aa")
	log.Println(anySet.Contains(10))
	log.Println(anySet.Contains("bb"))
}

type comparableSet[T comparable] struct {
	storage []T
}

func (t *comparableSet[T]) Init() {
	if t.storage != nil {
		return
	}
	t.storage = make([]T, 0, 8)
}

func (t *comparableSet[T]) Add(value T) {
	if t.storage == nil {
		return
	}
	t.storage = append(t.storage, value)
}

func (t *comparableSet[T]) Contains(value T) bool {
	for _, v := range t.storage {
		if value == v {
			return true
		}
	}
	return false
}

func TestComparableSet(t *testing.T) {
	var intSet comparableSet[int]
	intSet.Init()
	intSet.Add(10)
	log.Println(intSet.Contains(10))
	log.Println(intSet.Contains(20))

	var strSet comparableSet[string]
	strSet.Init()
	strSet.Add("aa")
	log.Println(strSet.Contains("aa"))
	log.Println(strSet.Contains("bb"))
}
