package delivery

import (
	"net/http"
	"uc_assignment/helper"
	customjson "uc_assignment/helper/json"
	requestresponse "uc_assignment/helper/requestResponse"
	"uc_assignment/model"

	"github.com/julienschmidt/httprouter"
)

func AddCoupon(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var coupon model.Coupon

	err := customjson.Decode(r.Body, &coupon)
	if err != nil || coupon.Code == "" {
		requestresponse.MakeResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}

	newCoupon, newCouponError := g.couponUsecase.AddCoupon(coupon)
	if newCouponError != (helper.Error{}) {
		requestresponse.MakeResponse(w, newCouponError.Code, newCouponError)
		return
	}
	requestresponse.MakeResponse(w, http.StatusCreated, newCoupon)
}

func DeleteCoupon(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var coupon model.Coupon

	err := customjson.Decode(r.Body, &coupon)
	if err != nil || coupon.Code == "" {
		requestresponse.MakeResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}

	deleteCouponError := g.couponUsecase.DeleteCoupon(coupon.Code)
	if deleteCouponError != (helper.Error{}) {
		requestresponse.MakeResponse(w, deleteCouponError.Code, deleteCouponError)
		return
	}
	requestresponse.MakeResponse(w, http.StatusCreated, "Coupon deleted successfully")
}

func AssignCoupon(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var assignCoupon requestresponse.AssignCoupon

	err := customjson.Decode(r.Body, &assignCoupon)
	if err != nil || assignCoupon.CouponCode == "" || assignCoupon.UserNumber == 0 {
		requestresponse.MakeResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	assignCouponError := g.couponUsecase.AssignCoupon(assignCoupon)
	if assignCouponError != (helper.Error{}) {
		requestresponse.MakeResponse(w, assignCouponError.Code, assignCouponError)
		return
	}
	requestresponse.MakeResponse(w, http.StatusCreated, "Coupon Assigned successfully")
}
