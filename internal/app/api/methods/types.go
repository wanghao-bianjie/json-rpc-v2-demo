package methods

import "myjsonrpcv2/internal/app/service"

var (
	userService service.IUserService = new(service.UserService)
)
