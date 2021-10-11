package problem

import (
	"math"
	"sort"
)

// 给定一个包括n 个整数的数组nums和 一个目标值target。找出nums中的三个整数，使得它们的和与target最接近。返回这三个数的和。假定每组输入只存在唯一答案。
//
// 示例：
// 输入：nums = [-1,2,1,-4], target = 1
// 输出：2
// 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//
// 提示：
// 3 <= nums.length <= 10^3
// -10^3<= nums[i]<= 10^3
// -10^4<= target<= 10^4
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/3sum-closest
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// threeSumClosest 最接近的三数之和
func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	dis := int(math.Abs(float64(res - target)))
	for i, v := range nums {
		if i > 0 && nums[i-1] == v {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := v + nums[l] + nums[r]
			if sum == target {
				return sum
			}
			n := int(math.Abs(float64(sum - target)))
			if n < dis {
				dis = n
				res = sum
			}
			if sum > target {
				r -= 1
			} else {
				l += 1
			}
		}
	}
	return res
}
