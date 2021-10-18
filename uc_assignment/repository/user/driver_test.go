package user

import (
	"reflect"
	"testing"
	"uc_assignment/helper"
	requestresponse "uc_assignment/helper/requestResponse"
	"uc_assignment/model"
)

func TestDriver_TripHistory(t *testing.T) {
	type fields struct {
		role string
	}
	type args struct {
		phoneNumber int
	}
	ongoingTrip := model.Trip{
		ID:          1,
		Source:      &model.Location{Latitude: 0, Longitude: 0},
		Destination: &model.Location{Latitude: 10, Longitude: 0},
		Cab:         &model.Cab{Id: 1, RegistrationNumber: "1", IsAvailable: true},
		Rider:       &model.User{Id: 1, Number: 12345},
		Driver:      &model.User{Number: 1, Id: 1, Name: "1", Role: DRIVER},
		Cost:        100,
		Status:      "Ongoing",
	}
	expectedResponse := requestresponse.TripHistory{
		Ongoing:   []*model.Trip{&ongoingTrip},
		Completed: []*model.Trip{},
	}
	driverOngoingTrips[12345] = []*model.Trip{&ongoingTrip}
	driverCompletedTrips[12345] = []*model.Trip{}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   requestresponse.TripHistory
		want1  helper.Error
	}{
		{
			name:  "Test Trip History",
			args:  args{phoneNumber: 12345},
			want:  expectedResponse,
			want1: (helper.Error{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Driver{
				role: tt.fields.role,
			}
			got, got1 := d.TripHistory(tt.args.phoneNumber)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Driver.TripHistory() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Driver.TripHistory() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDriver_Register(t *testing.T) {
	type fields struct {
		role string
	}
	type args struct {
		userParam model.User
	}
	x := model.User{
		Number: 1,
		Name:   "Shivam",
	}
	u := &model.User{
		Id:     1,
		Number: 1,
		Name:   "Shivam",
		Role:   "driver",
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
			fields: fields{role: DRIVER},
			args:   args{userParam: x},
			want:   u,
			want1:  (helper.Error{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Driver{
				role: tt.fields.role,
			}
			got, got1 := d.Register(tt.args.userParam)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Driver.Register() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Driver.Register() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDriver_CompleteTrip(t *testing.T) {
	type fields struct {
		role string
	}
	driver := model.User{
		Id:     1,
		Name:   "Driver 1",
		Number: 1,
		Role:   DRIVER,
	}

	type args struct {
		driver *model.User
		trip   *model.Trip
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Test complete trip",
			fields: fields{role: DRIVER},
			args:   args{driver: &driver},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Driver{
				role: tt.fields.role,
			}
			d.CompleteTrip(tt.args.driver, tt.args.trip)
		})
	}
}
