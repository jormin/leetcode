package problem

import (
	"sort"
)

// 给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
// 示例 1：
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
//
// 示例 2：
// 输入：nums = []
// 输出：[]
//
// 示例 3：
// 输入：nums = [0]
// 输出：[]
//
// 提示：
// 0 <= nums.length <= 3000
// -105 <= nums[i] <= 105
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/3sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// threeSum 三数之和
func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}
	// 排序
	sort.Ints(nums)
	for i, v := range nums {
		// 如果元素大于0，则后面的元素都大于0，和不可能为0
		if v > 0 {
			return res
		}
		// 去重重复元素
		if i > 0 && v == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := v + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{v, nums[l], nums[r]})
				// 去掉左侧和右侧的重复元素
				for l < r && nums[l] == nums[l+1] {
					l += 1
				}
				for l < r && nums[r] == nums[r-1] {
					r -= 1
				}
				l += 1
				r -= 1
			} else if sum > 0 {
				// 和大于0，说明右侧元素过大
				r -= 1
			} else {
				// 和小于0，说明左侧元素过小
				l += 1
			}
		}
	}
	return res
}
