package coupon

import (
	"uc_assignment/helper"
	"uc_assignment/model"
)

var FLAT_DICOUNT = 10
var coupons = make(map[string]*model.Coupon)
var couponUsers = make(map[string][]*model.User)
var lastCouponId int

type CouponRepository interface {
	AddCoupon(couponParam model.Coupon) (*model.Coupon, helper.Error)
	GetCoupon(couponCode string) (*model.Coupon, helper.Error)
	DeleteCoupon(couponCode string) helper.Error
	AssignCoupon(rider *model.User, couponCode string) helper.Error
	// GetDiscount(rider *model.User) int
}
