package requestresponse

import "uc_assignment/model"

type HttpResponseBody []byte

type CabLocation struct {
	Cab      *model.Cab
	Location *model.Location
	Driver   *model.User
}

type AssignCab struct {
	RegistrationNumber string
	DriverNumber       int
}

type BookRide struct {
	Source      model.Location
	Destination model.Location
	Rider       model.User
}

type AssignCoupon struct {
	CouponCode string
	UserNumber int
}
