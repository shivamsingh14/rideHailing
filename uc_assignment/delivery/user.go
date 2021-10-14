package delivery

import (
	"fmt"
	"net/http"
	"uc_assignment/helper"
	customjson "uc_assignment/helper/json"
	"uc_assignment/model"

	"github.com/julienschmidt/httprouter"
)

func RegisterRider(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var err error
	fmt.Fprint(w, "Welcome!\n")
	var user model.User
	err = customjson.Decode(r.Body, &user)
	if err != nil || user.Name == "" {
		fmt.Fprint(w, "Bad Request")
		return
	}
	newRider, riderError := g.userUsecase.RegisterUser(user, helper.RIDER)
	if riderError != nil {
		fmt.Fprint(w, riderError)
	}
	fmt.Fprint(w, newRider)
}

func RegisterDriver(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var err error
	fmt.Fprint(w, "Welcome!\n")
	var user model.User
	err = customjson.Decode(r.Body, &user)
	if err != nil || user.Name == "" {
		fmt.Fprint(w, "Bad Request")
		return
	}
	newUser, userError := g.userUsecase.RegisterUser(user, helper.DRIVER)
	if userError != nil {
		fmt.Fprint(w, userError)
	}
	fmt.Fprint(w, newUser)
}

func DriverTripHistory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var err error
	fmt.Fprint(w, "Welcome!\n")
	var driver model.User
	err = customjson.Decode(r.Body, &driver)
	if err != nil {
		fmt.Fprint(w, "Bad Request")
		return
	}
	driverTrips, driverTripsError := g.bookingUsecase.DriverTripHistory(driver)
	if driverTripsError != nil {
		fmt.Fprint(w, driverTripsError)
	}
	fmt.Fprint(w, driverTrips)
}

func RiderTripHistory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var err error
	fmt.Fprint(w, "Welcome!\n")
	var rider model.User
	err = customjson.Decode(r.Body, &rider)
	if err != nil {
		fmt.Fprint(w, "Bad Request")
		return
	}
	riderTrips, riderTripsError := g.bookingUsecase.RiderTripHistory(rider)
	if riderTripsError != nil {
		fmt.Fprint(w, riderTripsError)
	}
	fmt.Fprint(w, riderTrips)
}
