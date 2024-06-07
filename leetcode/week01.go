package leetcode

// TwoSum
// 1.两束之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。
// 示例 1：
// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
// 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
// 示例 2：
// 输入：nums = [3,2,4], target = 6
// 输出：[1,2]
// 示例 3：
// 输入：nums = [3,3], target = 6
// 输出：[0,1]
func TwoSum(nums []int, target int) []int {
	result := []int{}
	for k, v := range nums {
		for kk, vv := range nums[k+1:] {
			if v+vv == target {
				result = append(result, k, kk+k+1)
			}
		}
	}
	return result
}

// TwoSumV2 基于map的方式
func TwoSumV2(nums []int, target int) []int {
	result := []int{}
	numsMap := map[int]int{}
	for k, v := range nums {
		if i, ok := numsMap[target-v]; ok {
			result = append(result, i, k)
		}
		numsMap[v] = k
	}
	return result
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers
// 2. 两数相加
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
// https://leetcode.cn/problems/add-two-numbers/description/
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	if l1 == nil {
		l1 = &ListNode{Val: 0}
	}
	if l2 == nil {
		l2 = &ListNode{Val: 0}
	}
	if l1.Val+l2.Val >= 10 {
		// 计算当前的个位值
		res.Val = (l1.Val + l2.Val) % 10
		// 计算下一位的该加的值
		if l1.Next == nil {
			l1.Next = &ListNode{}
		}
		l1.Next.Val = (l1.Val+l2.Val)/10 + l1.Next.Val
	} else {
		res.Val = l1.Val + l2.Val
	}
	if l1.Next != nil || l2.Next != nil {
		res.Next = addTwoNumbers(l1.Next, l2.Next)
	}
	return res
}

func addTwoNumbersV2(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

// LengthOfLongestSubstring
// 3. 无重复字符的最长子串
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度
// 示例 1:
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//
//	请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
func LengthOfLongestSubstring(s string) int {
	uniqMap := map[string]int{}
	max := 0
	start := 0
	for i := 0; i < len(s); i++ {
		if _, ok := uniqMap[s[i:i+1]]; !ok {
			uniqMap[s[i:i+1]] = i
			if len(uniqMap) > max {
				max = len(uniqMap)
			}
		} else {
			uniqMap = map[string]int{}
			start = start + 1
			i = start
			uniqMap[s[i:i+1]] = i
		}
	}
	return max
}

func LengthOfLongestSubstringV2(s string) int {
	//用map来表示窗口
	uniqMap := map[byte]int{}
	//窗口的开始位置
	//右边的指针从-1 ,未移动的时候在窗口的左边界
	rk, maxLen := -1, 0
	//左边的指针移动，从0开始
	for lk := 0; lk < len(s); lk++ {
		if lk != 0 {
			//最左边的指针右移
			delete(uniqMap, s[lk-1])
		}
		//如果右边指针没超，且数据不再窗口里面
		for rk+1 < len(s) && uniqMap[s[rk+1]] == 0 {
			uniqMap[s[rk+1]]++
			rk++
		}
		//取最长的窗口宽度
		if maxLen < len(uniqMap) {
			maxLen = len(uniqMap)
		}
	}
	return maxLen
}
