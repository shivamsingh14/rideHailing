package model

type Cab struct {
	Id          int
	driver      User
	License     string
	IsAvailable bool
}
