package methods

import (
	"context"
	"encoding/json"
	"github.com/osamingo/jsonrpc/v2"
	"myjsonrpcv2/internal/app/api/export"
	"myjsonrpcv2/internal/app/api/types"
)

const methodRouteNameUserDetail = "Detail"

type userDetailHandler struct{}

var _ export.Handler = &userDetailHandler{}

func (h userDetailHandler) Name() string {
	return methodRouteNameUserDetail;
}

func (h userDetailHandler) Params() interface{} {
	return &types.DetailParams{}
}

func (h userDetailHandler) Result() interface{} {
	return &types.DetailResult{}
}

func (h userDetailHandler) ServeJSONRPC(c context.Context, params *json.RawMessage) (interface{}, *jsonrpc.Error) {
	var q types.DetailParams
	if err := jsonrpc.Unmarshal(params, &q); err != nil {
		return nil, err
	}

	user, err := userService.FindUserByName(q.Name)
	if err != nil {
		return nil, &jsonrpc.Error{
			Code:    jsonrpc.ErrorCodeInternal,
			Message: err.Msg(),
		}
	}
	return types.DetailResult{
		Name: user.Name,
		Age:  user.Age,
	}, nil
}
