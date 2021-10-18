package user

import (
	"reflect"
	"testing"
	"uc_assignment/helper"
	"uc_assignment/model"
)

func TestRider_Register(t *testing.T) {
	type fields struct {
		role string
	}
	type args struct {
		userParam model.User
	}
	userParam := model.User{
		Number: 1,
		Name:   "Shivam",
	}
	expectedResponse := model.User{
		Id:     1,
		Number: 1,
		Name:   "Shivam",
		Role:   RIDER,
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *model.User
		want1  helper.Error
	}{
		{
			name:   "Test Register Rider",
			fields: fields{role: RIDER},
			args:   args{userParam: userParam},
			want:   &expectedResponse,
			want1:  (helper.Error{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rider{
				role: tt.fields.role,
			}
			got, got1 := r.Register(tt.args.userParam)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rider.Register() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Rider.Register() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
