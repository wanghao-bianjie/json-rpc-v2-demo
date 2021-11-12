package methods

import (
	"fmt"
	"myjsonrpcv2/internal/app/api"
	"myjsonrpcv2/internal/app/api/export"
)

const methodGroupNameUser = "User"

type User struct {
}

func init() {
	//register method
	api.RegisterMethod(&User{})
}

func (us *User) MethodName(h export.Handler) string {
	return fmt.Sprintf("%s.%s", methodGroupNameUser, h.Name())
}

func (us *User) Handlers() []export.Handler {
	//register method handlers
	return []export.Handler{
		&userEchoHandler{},
		&userPositionalHandler{},
		&userAddHandler{},
		&userDetailHandler{},
	}
}
