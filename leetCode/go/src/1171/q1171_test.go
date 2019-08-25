package q1171

import (
	"reflect"
	"testing"
)

func Test_removeZeroSumSublists(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "-1 1 0 1",
			args: args{head: &ListNode{Val: -1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}}}},
			want: &ListNode{Val: 1, Next: nil},
		},
		{
			name: "zero",
			args: args{head: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: nil}}},
			want: nil,
		},
		{
			name: "sample",
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: -3, Next: &ListNode{Val: 4, Next: nil}}}}}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeZeroSumSublists(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeZeroSumSublists() = %v, want %v", got, tt.want)
			}
		})
	}
}
