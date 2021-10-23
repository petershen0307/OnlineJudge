package solution

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{N: 268},
			want: 5268,
		}, {
			name: "2",
			args: args{N: -999},
			want: -5999,
		},
		{
			name: "3",
			args: args{N: 670},
			want: 6750,
		},
		{
			name: "4",
			args: args{N: 0},
			want: 50,
		},
		{
			name: "5",
			args: args{N: -1670},
			want: -15670,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.N); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
