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

func (p7this *HTTPService) ServeHTTP(i9w http.ResponseWriter, p7r *http.Request) {
	p7ctx := &HTTPContext{
		I9writer:  i9w,
		P7request: p7r,
	}
	p7this.doServeHTTP(p7ctx)
}

func (p7this *HTTPService) doServeHTTP(p7ctx *HTTPContext) {

}

func (p7this *HTTPService) Start(addr string) error {
	return nil
}
