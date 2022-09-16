## 使用 Go 的 HTTP 包开启 HTTP 服务

- 2022-09-15
- go version go1.19

在 Go 中有多种方式，可以实现 HTTP 服务。可以使用官方提供的封装好的 `net/http` 包，也可以直接使用 net 包从 TCP 开始自行实现。

使用 `net/http` 包的时候，需要关注的最核心的部分，就是 Handler 接口 (src/net/http/server.go)。

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

开启 HTTP 服务和 HTTP 请求被处理的流程大致如下：

- 直接或间接地创建 Server 结构体 (src/net/http/server.go) 的实例。
- 调用 Server 的 ListenAndServe() 方法。
- ListenAndServe() 调用 net.Listen() (src/net/dial.go)，启动 TCP 服务。
- net.Listen() 返回一个 net.Listener 接口 (src/net/dial.go) 的实例。
- 调用 net.Listener 的 Accept() 方法，获取连接上来的 TCP 连接。
- 新开启一个协程，把这个 TCP 连接丢进去处理。
- 继续往下，会遇到这行代码：`serverHandler{c.server}.ServeHTTP(w, w.req)`。
- 这里调用的就是 Handler 的 ServeHTTP() 方法。
