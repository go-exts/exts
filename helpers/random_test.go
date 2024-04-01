package helpers

import (
	"reflect"
	"testing"
)

func TestNewRandom(t *testing.T) {
	type args struct {
		rnd int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "RandomNumber",
			args: args{
				rnd: RandomNumber,
			},
			want: "0123456789",
		},
		{
			name: "RandomLower",
			args: args{
				rnd: RandomLower,
			},
			want: "abcdefghijklmnopqrstuvwxyz",
		},
		{
			name: "RandomUpper",
			args: args{
				rnd: RandomUpper,
			},
			want: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
		{
			name: "Lower&Upper",
			args: args{
				rnd: RandomLower | RandomUpper,
			},
			want: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
		{
			name: "RandomLowerNum",
			args: args{
				rnd: RandomLower | RandomNumber,
			},
			want: "0123456789abcdefghijklmnopqrstuvwxyz",
		},
		{
			name: "RandomUpperNum",
			args: args{
				rnd: RandomUpper | RandomNumber,
			},
			want: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
		{
			name: "RandomSymbol",
			args: args{
				rnd: RandomSymbol,
			},
			want: "!@#$%^&*-_=+",
		},
		{
			name: "RandomLowerUpperNum",
			args: args{
				rnd: RandomLower | RandomNumber | RandomUpper,
			},
			want: "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
		{
			name: "RandomAll",
			args: args{
				rnd: RandomAll,
			},
			want: "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*-_=+",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRand(tt.args.rnd)
			if !reflect.DeepEqual(got.charset, tt.want) {
				t.Errorf("NewRandom() = %v, want %v", got, tt.want)
			}
		})
	}
}
