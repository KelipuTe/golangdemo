package unsafe

import (
	"reflect"
	"unsafe"
)

// StructIntFieldAccessor 存取 int 类型的字段
type StructIntFieldAccessor interface {
	GetIntField(field string) (int, error)
	SetIntField(field string, val int) error
}

// StructAnyFieldAccessor 存取任意类型的字段
type StructAnyFieldAccessor interface {
	GetAnyField(field string) (any, error)
	SetAnyField(field string, val any) error
}

// StructFieldMeta 结构体字段的数据
type StructFieldMeta struct {
	// 字段的类型
	fieldType reflect.Type
	// 字段起始地址相对于结构体起始地址的偏移量
	// 结构体起始地址 + 字段起始地址相对于结构体起始地址的偏移量，就可以得到字段的起始地址
	// 在考虑组合或者复杂类型字段的时候，它的含义可以衍生为，相对于最外层的结构体起始地址的偏移量
	offset uintptr
}

// StructUnsafeAccessor 通过地址直接访问结构体的字段
type StructUnsafeAccessor struct {
	// 结构体起始地址
	p7entity unsafe.Pointer
	mapsfm   map[string]StructFieldMeta
}

func NewStructUnsafeAccessor(entity interface{}) (*StructUnsafeAccessor, error) {
	if nil == entity {
		return nil, ErrMustStruct
	}

	t4type := reflect.TypeOf(entity)

	for reflect.Pointer == t4type.Kind() {
		t4type = t4type.Elem()
	}

	if t4type.Kind() != reflect.Struct {
		return nil, ErrMustStruct
	}

	t4num := t4type.NumField()
	t4mapsfm := make(map[string]StructFieldMeta, t4num)
	for i := 0; i < t4num; i++ {
		t4f := t4type.Field(i)
		// 拿到字段的类型和字段起始地址相对于结构体起始地址的偏移量
		t4mapsfm[t4f.Name] = StructFieldMeta{fieldType: t4f.Type, offset: t4f.Offset}
	}

	t4val := reflect.ValueOf(entity)
	p7sua := &StructUnsafeAccessor{p7entity: t4val.UnsafePointer(), mapsfm: t4mapsfm}

	return p7sua, nil
}

func (p7this *StructUnsafeAccessor) GetIntField(field string) (int, error) {
	meta, ok := p7this.mapsfm[field]
	if !ok {
		return 0, ErrFieldNotFound
	}

	// uintptr(p7this.p7entity)，结构体起始地址
	// uintptr(p7this.p7entity) + meta.offset，字段的起始地址
	// unsafe.Pointer(uintptr(p7this.p7entity) + meta.offset)，得到一个指针
	// 这里假设知道是个 int 类型的字段，所以可以把这个指针转换成 int 指针
	// 然后通过 * 访问指针，就可以得到指针指向的数据
	data := *(*int)(unsafe.Pointer(uintptr(p7this.p7entity) + meta.offset))
	return data, nil
}

func (p7this *StructUnsafeAccessor) SetIntField(field string, val int) error {
	meta, ok := p7this.mapsfm[field]
	if !ok {
		return ErrFieldNotFound
	}

	// 取值操作反过来就是赋值
	*(*int)(unsafe.Pointer(uintptr(p7this.p7entity) + meta.offset)) = val
	return nil
}

func (p7this *StructUnsafeAccessor) GetAnyField(field string) (any, error) {
	meta, ok := p7this.mapsfm[field]
	if !ok {
		return 0, ErrFieldNotFound
	}

	// 指针指向的内存地址上已经有数据了，但是没有类型，无法解析成对应的数据。
	// 这个操作是给指针指向的内存地址，挂一个类型，相当于把这个指针转换成类型指针。
	// 但是，这里不是和 GetIntField 那里一样直接拿到指针，而是和 reflect.ValueOf() 的结果一样的东西。
	t4val := reflect.NewAt(
		meta.fieldType,
		unsafe.Pointer(uintptr(p7this.p7entity)+meta.offset),
	)

	// 注意这里要调 Elem 拿到原对象
	return t4val.Elem().Interface(), nil
}

func (p7this *StructUnsafeAccessor) SetAnyField(field string, val any) error {
	meta, ok := p7this.mapsfm[field]
	if !ok {
		return ErrFieldNotFound
	}

	t4val := reflect.NewAt(
		meta.fieldType,
		unsafe.Pointer(uintptr(p7this.p7entity)+meta.offset),
	)

	// 注意这里要调 Elem 拿到原对象
	if t4val.Elem().CanSet() {
		t4val.Elem().Set(reflect.ValueOf(val))
	}

	return nil
}
