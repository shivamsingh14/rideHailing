package delivery

import (
	"net/http"
	"uc_assignment/helper"
	customjson "uc_assignment/helper/json"
	requestresponse "uc_assignment/helper/requestResponse"
	"uc_assignment/model"

	"github.com/julienschmidt/httprouter"
)

func RegisterRider(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var user model.User

	err := customjson.Decode(r.Body, &user)
	if err != nil || user.Number == 0 {
		requestresponse.MakeResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}

	newRider, riderError := g.userUsecase.RegisterUser(user, helper.RIDER)
	if riderError != (helper.Error{}) {
		requestresponse.MakeResponse(w, riderError.Code, riderError)
		return
	}
	requestresponse.MakeResponse(w, http.StatusCreated, newRider)
}

func ListRider(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ridersList, riderError := g.userUsecase.ListUsers(helper.RIDER)
	if riderError != (helper.Error{}) {
		requestresponse.MakeResponse(w, riderError.Code, riderError)
		return
	}
	requestresponse.MakeResponse(w, http.StatusOK, ridersList)
}

func RegisterDriver(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var user model.User

	err := customjson.Decode(r.Body, &user)
	if err != nil || user.Number == 0 {
		requestresponse.MakeResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}

	newDriver, driverError := g.userUsecase.RegisterUser(user, helper.DRIVER)
	if driverError != (helper.Error{}) {
		requestresponse.MakeResponse(w, driverError.Code, driverError)
		return
	}
	requestresponse.MakeResponse(w, http.StatusCreated, newDriver)
}

func ListDriver(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	driversList, driverError := g.userUsecase.ListUsers(helper.DRIVER)
	if driverError != (helper.Error{}) {
		requestresponse.MakeResponse(w, driverError.Code, driverError)
		return
	}
	requestresponse.MakeResponse(w, http.StatusOK, driversList)
}

func DriverTripHistory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var driver model.User
	err := customjson.Decode(r.Body, &driver)
	if err != nil {
		requestresponse.MakeResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	driverTrips, driverTripsError := g.userUsecase.GetRideHistory(helper.DRIVER, driver.Number)
	if driverTripsError != (helper.Error{}) {
		requestresponse.MakeResponse(w, driverTripsError.Code, driverTripsError)
		return
	}
	requestresponse.MakeResponse(w, http.StatusOK, driverTrips)
}

func RiderTripHistory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var rider model.User
	err := customjson.Decode(r.Body, &rider)
	if err != nil {
		requestresponse.MakeResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	riderTrips, riderTripsError := g.userUsecase.GetRideHistory(helper.RIDER, rider.Number)
	if riderTripsError != (helper.Error{}) {
		requestresponse.MakeResponse(w, riderTripsError.Code, riderTripsError)
		return
	}
	requestresponse.MakeResponse(w, http.StatusOK, riderTrips)
}
