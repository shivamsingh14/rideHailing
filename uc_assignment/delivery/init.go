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
	}

	DeliveryParams struct {
		UserUsecase     usecase.UserUsecase
		BookingsUsecase usecase.BookingUsecase
		CabuseCase      usecase.CabUsecase
	}
)

func NewHTTPDelivery(params DeliveryParams) {

	g = delivery{
		userUsecase:    params.UserUsecase,
		bookingUsecase: params.BookingsUsecase,
		cabUsecase:     params.CabuseCase,
	}

	InitEndpoint()
}
