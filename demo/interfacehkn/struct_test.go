package interfacehkn

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStruct(p7s6t *testing.T) {
	var ok bool = false
	var s6dog S6Dog = S6Dog{"dog"}

	// s6dog 是没有 f8SetName 方法的，只有 &s6dog 才有 f8SetName 方法。
	// 下面的四个输出依次是：false，true，true，true。
	_, ok = interface{}(s6dog).(i9Pet)
	fmt.Printf("dog is i9Pet: %v\n", ok)
	_, ok = interface{}(s6dog).(i9PetV2)
	fmt.Printf("dog is i9PetV2: %v\n", ok)
	_, ok = interface{}(&s6dog).(i9Pet)
	fmt.Printf("*dog is i9Pet: %v\n", ok)
	_, ok = interface{}(&s6dog).(i9PetV2)
	fmt.Printf("*dog is i9PetV2: %v\n", ok)

	// 变量 i9Pet 的静态类型是 i9Pet（接口），动态类型是 *S6Dog（接口实现）。
	var i9Pet i9Pet = &s6dog
	fmt.Println("TypeOf i9Pet is", reflect.TypeOf(i9Pet))
	fmt.Println("pet is", i9Pet.f8GetCategory(), ", name is", i9Pet.f8GetName())

	// 变量 i9PetV2 的静态类型是 i9PetV2（接口），动态类型是 S6Dog（接口实现）。
	// 因为接口 i9PetV2 没有指针方法，所以变量 i9PetV2 的赋值就变成了 s6dog 的副本。
	var i9PetV2 i9PetV2 = s6dog
	fmt.Println("TypeOf i9PetV2 is", reflect.TypeOf(i9PetV2))
	fmt.Println("pet is", i9PetV2.f8GetCategory(), ", name is", i9PetV2.f8GetName())

	// 这里修改了 s6dog 的值，但是副本不受影响，所以 i9Pet 的值变，i9PetV2 的值不变。
	s6dog.f8SetName("dog2")
	fmt.Println("pet is", i9Pet.f8GetCategory(), ", name is", i9Pet.f8GetName())
	fmt.Println("pet is", i9PetV2.f8GetCategory(), ", name is", i9PetV2.f8GetName())
}
