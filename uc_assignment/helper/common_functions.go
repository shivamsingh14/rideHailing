package helper

import (
	"math"
	"math/rand"
	"uc_assignment/model"
)

func FindDistance(source, destination model.Location) float64 {

	first := math.Pow(float64(destination.Latitude-source.Latitude), 2)
	second := math.Pow(float64(destination.Longitude-source.Longitude), 2)
	return math.Sqrt(first + second)

}

func GetCurrentLocation() model.Location {

	return model.Location{Latitude: rand.Intn(RANDOM_UPPER_LIMIT), Longitude: rand.Intn(RANDOM_UPPER_LIMIT)}
}
