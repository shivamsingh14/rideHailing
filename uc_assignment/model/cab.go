package model

type Cab struct {
	Id          int
	Driver      User
	License     string
	IsAvailable bool
}
