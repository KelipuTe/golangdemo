package interfacehkn

import (
	"reflect"
	"testing"
)

// 任何数据类型，只要它的方法集合中包含了一个接口定义的全部方法，那么它就是这个接口的实现。
// 值类型的方法集合中仅会包含所有值方法，指针类型的方法集合中包会包含所有指针方法和所有值方法。

type Pet interface {
	SetName(name string)
	GetName() string
	GetCategory() string
}

type PetV2 interface {
	GetName() string
	GetCategory() string
}

type Dog struct {
	name string
}

func (t *Dog) SetName(name string) {
	t.name = name
}

func (t Dog) GetName() string {
	return t.name
}

func (t Dog) GetCategory() string {
	return "dog"
}

func TestStruct(t *testing.T) {
	var ok bool = false
	var dog Dog = Dog{"dog"}

	// dog 是没有 SetName 方法的，只有 &dog 才有 SetName 方法。
	// 下面的四个输出依次是：false，true，true，true。
	_, ok = interface{}(dog).(Pet)
	t.Log("dog is Pet:", ok)
	_, ok = interface{}(dog).(PetV2)
	t.Log("dog is PetV2:", ok)
	_, ok = interface{}(&dog).(Pet)
	t.Log("*dog is Pet:", ok)
	_, ok = interface{}(&dog).(PetV2)
	t.Log("*dog is PetV2:", ok)

	// 变量 Pet 的静态类型是 Pet（接口），动态类型是 *Dog（接口实现）。
	var pet Pet = &dog
	t.Log("TypeOf Pet is", reflect.TypeOf(pet))
	t.Log("pet is", pet.GetCategory(), ", name is", pet.GetName())

	// 变量 PetV2 的静态类型是 PetV2（接口），动态类型是 Dog（接口实现）。
	// 因为接口 PetV2 没有指针方法，所以变量 PetV2 的赋值就变成了 dog 的副本。
	var petV2 PetV2 = dog
	t.Log("TypeOf PetV2 is", reflect.TypeOf(petV2))
	t.Log("pet is", petV2.GetCategory(), ", name is", petV2.GetName())

	// 这里修改了 dog 的值，但是副本不受影响，所以 Pet 的值变，PetV2 的值不变。
	dog.SetName("dog2")
	t.Log("pet is", pet.GetCategory(), ", name is", pet.GetName())
	t.Log("pet is", petV2.GetCategory(), ", name is", petV2.GetName())
}
