package usecase

import (
	"uc_assignment/helper"
	requestresponse "uc_assignment/helper/requestResponse"
	"uc_assignment/model"
	bookingRepo "uc_assignment/repository/booking"
	userRepo "uc_assignment/repository/user"
)

var userRepository userRepo.UserRepository

type UserUsecase interface {
	RegisterUser(user model.User, role string) (*model.User, helper.Error)
	ListUsers(role string) ([]*model.User, helper.Error)
	GetRideHistory(role string, phoneNumber int) (requestresponse.TripHistory, helper.Error)
}

type UserUsecaseParam struct {
	DriverRepository  userRepo.UserRepository
	RiderRepository   userRepo.UserRepository
	BookingRepository bookingRepo.BookingRespository
}

type userUsecase struct {
	driverRepository  userRepo.UserRepository
	riderRepository   userRepo.UserRepository
	bookingRepository bookingRepo.BookingRespository
}

func NewUserUsecase(param UserUsecaseParam) UserUsecase {
	return userUsecase{
		driverRepository:  param.DriverRepository,
		riderRepository:   param.RiderRepository,
		bookingRepository: param.BookingRepository,
	}
}

func (u userUsecase) RegisterUser(userParam model.User, role string) (*model.User, helper.Error) {

	var user *model.User
	var err helper.Error

	if role == helper.DRIVER {
		userRepository = u.driverRepository
	} else if role == helper.RIDER {
		userRepository = u.riderRepository
	}
	user, err = userRepository.Register(userParam)

	if err != (helper.Error{}) {
		return nil, err
	}
	return user, helper.Error{}
}

func (u userUsecase) ListUsers(role string) ([]*model.User, helper.Error) {

	if role == helper.DRIVER {
		userRepository = u.driverRepository
	} else if role == helper.RIDER {
		userRepository = u.riderRepository
	}

	users, err := userRepository.ListUsers()
	if err != (helper.Error{}) {
		return nil, err
	}
	return users, helper.Error{}
}

func (u userUsecase) GetUser(phoneNumber int, role string) (*model.User, helper.Error) {

	if role == helper.DRIVER {
		userRepository = u.driverRepository
	} else if role == helper.RIDER {
		userRepository = u.riderRepository
	}

	user, err := userRepository.GetUser(phoneNumber)
	if err != (helper.Error{}) {
		return nil, err
	}
	return user, helper.Error{}
}

func (u userUsecase) GetRideHistory(role string, phoneNumber int) (requestresponse.TripHistory, helper.Error) {

	if role == helper.DRIVER {
		userRepository = u.driverRepository
	} else if role == helper.RIDER {
		userRepository = u.riderRepository
	}

	_, getUserError := userRepository.GetUser(phoneNumber)
	if getUserError != (helper.Error{}) {
		return requestresponse.TripHistory{}, getUserError
	}

	if role == helper.RIDER {
		return u.riderRepository.TripHistory(phoneNumber)
	}
	return u.driverRepository.TripHistory(phoneNumber)
}
