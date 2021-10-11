package problem

import (
	"sort"
)

// 给你一个由 n 个整数组成的数组nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组[nums[a], nums[b], nums[c], nums[d]] ：
//
// 0 <= a, b, c, d< n
// a、b、c 和 d 互不相同
// nums[a] + nums[b] + nums[c] + nums[d] == target
// 你可以按 任意顺序 返回答案 。
//
// 示例 1：
// 输入：nums = [1,0,-1,0,-2,2], target = 0
// 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
//
// 示例 2：
// 输入：nums = [2,2,2,2,2], target = 8
// 输出：[[2,2,2,2]]
//
// 提示：
// 1 <= nums.length <= 200
// -109 <= nums[i] <= 109
// -109 <= target <= 109
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/4sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// fourSum 四数之和
func fourSum(nums []int, target int) [][]int {
	var res [][]int
	if len(nums) < 4 {
		return res
	}
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i-1] == nums[i] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j-1] == nums[j] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for l, r := j+1, n-1; l < r; {
				if sum := nums[i] + nums[j] + nums[l] + nums[r]; sum < target {
					l += 1
				} else if sum > target {
					r -= 1
				} else {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					// 去掉左侧和右侧的重复元素
					for l++; l < r && nums[l] == nums[l-1]; l++ {
					}
					for r--; l < r && nums[r] == nums[r+1]; r-- {
					}
				}
			}
		}
	}
	return res
}
