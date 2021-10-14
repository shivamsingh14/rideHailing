package usecase

import (
	"uc_assignment/helper"
	"uc_assignment/model"
	repository "uc_assignment/repository/booking"
	cabRepository "uc_assignment/repository/cab"
	pricingRepository "uc_assignment/repository/pricing"
)

type BookingUsecase interface {
	BookRide(rider model.User, source model.Location, destination model.Location) (model.Trip, error)
	GetRide(tripId int) (model.Trip, error)
	RiderTripHistory(rider model.User) ([]model.Trip, error)
	DriverTripHistory(driver model.User) ([]model.Trip, error)
	CompleteTrip(trip model.Trip) (float64, error)
	calculateFare(distance float64, source, destination model.Location) float64
}

type BookingUsecaseParam struct {
	BookRide      repository.BookingRespository
	BasePricing   pricingRepository.PricingRepository
	PremimPricing pricingRepository.PricingRepository
	LongPricing   pricingRepository.PricingRepository
	CabRepo       cabRepository.CabRepository
}

type bookingUsecase struct {
	bookride       repository.BookingRespository
	basePricing    pricingRepository.PricingRepository
	premiumPricing pricingRepository.PricingRepository
	longPricing    pricingRepository.PricingRepository
	cabRepo        cabRepository.CabRepository
}

func NewBookingUsecase(param BookingUsecaseParam) BookingUsecase {

	return bookingUsecase{
		bookride:       param.BookRide,
		basePricing:    param.BasePricing,
		premiumPricing: param.PremimPricing,
		longPricing:    param.LongPricing,
		cabRepo:        param.CabRepo,
	}

}

func (booking bookingUsecase) BookRide(rider model.User, source model.Location, destination model.Location) (model.Trip, error) {

	availableCabs, _ := booking.cabRepo.GetAvailableCabs(helper.GetCurrentLocation(), helper.RADIUS)
	selectedCab := availableCabs[0]
	tripDistance := helper.FindDistance(source, destination)
	price := booking.calculateFare(tripDistance, source, destination)

	trip, err := booking.bookride.BookRide(source, destination, selectedCab, rider, price)

	if err != nil {
		return model.Trip{}, err
	}
	return trip, nil

}

func (booking bookingUsecase) calculateFare(distance float64, source, destination model.Location) float64 {

	var price float64

	if distance <= 2.0 {
		price = booking.basePricing.FindFare(source, destination)
	} else if distance > 2.0 && distance <= 6.0 {
		price = booking.basePricing.FindFare(source, destination)
	} else if distance > 6.0 {
		price = booking.basePricing.FindFare(source, destination)
	}

	return price
}

func (booking bookingUsecase) CompleteTrip(trip model.Trip) (float64, error) {

	trip, err := booking.bookride.GetRide(trip.ID)
	if err != nil {
		return 0.0, err
	}

	tripDistance := helper.FindDistance(trip.Source, trip.Destination)
	price := booking.calculateFare(tripDistance, trip.Source, trip.Destination)

	booking.cabRepo.CompleteTrip(trip)
	return price, nil

}

func (booking bookingUsecase) RiderTripHistory(rider model.User) ([]model.Trip, error) {

	var userTrips []model.Trip

	allTrips, err := booking.bookride.RideHistory()
	for _, trip := range allTrips {
		if trip.Rider.Id == rider.Id {
			userTrips = append(userTrips, trip)
		}
	}

	return userTrips, err

}

func (booking bookingUsecase) DriverTripHistory(driver model.User) ([]model.Trip, error) {

	var driverTrips []model.Trip

	allTrips, err := booking.bookride.RideHistory()
	for _, trip := range allTrips {
		if trip.Cab.Driver.Id == driver.Id {
			driverTrips = append(driverTrips, trip)
		}
	}

	return driverTrips, err

}

func (booking bookingUsecase) GetRide(tripId int) (model.Trip, error) {
	trip, err := booking.bookride.GetRide(tripId)
	if err != nil {
		return model.Trip{}, err
	}
	return trip, nil
}
