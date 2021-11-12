package service

import "myjsonrpcv2/internal/app/repository"

var (
	userRepo repository.IUserRepo = new(repository.UserRepo)
)
