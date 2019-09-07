package shortestdistancetotargetcolor

import (
	"reflect"
	"testing"
)

func Test_shortestDistanceColor(t *testing.T) {
	type args struct {
		colors  []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "3",
			args: args{
				colors:  []int{3, 2, 2, 1, 3, 1, 1, 1, 3, 1},
				queries: [][]int{[]int{4, 1}, []int{9, 2}, []int{4, 2}, []int{8, 1}, []int{0, 3}, []int{2, 1}, []int{2, 3}, []int{6, 3}, []int{4, 1}, []int{1, 2}},
			},
			want: []int{1, 7, 2, 1, 0, 1, 2, 2, 1, 0},
		},
		{
			name: "1",
			args: args{
				colors:  []int{1, 1, 2, 1, 3, 2, 2, 3, 3},
				queries: [][]int{[]int{1, 3}, []int{2, 2}, []int{6, 1}},
			},
			want: []int{3, 0, 3},
		},
		{
			name: "2",
			args: args{
				colors:  []int{1, 2},
				queries: [][]int{[]int{0, 3}},
			},
			want: []int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestDistanceColor(tt.args.colors, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestDistanceColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
