package user

import (
	"net/http"
	"uc_assignment/helper"
	requestresponse "uc_assignment/helper/requestResponse"
	"uc_assignment/model"
)

type Driver struct {
	role string
}

func NewDriverRepo() UserRepository {
	return &Driver{
		role: DRIVER,
	}
}

func (d *Driver) Register(userParam model.User) (*model.User, helper.Error) {

	_, found := driverRecords[userParam.Number]

	if found {
		return nil, helper.Error{Code: http.StatusForbidden, Message: "already existing driver"}
	}
	lastDriverID += 1
	newDriver := model.User{Id: lastDriverID, Number: userParam.Number, Name: userParam.Name, Role: d.role}
	driverRecords[userParam.Number] = &newDriver

	return &newDriver, helper.Error{}

}

func (d *Driver) ListUsers() ([]*model.User, helper.Error) {

	var users []*model.User

	for _, user := range driverRecords {
		users = append(users, user)
	}
	if len(users) == 0 {
		return nil, helper.Error{Code: http.StatusOK, Message: "No Drivers Registered"}
	}
	return users, helper.Error{}

}

func (d *Driver) GetUser(phoneNumber int) (*model.User, helper.Error) {

	driver, found := driverRecords[phoneNumber]

	if !found {
		return nil, helper.Error{Code: http.StatusNotFound, Message: " No driver found"}
	}
	return driver, helper.Error{}
}

func (d *Driver) TripHistory(phoneNumber int) (requestresponse.TripHistory, helper.Error) {

	var ongoingTrips, completedTrips []*model.Trip

	ongoingTrips = driverOngoingTrips[phoneNumber]
	completedTrips = driverCompletedTrips[phoneNumber]
	return requestresponse.TripHistory{Ongoing: ongoingTrips, Completed: completedTrips}, helper.Error{}
}

func (d *Driver) CheckTripStatus(rider model.User) bool {

	_, found := driverOngoingTrips[rider.Number]
	return found

}

func (d *Driver) AddTrip(driver *model.User, trip *model.Trip) {

	_, found := driverOngoingTrips[driver.Number]
	if found {
		driverOngoingTrips[driver.Number] = append(driverOngoingTrips[driver.Number], trip)
	}
	driverOngoingTrips[driver.Number] = []*model.Trip{trip}

}

func (d *Driver) CompleteTrip(driver *model.User, trip *model.Trip) {

	delete(driverOngoingTrips, driver.Number)
	_, found := driverCompletedTrips[driver.Number]
	if found {
		driverCompletedTrips[driver.Number] = append(driverCompletedTrips[driver.Number], trip)
		return
	}
	driverCompletedTrips[driver.Number] = []*model.Trip{trip}

}
