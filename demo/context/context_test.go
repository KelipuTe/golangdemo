package context

import "testing"

func Test_f8ContextWithCancel(p7s6t *testing.T) {
	f8ContextWithCancel()
}

func Test_f8ContextWithDeadline(p7s6t *testing.T) {
	f8ContextWithDeadline()
	f8ContextWithDeadlineV2()
}

func Test_f8ContextWithTimeout(p7s6t *testing.T) {
	f8ContextWithTimeout()
	f8ContextWithTimeoutV2()
}

func Test_f8ContextWithValue(p7s6t *testing.T) {
	f8ContextWithValue()
}
