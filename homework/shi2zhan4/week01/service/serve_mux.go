package service

import "net/http"

type knHTTPServeMux struct {
	isStop bool
	*http.ServeMux
}

func (p1this *knHTTPServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if p1this.isStop {
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte("服务已关闭"))
		return
	}
	p1this.ServeMux.ServeHTTP(w, r)
}
