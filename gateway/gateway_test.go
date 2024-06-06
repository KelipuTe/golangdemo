package gateway

import (
	"demo-golang/signal"
	"testing"
)

func TestGateway(t *testing.T) {
	g := NewGateway("gateway", 9601, 9602)
	g.Start()
	signal.WaitForSIGINT()
}
