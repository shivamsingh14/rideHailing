package delivery

import (
	"fmt"
	"net/http"
	customjson "uc_assignment/helper/json"
	"uc_assignment/model"

	"github.com/julienschmidt/httprouter"
)

func CabRepository(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var err error
	fmt.Fprint(w, "Welcome!\n")
	var cab model.Cab
	err = customjson.Decode(r.Body, &cab)
	if err != nil {
		fmt.Fprint(w, "Bad Request")
		return
	}
	newCab, cabError := g.cabuseCase.CabRepository(cab)
	if cabError != nil {
		fmt.Fprint(w, cabError)
	}
	fmt.Fprint(w, newCab)
}

func UpdateCabLocation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var err error
	fmt.Fprint(w, "Welcome!\n")
	var cab model.Cab
	err = customjson.Decode(r.Body, &cab)
	if err != nil {
		fmt.Fprint(w, "Bad Request")
		return
	}
	err = g.cabuseCase.UpdateLocation(cab.Id)
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, "Updated location Successfully")
}
