package v2

import "net/http"

type HTTPContext struct {
	I9writer    http.ResponseWriter
	P7request   *http.Request
	S5pathParam map[string]string
}
