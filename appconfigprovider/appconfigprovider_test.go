package appconfigprovider

import (
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	options := []goleak.Option{
		goleak.IgnoreAnyFunction("github.com/open-feature/go-sdk/openfeature.(*eventExecutor).startEventListener.func1.1"),
	}
	goleak.VerifyTestMain(m, options...)
}
