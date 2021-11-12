package methods

import (
	"context"
	"encoding/json"
	"github.com/osamingo/jsonrpc/v2"
	"myjsonrpcv2/internal/app/api/export"
	"myjsonrpcv2/internal/app/api/types"
)

const methodRouteNameUserPositional = "Positional"

type userPositionalHandler struct{}

var _ export.Handler = &userPositionalHandler{}

func (h userPositionalHandler) Name() string {
	return methodRouteNameUserPositional
}

func (h userPositionalHandler) Params() interface{} {
	return &types.PositionalParams{}
}

func (h userPositionalHandler) Result() interface{} {
	return &types.PositionalResult{}
}

func (h userPositionalHandler) ServeJSONRPC(c context.Context, params *json.RawMessage) (interface{}, *jsonrpc.Error) {
	var p types.PositionalParams
	if err := jsonrpc.Unmarshal(params, &p); err != nil {
		return nil, err
	}
	return types.PositionalResult{
		Message: p,
	}, nil
}
