package methods

import (
	"context"
	"encoding/json"
	"github.com/osamingo/jsonrpc/v2"
	"myjsonrpcv2/internal/app/api/export"
	"myjsonrpcv2/internal/app/api/types"
)

const methodRouteNameUserEcho = "Echo"

type userEchoHandler struct{}

var _ export.Handler = &userEchoHandler{}

func (h userEchoHandler) Name() string {
	return methodRouteNameUserEcho;
}

func (h userEchoHandler) Params() interface{} {
	return &types.EchoParams{}
}

func (h userEchoHandler) Result() interface{} {
	return &types.EchoResult{}
}

func (h userEchoHandler) ServeJSONRPC(c context.Context, params *json.RawMessage) (interface{}, *jsonrpc.Error) {
	var p types.EchoParams
	if err := jsonrpc.Unmarshal(params, &p); err != nil {
		return nil, err
	}

	return types.EchoResult{
		Message: "Hello, " + p.Name,
	}, nil
}
