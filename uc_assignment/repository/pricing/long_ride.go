package pricing

import (
	"uc_assignment/helper"
	"uc_assignment/model"
)

var long_per_km_price float64 = 5

type LongRide struct {
	PER_KM_PRICE float64
}

func NewLongRideRepo() PricingRepository {
	return BaseRide{
		PER_KM_PRICE: long_per_km_price,
	}
}

func (l LongRide) FindFare(source model.Location, destination model.Location) float64 {

	return helper.FindDistance(source, destination)*l.PER_KM_PRICE + MIN_FARE

}
