package export

import "github.com/osamingo/jsonrpc/v2"

type (
	Handler interface {
		jsonrpc.Handler
		Name() string
		Params() interface{}
		Result() interface{}
	}

	Method interface {
		MethodName(Handler) string
		Handlers() []Handler
	}
)
