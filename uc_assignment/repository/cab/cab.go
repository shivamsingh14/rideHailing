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
	CabRepository(model.Cab) (model.Cab, error)
	GetAvailableCabs(source model.Location, radius float64) ([]model.Cab, error)
	UpdateLocation(model.Cab) error
	CompleteTrip(trip model.Trip)
	GetCab(cabId int) (model.Cab, error)
}

type Cab struct {
}

func NewCabRepository() CabRepository {
	return Cab{}
}

func (c Cab) CabRepository(cabParam model.Cab) (model.Cab, error) {
	_, found := cabs[cabParam.Id]

	if found {
		return model.Cab{}, errors.New("cab already exists")
	}

	lastCabId += 1
	newCab := model.Cab{Id: lastCabId, License: cabParam.License, IsAvailable: true}
	cabs[lastCabId] = newCab
	return newCab, nil
}

func (c Cab) UpdateLocation(cab model.Cab) error {

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

func (c Cab) GetCab(cabId int) (model.Cab, error) {

	_, found := cabs[cabId]

	if found {
		return cabs[cabId], nil
	}
	return model.Cab{}, errors.New("the cab is not present")
}
