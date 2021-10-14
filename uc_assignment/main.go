package main

import (
	"uc_assignment/delivery"
	bookingRepository "uc_assignment/repository/booking"
	cabRepository "uc_assignment/repository/cab"
	pricingRepository "uc_assignment/repository/pricing"
	repository "uc_assignment/repository/user"
	"uc_assignment/usecase"
)

var deliveryParams delivery.DeliveryParams

func main() {

	userUsecase := initializeUserUsecase()
	bookingsUsecase := initializeBookingUsecase()

	deliveryParams = delivery.DeliveryParams{
		UserUsecase:     userUsecase,
		BookingsUsecase: bookingsUsecase,
	}

	delivery.NewHTTPDelivery(deliveryParams)

}

func initializeUserUsecase() usecase.UserUsecase {

	driverRepo := repository.NewDriverRepo()
	riderRepo := repository.NewRiderRepo()

	userUsecaseParams := usecase.UserUsecaseParam{
		DriverRepository: driverRepo,
		RiderRepository:  riderRepo,
	}

	return usecase.NewUserUsecase(userUsecaseParams)
}

func initializeBookingUsecase() usecase.BookingUsecase {

	basePricingRepo := pricingRepository.NewBaseRideRepo()
	premiumPricingRepo := pricingRepository.NewPremiumRideRepo()
	longPricingRepo := pricingRepository.NewLongRideRepo()
	bookRideRepo := bookingRepository.NewBookRideRepo()
	cabRepo := cabRepository.NewCabRepository()

	bookingUseCaseParams := usecase.BookingUsecaseParam{
		BookRide:      bookRideRepo,
		BasePricing:   basePricingRepo,
		PremimPricing: premiumPricingRepo,
		LongPricing:   longPricingRepo,
		CabRepo:       cabRepo,
	}

	return usecase.NewBookingUsecase(bookingUseCaseParams)

}

func initializeCabUsecase() usecase.CabUsecase {

	cabUseCaseParams := usecase.CabUsecaseParam{}

	return usecase.NewCabUsecase(cabUseCaseParams)

}
