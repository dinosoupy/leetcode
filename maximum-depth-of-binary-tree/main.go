func maxDepth(root *TreeNode) int {
    if root==nil {
        return 0
    } 
    l := maxDepth(root.Left)
    r := maxDepth(root.Right)
    if l>r {
        return 1+l
    } else {
        return 1+r
    }
}