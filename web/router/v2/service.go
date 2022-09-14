package v2

import "net/http"

type HTTPHandleFunc func(p7ctx *HTTPContext)

type HTTPServiceInterface interface {
	RouterInterface
	http.Handler
	Start(addr string) error
}

type HTTPService struct {
	router
}

// 确保 HTTPService 实现了 HTTPServiceInterface 接口
var _ HTTPServiceInterface = &HTTPService{}

func NewHTTPService() *HTTPService {
	return &HTTPService{
		router: newRouter(),
	}
}

func (p7this *HTTPService) Get(path string, f4h HTTPHandleFunc) {
	p7this.router.addRoute(http.MethodGet, path, f4h)
}

func (p7this *HTTPService) Post(path string, f4h HTTPHandleFunc) {
	p7this.router.addRoute(http.MethodPost, path, f4h)
}

// ServeHTTP Handler.ServeHTTP，把 HTTPService 结构体变成 src/net/http/service.go 里 Handler 接口的实例。
// 在调用 http.ListenAndServe(addr string, handler Handler) 的时候，会把 HTTPService 的实例作为 handler 参数传进去。
// http.ListenAndServe() 会创建一个 src/net/http/server.go 里 Server 结构体的实例，保存 handler 参数。
// 然后 http.ListenAndServe() 调用 net.Listen(network, address string)，启动 TCP 服务。
// net.Listen() 返回一个 net.Listener 接口的实例，net.Listener 实例通过 Accept() 方法获取 TCP 连接。
// 获取到 TCP 连接之后，经过一系列的操作，最后会有这么一行代码 serverHandler{c.server}.ServeHTTP(w, w.req)。
// 这行代码会把一开始传进去的 Handler 接口的实例（HTTPService 的实例）取出来，然后调用 ServeHTTP 方法。
func (p7this *HTTPService) ServeHTTP(i9w http.ResponseWriter, p7r *http.Request) {
	i9ctx := &HTTPContext{
		P7request: p7r,
		I9writer:  i9w,
	}
	p7this.doServeHTTP(i9ctx)
}

func (p7this *HTTPService) doServeHTTP(*HTTPContext) {

}

func (p7this *HTTPService) Start(addr string) error {
	return nil
}
