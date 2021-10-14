package user

import "uc_assignment/model"

var userRecords = make(map[int]model.User)
var lastID int

const (
	DRIVER = "driver"
	RIDER  = "rider"
)

type UserRepository interface {
	Register(user model.User) (model.User, error)
}
