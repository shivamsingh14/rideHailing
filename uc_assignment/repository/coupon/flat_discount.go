package coupon

import (
	"net/http"
	"uc_assignment/helper"
	"uc_assignment/model"
)

type flatDiscount struct {
	FLAT_DISCOUNT int
}

func NewFlatDiscountRepo() CouponRepository {
	return &flatDiscount{
		FLAT_DISCOUNT: FLAT_DICOUNT,
	}
}

func (f *flatDiscount) AddCoupon(couponParam model.Coupon) (*model.Coupon, helper.Error) {

	_, found := coupons[couponParam.Code]
	if found {
		return nil, helper.Error{Code: http.StatusForbidden, Message: "the coupon code already exists"}
	}
	lastCouponId += 1
	newCoupon := model.Coupon{ID: lastCouponId, Code: couponParam.Code, Name: couponParam.Name, FlatDiscount: f.FLAT_DISCOUNT}
	coupons[couponParam.Code] = &newCoupon
	return &newCoupon, helper.Error{}
}

func (f *flatDiscount) DeleteCoupon(couponCode string) helper.Error {
	_, found := coupons[couponCode]
	if !found {
		return helper.Error{Code: http.StatusForbidden, Message: "the coupon does not  exists"}
	}
	delete(coupons, couponCode)
	delete(couponUsers, couponCode)
	return helper.Error{}
}

func (f *flatDiscount) AssignCoupon(rider *model.User, couponCode string) helper.Error {

	_, getCouponError := f.GetCoupon(couponCode)
	if getCouponError != (helper.Error{}) {
		return helper.Error{Code: http.StatusForbidden, Message: getCouponError.Message}
	}

	users, found := couponUsers[couponCode]
	if !found {
		couponUsers[couponCode] = []*model.User{rider}
		return helper.Error{}
	}
	for _, user := range users {
		if user == rider {
			return helper.Error{Code: http.StatusForbidden, Message: "The coupon is already assigned to the user"}
		}
	}
	couponUsers[couponCode] = append(couponUsers[couponCode], rider)
	return helper.Error{}
}

func (f *flatDiscount) GetCoupon(couponCode string) (*model.Coupon, helper.Error) {

	coupon, found := coupons[couponCode]
	if found {
		return coupon, helper.Error{}
	}
	return nil, helper.Error{Code: http.StatusNotFound, Message: "No such coupon exists"}
}
