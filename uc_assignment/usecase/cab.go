package usecase

import (
	"errors"
	"uc_assignment/model"
	repository "uc_assignment/repository/cab"
)

type CabUsecase interface {
	RegisterCab(cabParam model.Cab) (model.Cab, error)
	UpdateLocation()
	UpdateCabAvailability()
}

type CabUsecaseParam struct {
	RegisterCab repository.CabRepository
}

type cabUsecase struct {
	registerCab repository.CabRepository
}

func NewCabUsecase(param CabUsecaseParam) CabUsecase {
	return cabUsecase{
		registerCab: param.RegisterCab,
	}
}

func (cab cabUsecase) RegisterCab(cabParam model.Cab) (model.Cab, error) {

	newCab, cabError := cab.registerCab.RegisterCab(cabParam)
	if cabError != nil {
		return model.Cab{}, errors.New(cabError.Error())
	}
	return newCab, nil
}

func (cab cabUsecase) UpdateLocation() {

}

func (cab cabUsecase) UpdateCabAvailability() {

}
