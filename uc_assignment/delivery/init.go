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
		cabuseCase     usecase.CabUsecase
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
		cabuseCase:     params.CabuseCase,
	}

	InitEndpoint()
}
