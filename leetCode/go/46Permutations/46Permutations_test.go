package permutations

import (
	"reflect"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "one",
			args: args{nums: []int{1}},
			want: [][]int{
				[]int{1},
			},
		},
		{
			name: "two",
			args: args{nums: []int{1, 2}},
			want: [][]int{
				[]int{1, 2},
				[]int{2, 1},
			},
		},
		{
			name: "three",
			args: args{nums: []int{1, 2, 3}},
			want: [][]int{
				[]int{1, 2, 3},
				[]int{1, 3, 2},
				[]int{2, 1, 3},
				[]int{2, 3, 1},
				[]int{3, 1, 2},
				[]int{3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.args.nums)
			tmp := make([][]int, len(tt.want))
			for i, t := range tt.want {
				tmp[i] = make([]int, len(t))
				copy(tmp[i], t)
			}
			for _, g := range got {
				found := false
				for i, w := range tmp {
					if reflect.DeepEqual(g, w) {
						found = true
						tmp = append(tmp[:i], tmp[i+1:]...)
						break
					}
				}
				if !found {
					t.Errorf("permute() = %v, lack %v", g, tmp)
					break
				}
			}
		})
	}
}
