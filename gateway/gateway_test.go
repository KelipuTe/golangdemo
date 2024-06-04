package gateway

import (
	"demo-golang/signal"
	"testing"
)

func Test_Gateway(t *testing.T) {
	g := NewGateway("gateway", 9501, 9602)
	g.Start()
	signal.WaitForSIGINT()
}
