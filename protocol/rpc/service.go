package rpc

import (
	"context"
	"demo-golang/rpc/protocol"
	"reflect"
)

// ServiceI9 实现这个接口，表示本地服务可以被改造成 RPC 服务
// 可以被改造成 RPC 服务的本地服务（结构体）里面，应该都只有方法
type ServiceI9 interface {
	// GetServiceName 获取本地服务对应的 RPC 服务的服务名
	GetServiceName() string
}

// CoverWithRPC 把结构体改造成 RPC 服务
// client 可以发起 RPC 调用的客户端
// service 一个可以被改造成 RPC 服务的结构体
func CoverWithRPC(client ClientI9, service ServiceI9) {
	// 这里肯定是拿到一个接口（结构体指针）
	srv := reflect.ValueOf(service)
	// 通过结构体指针拿到结构体值
	srve := srv.Elem()
	// 通过结构体值拿到结构体类型
	srt := srve.Type()

	// 这里应该全部都是方法
	s6RPCServiceFieldNum := srt.NumField()
	for i := 0; i < s6RPCServiceFieldNum; i++ {
		// 拿到结构体属性
		s6StructField := srt.Field(i)
		// 拿到结构体属性的值
		s6StructFieldValue := srve.Field(i)
		// 拿到结构体属性的类型
		s6StructFieldType := s6StructField.Type
		// 判断一下结构体属性是否可修改
		if !s6StructFieldValue.CanSet() {
			continue
		}
		// 用结构体原来的方法的信息构造一个新的 RPC 调用的方法
		f8NewFunc := func(args []reflect.Value) (results []reflect.Value) {
			// 处理方法的入参，这里只管第二个参数，第一个是 context
			input := args[1].Interface()
			// 处理方法的返回值，这里只管第一个参数，第二个是 error
			output := reflect.New(s6StructFieldType.Out(0).Elem()).Interface()
			// 把方法的入参序列化
			i9serialize := client.GetSerialize()
			inputEncode, err := i9serialize.F8Encode(input)
			if err != nil {
				return []reflect.Value{reflect.ValueOf(output), reflect.ValueOf(err)}
			}
			// 从 context 获取元数据
			m3ExtraData := map[string]string{}
			i9ctx := args[0].Interface()
			if i9ctxValue, ok := i9ctx.(context.Context); ok {
				m3ExtraData["flowId"] = i9ctxValue.Value("flowId").(string)
			}
			// 组装调用的请求数据
			p7s6req := &protocol.Request{
				ServiceName:   service.GetServiceName(),
				FuncName:      s6StructField.Name,
				MetaData:      m3ExtraData,
				SerializeCode: i9serialize.F8GetCode(),
				FuncInput:     inputEncode,
			}
			// 向远端发起调用
			resp, err := client.SendRPC(args[0].Interface().(context.Context), p7s6req)
			if err != nil {
				return []reflect.Value{reflect.ValueOf(output), reflect.ValueOf(err)}
			}
			if resp.Error != nil {
				return []reflect.Value{reflect.ValueOf(output), reflect.ValueOf(resp.Error)}
			}
			// 把返回的数据反序列化
			err = i9serialize.F8Decode(resp.FuncOutput, output)
			if err != nil {
				return []reflect.Value{reflect.ValueOf(output), reflect.ValueOf(err)}
			}
			return []reflect.Value{reflect.ValueOf(output), reflect.Zero(reflect.TypeOf(new(error)).Elem())}
		}
		// 把结构体原来的方法替换成新构造的这个 RPC 调用的方法
		s6StructFieldValue.Set(reflect.MakeFunc(s6StructFieldType, f8NewFunc))
	}
}
