package repository

import "uc_assignment/model"

var trips = make(map[int]model.Trip)
var lastTripId int

type BookingRespository interface {
	BookRide(source model.Location, destination model.Location, cab model.Cab, rider model.User, price float64) (model.Trip, error)
}

type BookRide struct {
}

func NewBookRideRepo() BookingRespository {
	return BookRide{}
}

func (b BookRide) BookRide(source model.Location, destination model.Location, cab model.Cab, rider model.User, price float64) (model.Trip, error) {
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
