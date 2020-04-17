/*
 * Diameter of Binary Tree
 *
 * Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.
 *
 * Example:
 * Given a binary tree
 *
 * 1
 * / \
 * 2   3
 * / \
 * 4   5
 *
 * Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].
 *
 * Note: The length of path between two nodes is represented by the number of edges between them.
 */

package challenge

var cache = make(map[*TreeNode]int)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		curNode *TreeNode
		bestLen int
	)

	q := []*TreeNode{root}
	for len(q) > 0 {
		curNode = q[0]
		q = q[1:]
		candidateLen := maxDepth(curNode.Left) + maxDepth(curNode.Right)
		bestLen = max(bestLen, candidateLen)

		// let's try keeping append operations to a minimum (instead of just having the last two if statements)
		if curNode.Left != nil && curNode.Right != nil {
			q = append(q, curNode.Left, curNode.Right)
		} else if curNode.Left != nil {
			q = append(q, curNode.Left)
		} else if curNode.Right != nil {
			q = append(q, curNode.Right)
		}
	}
	return bestLen
}

func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	if value, ok := cache[node]; ok {
		return value
	}
	if node.Left == nil && node.Right == nil {
		cache[node] = 1
		return 1
	}

	curDepth := max(maxDepth(node.Left), maxDepth(node.Right)) + 1
	cache[node] = curDepth
	return curDepth
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
