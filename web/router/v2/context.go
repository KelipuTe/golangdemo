package v2

import "net/http"

type HTTPContext struct {
	P7request   *http.Request
	I9writer    http.ResponseWriter
	S5pathParam map[string]string
}
