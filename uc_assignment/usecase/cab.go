package usecase

import (
	"errors"
	"uc_assignment/model"
	repository "uc_assignment/repository/cab"
)

type CabUsecase interface {
	CabRepository(cabParam model.Cab) (model.Cab, error)
	UpdateLocation(cabId int) error
}

type CabUsecaseParam struct {
	CabRepository repository.CabRepository
}

type cabUsecase struct {
	cabRepository repository.CabRepository
}

func NewCabUsecase(param CabUsecaseParam) CabUsecase {
	return cabUsecase{
		cabRepository: param.CabRepository,
	}
}

func (cab cabUsecase) CabRepository(cabParam model.Cab) (model.Cab, error) {

	newCab, cabError := cab.cabRepository.CabRepository(cabParam)
	if cabError != nil {
		return model.Cab{}, errors.New(cabError.Error())
	}
	return newCab, nil
}

func (c cabUsecase) UpdateLocation(cabId int) error {

	cab, err := c.cabRepository.GetCab(cabId)
	if err != nil {
		return err
	}
	c.cabRepository.UpdateLocation(cab)
	return nil

}
