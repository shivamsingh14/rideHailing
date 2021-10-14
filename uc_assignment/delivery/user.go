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
		//TODO: fmt.Fprint(w, helper.BadError("code", "message"))
		fmt.Fprint(w, "Bad Request")
		return
	}
	newRider, riderError := g.userUsecase.RegisterUser(user, helper.RIDER)
	if riderError != nil {
		//TODO: fmt.Fprint(w, helper.BadError("code", "message"))
		fmt.Fprint(w, "Bad Request")
	}
	fmt.Fprint(w, newRider)
}

func RegisterDriver(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var err error
	fmt.Fprint(w, "Welcome!\n")
	var user model.User
	err = customjson.Decode(r.Body, &user)
	if err != nil || user.Name == "" {
		//fmt.Fprint(w, helper.BadError("code", "message"))
		fmt.Fprint(w, "Bad Request")
		return
	}
	newUser, userError := g.userUsecase.RegisterUser(user, helper.DRIVER)
	if userError != nil {
		//fmt.Fprint(w, helper.BadError("code", "message"))
		fmt.Fprint(w, "Bad Request")
	}
	fmt.Fprint(w, newUser)
}
