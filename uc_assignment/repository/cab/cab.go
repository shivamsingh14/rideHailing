package cab

import (
	"errors"
	"uc_assignment/helper"
	"uc_assignment/model"
)

var cabs = make(map[int]model.Cab)
var cabsLocation = make(map[model.Cab]model.Location)
var lastCabId int

type CabRepository interface {
	RegisterCab(model.Cab) (model.Cab, error)
	GetAvailableCabs(source model.Location, radius float64) ([]model.Cab, error)
	UpdateLocation(model.Cab, model.Location) error
	CompleteTrip(trip model.Trip)
}

type Cab struct {
}

func NewCabRepository() CabRepository {
	return Cab{}
}

func (c Cab) RegisterCab(cabParam model.Cab) (model.Cab, error) {
	_, found := cabs[cabParam.Id]

	if found {
		return model.Cab{}, errors.New("cab already exists")
	}

	lastCabId += 1
	newCab := model.Cab{Id: lastCabId, License: cabParam.License, IsAvailable: true}
	cabs[lastCabId] = newCab
	return newCab, nil
}

func (c Cab) UpdateLocation(cab model.Cab, location model.Location) error {

	_, found := cabsLocation[cab]

	if !found {
		return errors.New("invalid cab request")
	}
	cabsLocation[cab] = helper.GetCurrentLocation()
	return nil
}

func (c Cab) GetAvailableCabs(source model.Location, radius float64) ([]model.Cab, error) {

	var availableCabs []model.Cab

	for cab, cabCurrentLocation := range cabsLocation {
		if helper.FindDistance(source, cabCurrentLocation) <= radius {
			availableCabs = append(availableCabs, cab)
		}
	}
	return availableCabs, nil
}

func (c Cab) CompleteTrip(trip model.Trip) {

	trip.Status = model.FINISHED

}
