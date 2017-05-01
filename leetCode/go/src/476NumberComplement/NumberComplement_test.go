package main

import "testing"

func Test_findComplement(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "normal", args: args{num: 2}, want: 1},
		{name: "normal", args: args{num: 5}, want: 2},
		{name: "border", args: args{num: 0x40000000}, want: 0x3FFFFFFF},
		{name: "border", args: args{num: 0x80000000}, want: 0x7FFFFFFF},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findComplement(tt.args.num); got != tt.want {
				t.Errorf("findComplement() = %v, want %v", got, tt.want)
			}
		})
	}
}
