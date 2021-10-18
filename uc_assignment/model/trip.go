package model

const (
	REST        = "Not Started"
	IN_PROGRESS = "Ongoing"
	FINISHED    = "Completed"
)

type Trip struct {
	ID          int
	Source      *Location
	Destination *Location
	Cab         *Cab
	Rider       *User
	Driver      *User
	Cost        float64
	Status      string
}
