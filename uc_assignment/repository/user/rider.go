package user

import (
	"errors"
	"uc_assignment/model"
)

type Rider struct {
}

func NewRiderRepo() UserRepository {
	return Rider{}
}

func (r Rider) Register(userParam model.User) (model.User, error) {

	_, found := userRecords[userParam.Id]

	if found {
		return model.User{}, errors.New("user already exists")
	}

	lastID += 1
	newUser := model.User{Id: lastID, Name: userParam.Name, Role: RIDER}
	userRecords[userParam.Id] = newUser
	return newUser, nil
}
