package delivery

import (
	"fmt"
	"net/http"
	customjson "uc_assignment/helper/json"
	"uc_assignment/model"

	"github.com/julienschmidt/httprouter"
)

func RegisterCab(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var err error
	fmt.Fprint(w, "Welcome!\n")
	var cab model.Cab
	err = customjson.Decode(r.Body, &cab)
	if err != nil {
		//TODO: fmt.Fprint(w, helper.BadError("code", "message"))
		fmt.Fprint(w, "Bad Request")
		return
	}
	newCab, cabError := g.cabuseCase.RegisterCab(cab)
	if cabError != nil {
		//TODO: fmt.Fprint(w, helper.BadError("code", "message"))
		fmt.Fprint(w, cabError)
	}
	fmt.Fprint(w, newCab)
}
