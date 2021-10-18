package bookRide

import (
	"reflect"
	"testing"
	"uc_assignment/helper"
	"uc_assignment/model"
)

func Test_bookRide_BookRide(t *testing.T) {
	type args struct {
		source      *model.Location
		destination *model.Location
		cab         *model.Cab
		rider       *model.User
		driver      *model.User
		price       float64
	}
	expectedResponse := model.Trip{
		ID:          1,
		Source:      &model.Location{Latitude: 0, Longitude: 0},
		Destination: &model.Location{Latitude: 10, Longitude: 0},
		Cab:         &model.Cab{Id: 1, RegistrationNumber: "1", IsAvailable: true},
		Rider:       &model.User{Number: 1, Id: 1, Role: helper.RIDER},
		Driver:      &model.User{Number: 1, Id: 1, Role: helper.DRIVER},
		Cost:        100,
		Status:      "Ongoing",
	}
	tests := []struct {
		name  string
		b     *bookRide
		args  args
		want  *model.Trip
		want1 helper.Error
	}{
		{
			name: "Test Book Ride",
			b:    &bookRide{},
			args: args{
				source:      &model.Location{Latitude: 0, Longitude: 0},
				destination: &model.Location{Latitude: 10, Longitude: 0},
				cab:         &model.Cab{Id: 1, RegistrationNumber: "1", IsAvailable: true},
				rider:       &model.User{Number: 1, Id: 1, Role: helper.RIDER},
				driver:      &model.User{Number: 1, Id: 1, Role: helper.DRIVER},
				price:       100,
			},
			want:  &expectedResponse,
			want1: (helper.Error{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bookRide{}
			got, got1 := b.BookRide(tt.args.source, tt.args.destination, tt.args.cab, tt.args.rider, tt.args.driver, tt.args.price)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bookRide.BookRide() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("bookRide.BookRide() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
