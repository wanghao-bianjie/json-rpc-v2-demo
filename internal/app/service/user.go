package service

import (
	"myjsonrpcv2/internal/app/errors"
	"myjsonrpcv2/internal/app/model/entity"
)

type IUserService interface {
	FindUserByName(name string) (entity.User, errors.Error)
	AddUser(name string, age int) (uint, errors.Error)
}

type UserService struct {
}

func (svc UserService) FindUserByName(name string) (entity.User, errors.Error) {
	user, err := userRepo.FindByName(name)
	if err != nil {
		return entity.User{}, errors.Wrap(err)
	}
	return user, nil
}

func (svc UserService) AddUser(name string, age int) (uint, errors.Error) {
	user := entity.User{
		Name: name,
		Age:  age,
	}
	userId, err := userRepo.Add(&user)
	if err != nil {
		return 0, errors.Wrap(err)
	}
	return userId, nil
}
