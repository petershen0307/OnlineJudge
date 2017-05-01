package HammingDistance

import "testing"

func Test_hammingDistance(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "border", args: args{x: 0xF0000000, y: 4}, want: -1},
		{name: "border", args: args{x: 0, y: 0x80000000}, want: -1},
		{name: "border", args: args{x: -1, y: 0}, want: -1},
		{name: "normal", args: args{x: 1, y: 4}, want: 2},
		{name: "normal", args: args{x: 10, y: 5}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingDistance(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("hammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
