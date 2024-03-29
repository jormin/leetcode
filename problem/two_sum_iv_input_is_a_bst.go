package problem

// 给定一个二叉搜索树 root 和一个目标结果 k，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。
//
// 示例 1：
// 输入: root = [5,3,6,2,4,null,7], k = 9
// 输出: true
//
// 示例 2：
// 输入: root = [5,3,6,2,4,null,7], k = 28
// 输出: false
//
// 示例 3：
// 输入: root = [2,1,3], k = 4
// 输出: true
//
// 示例 4：
// 输入: root = [2,1,3], k = 1
// 输出: false
//
// 示例 5：
// 输入: root = [2,1,3], k = 3
// 输出: true
//
// 提示:
// 二叉树的节点个数的范围是[1, 104].
// -104<= Node.val <= 104
// root为二叉搜索树
// -105<= k <= 105
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// findTarget 两数之和 IV - 输入 BST
func findTarget(root *TreeNode, k int) bool {
	var stack []*TreeNode
	m := map[int]bool{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		v := k - root.Val
		if _, ok := m[v]; ok {
			return true
		}
		m[root.Val] = true
		root = root.Right
	}
	return false
}
