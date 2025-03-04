package interfacehkn

type interface1 interface {
	func1()
}

type struct1 struct {
}

func (t struct1) func1() {
}

//这两种写法可以检查 struct1 是否实现了 interface1

var _ interface1 = &struct1{}

var _ interface1 = (*struct1)(nil)
