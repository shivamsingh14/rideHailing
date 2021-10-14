package user

import (
	"errors"
	"uc_assignment/model"
)

type Driver struct {
}

func NewDriverRepo() UserRepository {
	return Driver{}
}

func (d Driver) Register(user model.User) (model.User, error) {

	_, found := userRecords[user.Id]

	if found {
		return user, errors.New("already existing driver")
	}
	lastID += 1
	userRecords[user.Id] = model.User{Id: lastID, Name: user.Name, Role: DRIVER}
	return user, nil

}
