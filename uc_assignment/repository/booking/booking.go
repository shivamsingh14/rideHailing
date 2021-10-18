package bookRide

import (
	"net/http"
	"uc_assignment/helper"
	"uc_assignment/model"
)

var trips = make(map[int]*model.Trip)
var lastTripId int

type BookingRespository interface {
	BookRide(source *model.Location, destination *model.Location, cab *model.Cab, rider *model.User, driver *model.User, price float64) (*model.Trip, helper.Error)
	GetRide(tripId int) (*model.Trip, helper.Error)
}

type bookRide struct {
}

func NewBookRideRepo() BookingRespository {
	return &bookRide{}
}

func (b *bookRide) BookRide(source *model.Location, destination *model.Location, cab *model.Cab, rider *model.User, driver *model.User, price float64) (*model.Trip, helper.Error) {
	lastTripId += 1

	newTrip := model.Trip{
		ID:          lastTripId,
		Source:      source,
		Destination: destination,
		Cab:         cab,
		Rider:       rider,
		Driver:      driver,
		Status:      model.IN_PROGRESS,
		Cost:        price,
	}
	trips[lastTripId] = &newTrip
	return &newTrip, helper.Error{}
}

func (b *bookRide) GetRide(tripId int) (*model.Trip, helper.Error) {

	_, found := trips[tripId]

	if found {
		return trips[tripId], helper.Error{}
	}
	return nil, helper.Error{Code: http.StatusForbidden, Message: "the ride is not present"}

}
