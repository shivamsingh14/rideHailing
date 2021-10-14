package pricing

import (
	"uc_assignment/helper"
	"uc_assignment/model"
)

var base_per_km_price float64 = 10

type BaseRide struct {
	PER_KM_PRICE float64
}

func NewBaseRideRepo() PricingRepository {
	return BaseRide{
		PER_KM_PRICE: base_per_km_price,
	}
}

func (b BaseRide) FindFare(source model.Location, destination model.Location) float64 {

	return helper.FindDistance(source, destination)*b.PER_KM_PRICE + MIN_FARE

}
