package delivery

import (
	"uc_assignment/usecase"

	"github.com/julienschmidt/httprouter"
)

var router *httprouter.Router
var g delivery

type (
	delivery struct {
		userUsecase    usecase.UserUsecase
		bookingUsecase usecase.BookingUsecase
		cabUsecase     usecase.CabUsecase
		couponUsecase  usecase.CouponUsecase
	}

	DeliveryParams struct {
		UserUsecase     usecase.UserUsecase
		BookingsUsecase usecase.BookingUsecase
		CabuseCase      usecase.CabUsecase
		CouponUsecase   usecase.CouponUsecase
	}
)

func NewHTTPDelivery(params DeliveryParams) {

	g = delivery{
		userUsecase:    params.UserUsecase,
		bookingUsecase: params.BookingsUsecase,
		cabUsecase:     params.CabuseCase,
		couponUsecase:  params.CouponUsecase,
	}

	InitEndpoint()
}
