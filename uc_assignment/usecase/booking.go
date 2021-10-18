package usecase

import (
	"net/http"
	"uc_assignment/helper"
	"uc_assignment/model"
	repository "uc_assignment/repository/booking"
	cabRepository "uc_assignment/repository/cab"
	couponRepo "uc_assignment/repository/coupon"
	pricingRepository "uc_assignment/repository/pricing"
	userRepo "uc_assignment/repository/user"
)

type BookingUsecase interface {
	BookRide(rider model.User, source model.Location, destination model.Location) (*model.Trip, helper.Error)
	GetRide(tripId int) (*model.Trip, helper.Error)
	CompleteTrip(trip *model.Trip) (float64, helper.Error)
	calculateFare(distance float64, source, destination model.Location) float64
}

type BookingUsecaseParam struct {
	BookRide         repository.BookingRespository
	BasePricing      pricingRepository.PricingRepository
	PremimPricing    pricingRepository.PricingRepository
	LongPricing      pricingRepository.PricingRepository
	CabRepo          cabRepository.CabRepository
	DriverRepository userRepo.UserRepository
	RiderRepository  userRepo.UserRepository
	CouponRepo       couponRepo.CouponRepository
}

type bookingUsecase struct {
	bookride         repository.BookingRespository
	basePricing      pricingRepository.PricingRepository
	premiumPricing   pricingRepository.PricingRepository
	longPricing      pricingRepository.PricingRepository
	cabRepo          cabRepository.CabRepository
	driverRepository userRepo.UserRepository
	riderRepository  userRepo.UserRepository
	couponRepo       couponRepo.CouponRepository
}

func NewBookingUsecase(param BookingUsecaseParam) BookingUsecase {

	return &bookingUsecase{
		bookride:         param.BookRide,
		basePricing:      param.BasePricing,
		premiumPricing:   param.PremimPricing,
		longPricing:      param.LongPricing,
		cabRepo:          param.CabRepo,
		driverRepository: param.DriverRepository,
		riderRepository:  param.RiderRepository,
		couponRepo:       param.CouponRepo,
	}

}

func (booking *bookingUsecase) BookRide(rider model.User, source model.Location, destination model.Location) (*model.Trip, helper.Error) {

	if source == (model.Location{}) {
		source = helper.GetCurrentLocation()
	}
	getRider, getRiderErr := booking.riderRepository.GetUser(rider.Number)
	if getRiderErr != (helper.Error{}) {
		return nil, helper.Error{Code: getRiderErr.Code, Message: "Please register before booking"}
	}
	if booking.riderRepository.CheckTripStatus(rider) {
		return nil, helper.Error{Code: http.StatusForbidden, Message: "Rider's trip is already in progress"}
	}
	availableCabs, availableCabsErr := booking.cabRepo.GetAvailableCabs(&source, helper.RADIUS)
	if availableCabsErr != (helper.Error{}) {
		return nil, availableCabsErr
	}
	if len(availableCabs) == 0 {
		return nil, helper.Error{Code: http.StatusNotFound, Message: "No cabs available"}
	}
	selectedCab := availableCabs[0]
	getDriver := booking.cabRepo.GetDriver(selectedCab)
	booking.cabRepo.MarkCabAvailability(selectedCab, false)
	tripDistance := helper.FindDistance(source, destination)
	price := booking.calculateFare(tripDistance, source, destination)

	trip, err := booking.bookride.BookRide(&source, &destination, selectedCab, getRider, getDriver, price)
	booking.driverRepository.AddTrip(getDriver, trip)
	booking.riderRepository.AddTrip(getRider, trip)

	if err != (helper.Error{}) {
		return nil, err
	}
	return trip, helper.Error{}

}

func (booking *bookingUsecase) calculateFare(distance float64, source, destination model.Location) float64 {

	var price float64

	if distance <= 2.0 {
		price = booking.basePricing.FindFare(source, destination)
	} else if distance > 2.0 && distance <= 6.0 {
		price = booking.premiumPricing.FindFare(source, destination)
	} else if distance > 6.0 {
		price = booking.longPricing.FindFare(source, destination)
	}

	return price
}

func (booking *bookingUsecase) CompleteTrip(trip *model.Trip) (float64, helper.Error) {

	trip, err := booking.bookride.GetRide(trip.ID)
	if err != (helper.Error{}) {
		return 0.0, err
	}

	tripDistance := helper.FindDistance(*trip.Source, *trip.Destination)
	price := booking.calculateFare(tripDistance, *trip.Source, *trip.Destination)

	completeTripError := booking.cabRepo.CompleteTrip(trip)
	if completeTripError != (helper.Error{}) {
		return 0.0, completeTripError
	}
	booking.driverRepository.CompleteTrip(trip.Driver, trip)
	booking.riderRepository.CompleteTrip(trip.Rider, trip)
	return price, helper.Error{}

}

func (booking *bookingUsecase) GetRide(tripId int) (*model.Trip, helper.Error) {
	trip, err := booking.bookride.GetRide(tripId)
	if err != (helper.Error{}) {
		return nil, err
	}
	return trip, helper.Error{}
}
