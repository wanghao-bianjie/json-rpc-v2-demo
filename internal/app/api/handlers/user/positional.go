package user

import (
	"context"
	"encoding/json"
	"github.com/osamingo/jsonrpc/v2"
	export2 "myjsonrpcv2/internal/app/api/export"
	types2 "myjsonrpcv2/internal/app/api/types"
)

const methodRouteNamePositional = "Positional"

type PositionalHandler struct{}

var _ export2.Handler = &PositionalHandler{}

func (h PositionalHandler) Name() string {
	return methodRouteNamePositional
}

func (h PositionalHandler) Params() interface{} {
	return &types2.PositionalParams{}
}

func (h PositionalHandler) Result() interface{} {
	return &types2.PositionalResult{}
}

func (h PositionalHandler) ServeJSONRPC(c context.Context, params *json.RawMessage) (interface{}, *jsonrpc.Error) {
	var p types2.PositionalParams
	if err := jsonrpc.Unmarshal(params, &p); err != nil {
		return nil, err
	}
	return types2.PositionalResult{
		Message: p,
	}, nil
}
