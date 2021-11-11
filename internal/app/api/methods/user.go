package methods

import (
	"fmt"
	api2 "myjsonrpcv2/internal/app/api"
	export2 "myjsonrpcv2/internal/app/api/export"
	user2 "myjsonrpcv2/internal/app/api/handlers/user"
)

const methodGroupNameUser = "User"

type User struct{}

func init() {
	//register method
	api2.RegisterMethod(&User{})
}

func (us *User) MethodName(h export2.Handler) string {
	return fmt.Sprintf("%s.%s", methodGroupNameUser, h.Name())
}

func (us *User) Handlers() []export2.Handler {
	//register method handlers
	return []export2.Handler{
		&user2.EchoHandler{},
		&user2.PositionalHandler{},
	}
}
