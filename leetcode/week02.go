package leetcode

// FindMedianSortedArrays
// 4. 寻找两个正序数组的中位数
// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
// 算法的时间复杂度应该为 O(log (m+n))
// 示例 1：
// 输入：nums1 = [1,3], nums2 = [2]
// 输出：2.00000
// 解释：合并数组 = [1,2,3] ，中位数 2
// 示例 2：
// 输入：nums1 = [1,2], nums2 = [3,4]
// 输出：2.50000
// 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if (len(nums1)+len(nums2))%2 == 0 {
		//(len(nums1) + len(nums2) - 1) / 2 = midKey
	}
	return 0
}

// LongestPalindrome 5. 最长回文子串
// 中等
// 给你一个字符串 s，找到 s 中最长的回文子串。
//
// 示例 1：
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
// 示例 2：
// 输入：s = "cbbd"
// 输出："bb"
// 提示：
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成
func LongestPalindrome(s string) string {
	//12345678987654321
	//计算总任务数， 字符长度，由于只要查询出一个就行，所以这边就字符长度减1
	taskNums := len(s)
	//存储最大回文串
	longestPalindrome := ""
	//判断是否是回文
	isPalindrome := func(tempS string) bool {
		for h := 0; h < len(tempS)/2; h++ {
			if tempS[h] != tempS[len(tempS)-1-h] {
				return false
			}
		}
		return true
	}
	for i := 0; i < taskNums; i++ {
		//判断当前是否满足回文串
		for j := i + 1; j <= taskNums; j++ {
			tempS := s[i:j]
			if isPalindrome(tempS) {
				if len(tempS) > len(longestPalindrome) {
					longestPalindrome = tempS
				}
			}
		}
	}
	return longestPalindrome
}
