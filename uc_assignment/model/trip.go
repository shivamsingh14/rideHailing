package model

const (
	REST        = 0
	IN_PROGRESS = 1
	FINISHED    = 2
)

type Trip struct {
	ID          int
	Source      Location
	Destination Location
	Cab         Cab
	Rider       User
	Cost        float64
	Status      int
}
