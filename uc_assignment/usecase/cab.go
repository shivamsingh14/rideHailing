package usecase

import (
	"uc_assignment/helper"
	requestresponse "uc_assignment/helper/requestResponse"
	"uc_assignment/model"
	repository "uc_assignment/repository/cab"
	userRepo "uc_assignment/repository/user"
)

type CabUsecase interface {
	CreateCab(cabParam model.Cab) (*model.Cab, helper.Error)
	UpdateLocation(cabNumber string, location *model.Location) helper.Error
	GetAvailableCabs(source model.Location, radius float64) ([]*model.Cab, helper.Error)
	ListCabs() ([]requestresponse.CabLocation, helper.Error)
	AssignCab(registrationNumber string, phoneNumber int) helper.Error
}

type CabUsecaseParam struct {
	CabRepository    repository.CabRepository
	DriverRepository userRepo.UserRepository
}

type cabUsecase struct {
	cabRepository    repository.CabRepository
	driverRepository userRepo.UserRepository
}

func NewCabUsecase(param CabUsecaseParam) CabUsecase {
	return cabUsecase{
		cabRepository:    param.CabRepository,
		driverRepository: param.DriverRepository,
	}
}

func (cab cabUsecase) CreateCab(cabParam model.Cab) (*model.Cab, helper.Error) {

	newCab, cabError := cab.cabRepository.CreateCab(cabParam)
	if cabError != (helper.Error{}) {
		return nil, cabError
	}
	return newCab, helper.Error{}
}

func (c cabUsecase) UpdateLocation(cabNumber string, location *model.Location) helper.Error {

	var newLocation model.Location
	cab, err := c.cabRepository.GetCab(cabNumber)
	if err != (helper.Error{}) {
		return err
	}
	if location != nil {
		c.cabRepository.UpdateLocation(cab, location)
		return helper.Error{}
	}
	newLocation = helper.GetCurrentLocation()
	c.cabRepository.UpdateLocation(cab, &newLocation)
	return helper.Error{}
}

func (c cabUsecase) GetAvailableCabs(source model.Location, radius float64) ([]*model.Cab, helper.Error) {
	return c.cabRepository.GetAvailableCabs(&source, radius)
}

func (c cabUsecase) ListCabs() ([]requestresponse.CabLocation, helper.Error) {
	cabs, err := c.cabRepository.ListCabs()
	if err != (helper.Error{}) {
		return []requestresponse.CabLocation{}, err
	}
	return cabs, helper.Error{}
}

func (c cabUsecase) AssignCab(registrationNumber string, phoneNumber int) helper.Error {

	cab, err := c.cabRepository.GetCab(registrationNumber)
	if err != (helper.Error{}) {
		return err
	}
	driver, getDriverErr := c.driverRepository.GetUser(phoneNumber)
	if getDriverErr != (helper.Error{}) {
		return getDriverErr
	}

	assignCabError := c.cabRepository.AssignCab(cab, driver)
	if assignCabError != (helper.Error{}) {
		return assignCabError
	}
	return helper.Error{}

}
