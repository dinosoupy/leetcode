/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    // corner case
    if root == nil {
        return [][]int{}
    }

    var ans [][]int
    // array for storing children in a level
    var level []*TreeNode
    // queue for storeing nodes
    // order of dequeue is actually level order traversal
    q := []*TreeNode{root}

    for len(q) > 0 {
        level = q
        q = []*TreeNode{}
        // append empty array for current level
        ans = append(ans, []int{})

        // level is an an array of nodes 
        // go through each node and add its value
        for _, curr := range level {
            ans[len(ans)-1] = append(ans[len(ans)-1], curr.Val)

            // if node has children, queue them
            if curr.Left != nil {
                q = append(q, curr.Left)
            }
            if curr.Right != nil {
                q = append(q, curr.Right)
            }
        } 
    }

    return ans
}