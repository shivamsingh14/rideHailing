package pricing

import (
	"uc_assignment/helper"
	"uc_assignment/model"
)

var premium_per_km_price float64 = 8

type PremiumRide struct {
	PER_KM_PRICE float64
}

func NewPremiumRideRepo() PricingRepository {
	return BaseRide{
		PER_KM_PRICE: premium_per_km_price,
	}
}

func (p PremiumRide) FindFare(source model.Location, destination model.Location) float64 {

	return helper.FindDistance(source, destination)*p.PER_KM_PRICE + MIN_FARE

}
