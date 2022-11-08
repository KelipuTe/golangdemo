package selectkn

import "fmt"

func NewErrUnsupportedExpressionType(e any) error {
	return fmt.Errorf("orm: 不支持的表达式 %v", e)
}
