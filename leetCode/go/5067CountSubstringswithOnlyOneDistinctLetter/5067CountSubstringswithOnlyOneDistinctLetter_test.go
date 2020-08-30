package countsubstringswithonlyonedistinctletter

import "testing"

func Test_countLetters(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "1", args: args{S: "aaa"}, want: 6},
		{name: "2", args: args{S: "aaaba"}, want: 8},
		{name: "3", args: args{S: "aaabab"}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countLetters(tt.args.S); got != tt.want {
				t.Errorf("countLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
