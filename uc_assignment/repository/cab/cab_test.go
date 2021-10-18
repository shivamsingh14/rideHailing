package cab

import (
	"testing"
	"uc_assignment/model"
)

func TestCab_UpdateLocation(t *testing.T) {
	type args struct {
		cab      *model.Cab
		location *model.Location
	}
	cab := model.Cab{
		Id:                 1,
		RegistrationNumber: "1",
		IsAvailable:        true,
	}
	tests := []struct {
		name string
		c    *Cab
		args args
	}{
		{
			name: "Test update Location",
			c:    &Cab{},
			args: args{cab: &cab, location: &model.Location{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cab{}
			c.UpdateLocation(tt.args.cab, tt.args.location)
		})
	}
}
