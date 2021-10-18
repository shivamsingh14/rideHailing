package usecase

import (
	"uc_assignment/helper"
	requestresponse "uc_assignment/helper/requestResponse"
	"uc_assignment/model"
	couponRepo "uc_assignment/repository/coupon"
	userRepo "uc_assignment/repository/user"
)

type CouponUsecase interface {
	AddCoupon(couponParam model.Coupon) (*model.Coupon, helper.Error)
	DeleteCoupon(couponCode string) helper.Error
	AssignCoupon(assignCouponParam requestresponse.AssignCoupon) helper.Error
}

type CouponUsecaseParam struct {
	CouponRepo couponRepo.CouponRepository
	RiderRepo  userRepo.UserRepository
}

type couponUsecase struct {
	couponRepo couponRepo.CouponRepository
	riderRepo  userRepo.UserRepository
}

func NewCouponUsecase(couponParam CouponUsecaseParam) CouponUsecase {
	return &couponUsecase{
		couponRepo: couponParam.CouponRepo,
		riderRepo:  couponParam.RiderRepo,
	}
}

func (c *couponUsecase) AddCoupon(couponParam model.Coupon) (*model.Coupon, helper.Error) {
	return c.couponRepo.AddCoupon(couponParam)
}

func (c *couponUsecase) DeleteCoupon(couponCode string) helper.Error {
	return c.couponRepo.DeleteCoupon(couponCode)
}

func (c *couponUsecase) AssignCoupon(assignCouponParam requestresponse.AssignCoupon) helper.Error {

	getRider, getRiderError := c.riderRepo.GetUser(assignCouponParam.UserNumber)
	if getRiderError != (helper.Error{}) {
		return getRiderError
	}
	return c.couponRepo.AssignCoupon(getRider, assignCouponParam.CouponCode)
}
