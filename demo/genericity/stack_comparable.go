package genericity

// StackComparable 泛型栈
type StackComparable[T comparable] struct {
	s5value []T
}

// Init 初始化
func (p7this *StackComparable[T]) Init() {
	if nil == p7this.s5value {
		p7this.s5value = make([]T, 0, 2)
	}
}

// Push 入栈
func (p7this *StackComparable[T]) Push(value T) {
	if nil == p7this.s5value {
		return
	}

	p7this.s5value = append(p7this.s5value, value)
}

// Contains 是否存在元素
func (p7this *StackComparable[T]) Contains(value T) bool {
	for _, v := range p7this.s5value {
		// 声明为 StackComparable[T comparable] 时，这里就可以直接比较
		if v == value {
			return true
		}
	}
	return false
}
