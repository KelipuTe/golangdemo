package interfacehkn

type checkInterface interface {
	checkFunc()
}

type checkStruct struct {
}

func (t checkStruct) checkFunc() {
}

//这两种写法可以检查 struct1 是否实现了 interface1

var _ checkInterface = &checkStruct{}

var _ checkInterface = (*checkStruct)(nil)
