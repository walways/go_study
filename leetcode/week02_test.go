package leetcode

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {

}

func TestLongestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "test2",
			args: args{
				s: "babaab",
			},
			want: "baab",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("LongestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
