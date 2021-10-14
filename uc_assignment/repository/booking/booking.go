package repository

import (
	"errors"
	"uc_assignment/model"
)

var trips = make(map[int]model.Trip)
var lastTripId int

type BookingRespository interface {
	BookRide(source model.Location, destination model.Location, cab model.Cab, rider model.User, price float64) (model.Trip, error)
	RideHistory() ([]model.Trip, error)
	GetRide(tripId int) (model.Trip, error)
}

type bookRide struct {
}

func NewBookRideRepo() BookingRespository {
	return bookRide{}
}

func (b bookRide) BookRide(source model.Location, destination model.Location, cab model.Cab, rider model.User, price float64) (model.Trip, error) {
	lastTripId += 1

	newTrip := model.Trip{
		Source:      source,
		Destination: destination,
		Cab:         cab,
		Rider:       rider,
		Status:      model.IN_PROGRESS,
		Cost:        price,
	}
	trips[lastTripId] = newTrip
	return newTrip, nil
}

func (b bookRide) RideHistory() ([]model.Trip, error) {

	var rideHistory []model.Trip

	for _, trip := range trips {

		rideHistory = append(rideHistory, trip)

	}
	return rideHistory, nil

}

func (b bookRide) GetRide(tripId int) (model.Trip, error) {

	_, found := trips[tripId]

	if found {
		return trips[tripId], nil
	}
	return model.Trip{}, errors.New("the ride is not present")

}
