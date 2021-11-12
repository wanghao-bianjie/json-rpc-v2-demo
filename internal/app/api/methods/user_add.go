package methods

import (
	"context"
	"encoding/json"
	"github.com/osamingo/jsonrpc/v2"
	"myjsonrpcv2/internal/app/api/export"
	"myjsonrpcv2/internal/app/api/types"
)

const methodRouteNameUserAdd = "Add"

type userAddHandler struct{}

var _ export.Handler = &userAddHandler{}

func (h userAddHandler) Name() string {
	return methodRouteNameUserAdd;
}

func (h userAddHandler) Params() interface{} {
	return &types.AddParams{}
}

func (h userAddHandler) Result() interface{} {
	return &types.AddResult{}
}

func (h userAddHandler) ServeJSONRPC(c context.Context, params *json.RawMessage) (interface{}, *jsonrpc.Error) {
	var q types.AddParams
	if err := jsonrpc.Unmarshal(params, &q); err != nil {
		return nil, err
	}

	userId, err := userService.AddUser(q.Name, q.Age)
	if err != nil {
		return nil, &jsonrpc.Error{
			Code:    jsonrpc.ErrorCodeInternal,
			Message: err.Msg(),
		}
	}
	return types.AddResult{
		Id: userId,
	}, nil
}
