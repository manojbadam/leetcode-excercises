package solution

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	} else if root.Left == nil && root.Right == nil {
		return true
	} else {
		return isValidBSTHelper(root, -2147483648, 2147483648)
	}
}

func isValidBSTHelper(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	fmt.Println(root.Val, min, max)
	if root.Val > min && root.Val < max {
		result := true
		if root.Left != nil {
			result = result && isValidBSTHelper(root.Left, int(math.Min(float64(min), float64(root.Val))), int(math.Min(float64(max), float64(root.Val))))
		}
		if root.Right != nil {
			result = result && isValidBSTHelper(root.Right, int(math.Max(float64(min), float64(root.Val))), int(math.Max(float64(max), float64(root.Val))))
		}
		return result
	} else if (root.Val == -2147483648) || (root.Val == 2147483648) {
		return true
	} else {
		return false
	}
}
