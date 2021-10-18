package delivery

import (
	"fmt"
	"net/http"
	"uc_assignment/helper"
	customjson "uc_assignment/helper/json"
	requestresponse "uc_assignment/helper/requestResponse"
	"uc_assignment/model"

	"github.com/julienschmidt/httprouter"
)

func BookRide(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var bookRide requestresponse.BookRide
	bookRideError := customjson.Decode(r.Body, &bookRide)
	if bookRideError != nil {
		requestresponse.MakeResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}

	ride, rideErr := g.bookingUsecase.BookRide(bookRide.Rider, bookRide.Source, bookRide.Destination)
	if rideErr != (helper.Error{}) {
		requestresponse.MakeResponse(w, rideErr.Code, rideErr)
		return
	}
	requestresponse.MakeResponse(w, http.StatusCreated, ride)
}

func CompleteTrip(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var trip *model.Trip
	err := customjson.Decode(r.Body, &trip)
	if err != nil {
		requestresponse.MakeResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	tripPrice, errPrice := g.bookingUsecase.CompleteTrip(trip)
	if errPrice != (helper.Error{}) {
		requestresponse.MakeResponse(w, errPrice.Code, errPrice)
		return
	}
	successResponse := fmt.Sprintf("Trip has been completed, Your fare is %v", tripPrice)
	requestresponse.MakeResponse(w, http.StatusOK, successResponse)
}
