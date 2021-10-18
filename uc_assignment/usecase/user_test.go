package usecase

import (
	"reflect"
	"testing"
	"uc_assignment/helper"
	"uc_assignment/model"
	bookingRepo "uc_assignment/repository/booking"
	userRepo "uc_assignment/repository/user"
)

func Test_userUsecase_RegisterUser(t *testing.T) {
	type fields struct {
		driverRepository  userRepo.UserRepository
		riderRepository   userRepo.UserRepository
		bookingRepository bookingRepo.BookingRespository
	}
	driverRepository := userRepo.NewDriverRepo()
	riderRepository := userRepo.NewRiderRepo()
	bookingRepository := bookingRepo.NewBookRideRepo()
	userParam := model.User{
		Number: 1,
		Name:   "Shivam",
	}
	expectedResponse := model.User{
		Id:     1,
		Number: 1,
		Name:   "Shivam",
		Role:   helper.DRIVER,
	}
	type args struct {
		userParam model.User
		role      string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *model.User
		want1  helper.Error
	}{
		{
			name:   "Test Register User",
			fields: fields{driverRepository: driverRepository, riderRepository: riderRepository, bookingRepository: bookingRepository},
			args:   args{userParam: userParam, role: helper.DRIVER},
			want:   &expectedResponse,
			want1:  (helper.Error{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := userUsecase{
				driverRepository:  tt.fields.driverRepository,
				riderRepository:   tt.fields.riderRepository,
				bookingRepository: tt.fields.bookingRepository,
			}
			got, got1 := u.RegisterUser(tt.args.userParam, tt.args.role)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.RegisterUser() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("userUsecase.RegisterUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
