package delivery

import (
	"fmt"
	"net/http"
	customjson "uc_assignment/helper/json"
	"uc_assignment/model"

	"github.com/julienschmidt/httprouter"
)

func BookRide(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var errRider, errSource, errDestination error
	fmt.Fprint(w, "Welcome!\n")
	var rider model.User
	var source, destination model.Location
	errRider = customjson.Decode(r.Body, &rider)
	errSource = customjson.Decode(r.Body, &source)
	errDestination = customjson.Decode(r.Body, &destination)
	if errRider != nil || errSource != nil || errDestination != nil {
		fmt.Fprint(w, "Bad Request")
		return
	}

	ride, rideErr := g.bookingUsecase.BookRide(rider, source, destination)
	if rideErr != nil {
		fmt.Fprint(w, rideErr)
	}
	fmt.Fprint(w, ride)
}

func CompleteTrip(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var err error
	fmt.Fprint(w, "Welcome!\n")
	var trip model.Trip
	err = customjson.Decode(r.Body, &trip)
	if err != nil {
		fmt.Fprint(w, "Bad Request")
		return
	}
	tripPrice, errPrice := g.bookingUsecase.CompleteTrip(trip)
	if errPrice != nil {
		fmt.Fprint(w, errPrice)
	}
	fmt.Fprint(w, tripPrice)
}
