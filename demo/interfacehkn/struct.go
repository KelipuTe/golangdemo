package interfacehkn

// 任何数据类型，只要它的方法集合中包含了一个接口定义的全部方法，那么它就是这个接口的实现。
// 值类型的方法集合中仅会包含所有值方法，指针类型的方法集合中包会包含所有指针方法和所有值方法。

type i9Pet interface {
	f8SetName(name string)
	f8GetName() string
	f8GetCategory() string
}

type i9PetV2 interface {
	f8GetName() string
	f8GetCategory() string
}

type S6Dog struct {
	name string
}

func (p7this *S6Dog) f8SetName(name string) {
	p7this.name = name
}

func (this S6Dog) f8GetName() string {
	return this.name
}

func (this S6Dog) f8GetCategory() string {
	return "dog"
}
