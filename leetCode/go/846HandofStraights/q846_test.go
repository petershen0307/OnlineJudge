package q846

import "testing"

func Test_isNStraightHand(t *testing.T) {
	type args struct {
		hand []int
		W    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "group 3",
			args: args{
				hand: []int{8, 11, 12},
				W:    3,
			},
			want: false,
		},
		{
			name: "group 2",
			args: args{
				hand: []int{1, 2, 3, 4, 5, 6},
				W:    2,
			},
			want: true,
		},
		{
			name: "group 0",
			args: args{
				hand: []int{1, 2, 3, 4, 5},
				W:    4,
			},
			want: false,
		},
		{
			name: "group 3",
			args: args{
				hand: []int{1, 2, 3, 6, 2, 3, 4, 7, 8},
				W:    3,
			},
			want: true,
		},
		{
			name: "give 2 value and only 1 consecutive",
			args: args{
				hand: []int{1, 2},
				W:    1,
			},
			want: true,
		},
		{
			name: "give 1 value and only 1 consecutive",
			args: args{
				hand: []int{1},
				W:    1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNStraightHand(tt.args.hand, tt.args.W); got != tt.want {
				t.Errorf("isNStraightHand() = %v, want %v", got, tt.want)
			}
		})
	}
}
