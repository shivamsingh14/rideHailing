package user

import (
	"net/http"
	"uc_assignment/helper"
	requestresponse "uc_assignment/helper/requestResponse"
	"uc_assignment/model"
)

type Rider struct {
	role string
}

func NewRiderRepo() UserRepository {
	return &Rider{
		role: RIDER,
	}
}

func (r *Rider) Register(userParam model.User) (*model.User, helper.Error) {

	_, found := riderRecords[userParam.Number]

	if found {
		return nil, helper.Error{Code: http.StatusForbidden, Message: "User already exists"}
	}

	lastRiderID += 1
	newUser := model.User{Id: lastRiderID, Number: userParam.Number, Name: userParam.Name, Role: r.role}
	riderRecords[userParam.Number] = &newUser
	return &newUser, helper.Error{}
}

func (r *Rider) ListUsers() ([]*model.User, helper.Error) {

	var users []*model.User

	for _, user := range riderRecords {
		users = append(users, user)
	}
	if len(users) == 0 {
		return nil, helper.Error{Code: http.StatusOK, Message: "No Riders Registered"}
	}
	return users, helper.Error{}
}

func (r *Rider) GetUser(phoneNumber int) (*model.User, helper.Error) {

	rider, found := riderRecords[phoneNumber]

	if !found {
		return nil, helper.Error{Code: http.StatusNotFound, Message: " No Rider found"}
	}
	return rider, helper.Error{}
}

func (r *Rider) TripHistory(phoneNumber int) (requestresponse.TripHistory, helper.Error) {

	var ongoingTrips, completedTrips []*model.Trip

	// for _, trip := range trips {
	// 	if trip.Rider.Number == phoneNumber {
	// 		if trip.Status == model.IN_PROGRESS {
	// 			ongoingTrips = append(ongoingTrips, trip)
	// 		} else if trip.Status == model.FINISHED {
	// 			completedTrips = append(completedTrips, trip)
	// 		}
	// 	}
	// }
	ongoingTrips = riderOngoingTrips[phoneNumber]
	completedTrips = riderCompletedTrips[phoneNumber]
	return requestresponse.TripHistory{Ongoing: ongoingTrips, Completed: completedTrips}, helper.Error{}
}

func (r *Rider) CheckTripStatus(rider model.User) bool {

	_, found := riderOngoingTrips[rider.Number]

	return found

}

func (r *Rider) AddTrip(rider *model.User, trip *model.Trip) {
	_, found := riderOngoingTrips[rider.Number]
	if found {
		riderOngoingTrips[rider.Number] = append(riderOngoingTrips[rider.Number], trip)
		return
	}
	riderOngoingTrips[rider.Number] = []*model.Trip{trip}
}

func (r *Rider) CompleteTrip(rider *model.User, trip *model.Trip) {

	delete(riderOngoingTrips, rider.Number)
	_, found := riderCompletedTrips[rider.Number]
	if found {
		riderCompletedTrips[rider.Number] = append(riderCompletedTrips[rider.Number], trip)
		return
	}
	riderCompletedTrips[rider.Number] = []*model.Trip{trip}

}
