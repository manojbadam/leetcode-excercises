package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		return DepthHelper(root, 1)
	}
}

func DepthHelper(root *TreeNode, pos int) int {
	if root == nil {
		return pos
	} else {
		left := pos
		right := pos
		if root.Left != nil {
			left = DepthHelper(root.Left, pos+1)
		}
		if root.Right != nil {
			right = DepthHelper(root.Right, pos+1)
		}
		if left < right {
			return right
		} else {
			return left
		}
	}
}
