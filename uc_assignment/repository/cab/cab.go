package cab

import (
	"net/http"
	"uc_assignment/helper"
	requestresponse "uc_assignment/helper/requestResponse"
	"uc_assignment/model"
)

var cabs = make(map[string]*model.Cab)
var cabsLocation = make(map[*model.Cab]*model.Location)
var cabUserMapping = make(map[*model.Cab]*model.User)
var lastCabId int

type CabRepository interface {
	CreateCab(model.Cab) (*model.Cab, helper.Error)
	GetAvailableCabs(source *model.Location, radius float64) ([]*model.Cab, helper.Error)
	UpdateLocation(cab *model.Cab, location *model.Location)
	CompleteTrip(trip *model.Trip) helper.Error
	GetCab(cabNumber string) (*model.Cab, helper.Error)
	ListCabs() ([]requestresponse.CabLocation, helper.Error)
	AssignCab(cab *model.Cab, driver *model.User) helper.Error
	MarkCabAvailability(cab *model.Cab, status bool)
	GetDriver(cab *model.Cab) *model.User
}

type Cab struct {
}

func NewCabRepository() CabRepository {
	return &Cab{}
}

func (c *Cab) CreateCab(cabParam model.Cab) (*model.Cab, helper.Error) {

	_, found := cabs[cabParam.RegistrationNumber]

	if found {
		return nil, helper.Error{Code: http.StatusForbidden, Message: helper.CAB_ALREADY_CREATED}
	}

	lastCabId += 1
	newCab := model.Cab{Id: lastCabId, RegistrationNumber: cabParam.RegistrationNumber, IsAvailable: false}
	cabs[cabParam.RegistrationNumber] = &newCab
	cabsLocation[&newCab] = &model.Location{}
	return &newCab, helper.Error{}
}

func (c *Cab) UpdateLocation(cab *model.Cab, location *model.Location) {

	cabsLocation[cab] = location
}

func (c *Cab) GetAvailableCabs(source *model.Location, radius float64) ([]*model.Cab, helper.Error) {

	var availableCabs []*model.Cab
	for cab, cabCurrentLocation := range cabsLocation {
		if helper.FindDistance(*source, *cabCurrentLocation) <= radius && cab.IsAvailable {
			availableCabs = append(availableCabs, cab)
		}
	}
	if len(availableCabs) == 0 {
		return nil, helper.Error{Code: http.StatusOK, Message: "No available cabs found nearby"}
	}
	return availableCabs, helper.Error{}
}

func (c *Cab) CompleteTrip(trip *model.Trip) helper.Error {

	if trip.Status == model.FINISHED {
		return helper.Error{Code: http.StatusBadRequest, Message: "Trip has already finished"}
	}

	trip.Status = model.FINISHED
	trip.Cab.IsAvailable = true
	return helper.Error{}

}

func (c *Cab) GetCab(registrationNumber string) (*model.Cab, helper.Error) {

	_, found := cabs[registrationNumber]

	if found {
		return cabs[registrationNumber], helper.Error{}
	}
	return nil, helper.Error{Code: http.StatusNotFound, Message: helper.CAB_NOT_FOUND}
}

func (c *Cab) ListCabs() ([]requestresponse.CabLocation, helper.Error) {

	var cabsList []requestresponse.CabLocation

	for _, cab := range cabs {
		driver := cabUserMapping[cab]
		cabsList = append(cabsList, requestresponse.CabLocation{Cab: cab, Location: cabsLocation[cab], Driver: driver})
	}
	if len(cabs) == 0 {
		return nil, helper.Error{Code: http.StatusOK, Message: "No cabs found"}
	}
	return cabsList, helper.Error{}
}

func (c *Cab) AssignCab(cab *model.Cab, driver *model.User) helper.Error {

	_, found := cabUserMapping[cab]
	if found {
		return helper.Error{Code: http.StatusForbidden, Message: "Cab already assigned driver"}
	}

	for _, user := range cabUserMapping {
		if user.Number == driver.Number {
			return helper.Error{Code: http.StatusForbidden, Message: "Driver already assigned"}
		}
	}

	cabUserMapping[cab] = driver
	cab.IsAvailable = true
	return helper.Error{}
}

func (c *Cab) MarkCabAvailability(cab *model.Cab, status bool) {
	cab.IsAvailable = status
}

func (c *Cab) GetDriver(cab *model.Cab) *model.User {
	return cabUserMapping[cab]
}
