package rpc

import (
	"context"
	"demo-golang/rpc/serialize"
	"reflect"
)

// ServiceReflect 本地服务的反射
type ServiceReflect struct {
	service ServiceI9
	reflect reflect.Value
}

func (t *ServiceReflect) handleRPC(ctx context.Context, funcName string, funcIn []byte, serialize serialize.SerializeI9) ([]byte, error) {
	// 通过方法名，从结构体的反射中找到方法
	s6MethodValue := t.reflect.MethodByName(funcName)
	// 拿到方法的第二个入参的类型，第一个是 context
	inputType := s6MethodValue.Type().In(1)
	// 构造方法的第二个入参参
	inputValue := reflect.New(inputType.Elem())
	input := inputValue.Interface()
	// 把传过来的编码后的入参解码，然后放到构造的入参上去
	err := serialize.F8Decode(funcIn, input)
	if err != nil {
		return nil, err
	}
	output := s6MethodValue.Call([]reflect.Value{reflect.ValueOf(ctx), inputValue})
	// 判断有没有 error
	if len(output) > 1 && !output[1].IsZero() {
		return nil, output[1].Interface().(error)
	}
	return serialize.F8Encode(output[0].Interface())
}
