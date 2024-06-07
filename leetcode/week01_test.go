package leetcode

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	// 定义测试用例
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{nums: []int{-3, 4, 3, 90}, target: 0, want: []int{0, 2}},
		{nums: []int{3, 3}, target: 6, want: []int{0, 1}},
		{nums: []int{-1, -2, -3, -4, -5}, target: -8, want: []int{2, 4}},
		// 可以根据需要添加更多测试用例
	}

	// 遍历测试用例
	for _, tt := range tests {
		//got := TwoSum(tt.nums, tt.target)
		got := TwoSum(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("TwoSum(%v, %v) = %v, want %v", tt.nums, tt.target, got, tt.want)
		}
	}
	// 遍历测试用例
	for _, tt := range tests {
		//got := TwoSum(tt.nums, tt.target)
		got := TwoSumV2(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("TwoSum(%v, %v) = %v, want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		//l1 = [2,4,3], l2 = [5,6,4]
		//[7,0,8]
		{"test1",
			args{
				&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
				&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			},
			&ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}},
		},
		//{"test2",
		//	args{
		//		&ListNode{Val: 0},
		//		&ListNode{Val: 0},
		//	},
		//	&ListNode{Val: 0},
		//},
		{"test3",
			args{
				&ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}},
				&ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}},
			},
			&ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got.Val, tt.want.Val) ||
				!reflect.DeepEqual(got.Next.Val, tt.want.Next.Val) ||
				!reflect.DeepEqual(got.Next.Next.Val, tt.want.Next.Next.Val) {
				t.Errorf("addTwoNumbers() = %v,%v,%v, want %v,%v,%v", got.Val, got.Next.Val, got.Next.Next.Val,
					tt.want.Val, tt.want.Next.Val, tt.want.Next.Next.Val)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbersV2(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got.Val, tt.want.Val) ||
				!reflect.DeepEqual(got.Next.Val, tt.want.Next.Val) ||
				!reflect.DeepEqual(got.Next.Next.Val, tt.want.Next.Next.Val) {
				t.Errorf("addTwoNumbers() = %v,%v,%v, want %v,%v,%v", got.Val, got.Next.Val, got.Next.Next.Val,
					tt.want.Val, tt.want.Next.Val, tt.want.Next.Next.Val)
			}
		})
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{"abcabcbb"}, 3},
		{"test2", args{"bbbbb"}, 1},
		{"test3", args{"pwwkew"}, 3},
		{"test4", args{" "}, 1},
		{"test5", args{"au"}, 2},
		{"test6", args{"dvdf"}, 3},
		{"test7", args{"abba"}, 2},
	}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := LengthOfLongestSubstring(tt.args.s); got != tt.want {
	//			t.Errorf("LengthOfLongestSubstring() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstringV2(tt.args.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
