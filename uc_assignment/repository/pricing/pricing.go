package pricing

import "uc_assignment/model"

var MIN_FARE float64 = 50

type PricingRepository interface {
	FindFare(source, destination model.Location) float64
}
