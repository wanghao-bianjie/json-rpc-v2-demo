package repository

import (
	"myjsonrpcv2/internal/app/db"
	"myjsonrpcv2/internal/app/model/entity"
)

type IUserRepo interface {
	FindByName(name string) (entity.User, error)
	Add(user *entity.User) (uint, error)
}

type UserRepo struct {
}

func (repo UserRepo) FindByName(name string) (entity.User, error) {
	var user entity.User
	//err := repo.db.Where(&entity.User{Name: name}).First(&user).Error
	err := db.GetDb().Where(&entity.User{Name: name}).First(&user).Error
	return user, err
}

func (repo UserRepo) Add(user *entity.User) (uint, error) {
	err := db.GetDb().Save(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}
