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

func CreateCab(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var cab model.Cab
	err := customjson.Decode(r.Body, &cab)
	if err != nil || cab.RegistrationNumber == "" {
		requestresponse.MakeResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	newCab, cabError := g.cabUsecase.CreateCab(cab)
	if cabError != (helper.Error{}) {
		requestresponse.MakeResponse(w, cabError.Code, cabError)
		return
	}
	requestresponse.MakeResponse(w, http.StatusCreated, newCab)
}

func UpdateCabLocation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var cabLocation requestresponse.CabLocation
	err := customjson.Decode(r.Body, &cabLocation)
	if err != nil {
		requestresponse.MakeResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	updateLocationErr := g.cabUsecase.UpdateLocation(cabLocation.Cab.RegistrationNumber, cabLocation.Location)
	if updateLocationErr != (helper.Error{}) {
		requestresponse.MakeResponse(w, updateLocationErr.Code, updateLocationErr)
		return
	}
	fmt.Fprint(w, "Updated location Successfully")
}

func GetAvailableCabs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var location model.Location
	err := customjson.Decode(r.Body, &location)
	if err != nil {
		requestresponse.MakeResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	availableCabs, availableCabsErr := g.cabUsecase.GetAvailableCabs(location, helper.RADIUS)
	if availableCabsErr != (helper.Error{}) {
		requestresponse.MakeResponse(w, availableCabsErr.Code, availableCabsErr)
		return
	}
	requestresponse.MakeResponse(w, http.StatusOK, availableCabs)

}

func ListCabs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cabsList, cabsListError := g.cabUsecase.ListCabs()
	if cabsListError != (helper.Error{}) {
		requestresponse.MakeResponse(w, cabsListError.Code, cabsListError)
		return
	}
	requestresponse.MakeResponse(w, http.StatusOK, cabsList)
}

func AssignCab(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var assignCab requestresponse.AssignCab
	err := customjson.Decode(r.Body, &assignCab)
	if err != nil {
		requestresponse.MakeResponse(w, http.StatusBadRequest, "")
		return
	}
	assignCabError := g.cabUsecase.AssignCab(assignCab.RegistrationNumber, assignCab.DriverNumber)
	if assignCabError != (helper.Error{}) {
		requestresponse.MakeResponse(w, assignCabError.Code, assignCabError)
		return
	}
	requestresponse.MakeResponse(w, http.StatusOK, "Cab assigned to the driver")
}
