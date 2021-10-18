package user

import (
	"uc_assignment/helper"
	requestresponse "uc_assignment/helper/requestResponse"
	"uc_assignment/model"
)

var driverRecords = make(map[int]*model.User)
var riderRecords = make(map[int]*model.User)

var driverOngoingTrips = make(map[int][]*model.Trip)
var riderOngoingTrips = make(map[int][]*model.Trip)

var driverCompletedTrips = make(map[int][]*model.Trip)
var riderCompletedTrips = make(map[int][]*model.Trip)

var lastDriverID int
var lastRiderID int

const (
	DRIVER = "driver"
	RIDER  = "rider"
)

type UserRepository interface {
	Register(user model.User) (*model.User, helper.Error)
	ListUsers() ([]*model.User, helper.Error)
	GetUser(phoneNumber int) (*model.User, helper.Error)
	TripHistory(phoneNumber int) (requestresponse.TripHistory, helper.Error)
	CheckTripStatus(rider model.User) bool
	AddTrip(user *model.User, trip *model.Trip)
	CompleteTrip(user *model.User, trip *model.Trip)
}
