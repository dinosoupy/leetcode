/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    // lowest possible int32
    mx := math.MinInt32

    var dfs func(root *TreeNode) int 
    dfs = func(root *TreeNode) int {
        if root==nil {
            return 0
        }
        // at each node calculate sum of left and righ subtree
        // discard sum if it is -ve 
        left := max(dfs(root.Left), 0)
        right := max(dfs(root.Right), 0)

        // update global max if necessary
        mx = max(mx, root.Val + left + right)

        // return self plus highest sum from subtrees
        return root.Val + max(left, right)
    }

    dfs(root)
    return mx
}

func max(x, y int) int {
    if x > y {
        return x
    } else {
        return y
    }
}