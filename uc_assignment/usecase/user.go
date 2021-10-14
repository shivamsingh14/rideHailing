package usecase

import (
	"errors"
	"strings"
	"uc_assignment/helper"
	"uc_assignment/model"
	repository "uc_assignment/repository/user"
)

type UserUsecase interface {
	RegisterUser(user model.User, role string) (model.User, error)
}

type UserUsecaseParam struct {
	DriverRepository repository.UserRepository
	RiderRepository  repository.UserRepository
}

type userUsecase struct {
	driverRepository repository.UserRepository
	riderRepository  repository.UserRepository
}

func NewUserUsecase(param UserUsecaseParam) UserUsecase {
	return userUsecase{
		driverRepository: param.DriverRepository,
		riderRepository:  param.RiderRepository,
	}
}

func (u userUsecase) RegisterUser(userParam model.User, role string) (model.User, error) {

	var user model.User
	var err error

	if role == helper.DRIVER {
		user, err = u.driverRepository.Register(userParam)
	} else if role == helper.RIDER {
		user, err = u.riderRepository.Register(userParam)
	}

	if err != nil {
		return model.User{}, errors.New("could not create " + strings.ToLower(role))
	}
	return user, nil
}
