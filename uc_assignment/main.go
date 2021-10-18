package main

import (
	"uc_assignment/delivery"
	bookingRepository "uc_assignment/repository/booking"
	cabRepository "uc_assignment/repository/cab"
	couponRepository "uc_assignment/repository/coupon"
	pricingRepository "uc_assignment/repository/pricing"
	repository "uc_assignment/repository/user"
	"uc_assignment/usecase"
)

var deliveryParams delivery.DeliveryParams

func main() {

	userUsecase := initializeUserUsecase()
	bookingsUsecase := initializeBookingUsecase()
	cabUsecase := initializeCabUsecase()
	couponUsecase := initializeCouponUsecase()

	deliveryParams = delivery.DeliveryParams{
		UserUsecase:     userUsecase,
		BookingsUsecase: bookingsUsecase,
		CabuseCase:      cabUsecase,
		CouponUsecase:   couponUsecase,
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
	driverRepo := repository.NewDriverRepo()
	riderRepo := repository.NewRiderRepo()
	flatDiscountRepo := couponRepository.NewFlatDiscountRepo()

	bookingUseCaseParams := usecase.BookingUsecaseParam{
		BookRide:         bookRideRepo,
		BasePricing:      basePricingRepo,
		PremimPricing:    premiumPricingRepo,
		LongPricing:      longPricingRepo,
		CabRepo:          cabRepo,
		DriverRepository: driverRepo,
		RiderRepository:  riderRepo,
		CouponRepo:       flatDiscountRepo,
	}

	return usecase.NewBookingUsecase(bookingUseCaseParams)

}

func initializeCabUsecase() usecase.CabUsecase {

	cabRepository := cabRepository.NewCabRepository()
	driverRepo := repository.NewDriverRepo()

	cabUseCaseParams := usecase.CabUsecaseParam{
		CabRepository:    cabRepository,
		DriverRepository: driverRepo,
	}

	return usecase.NewCabUsecase(cabUseCaseParams)

}

func initializeCouponUsecase() usecase.CouponUsecase {

	flatDiscountRepo := couponRepository.NewFlatDiscountRepo()
	riderRepo := repository.NewRiderRepo()

	couponUseCaseParams := usecase.CouponUsecaseParam{
		CouponRepo: flatDiscountRepo,
		RiderRepo:  riderRepo,
	}

	return usecase.NewCouponUsecase(couponUseCaseParams)
}
