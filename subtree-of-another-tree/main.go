func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    // return true if both are nil, false if one of them isn't nil
    if root==nil || subRoot==nil {
        return root==subRoot
    }
    // check if subree and main tree are already equal
    if isSameTree(root, subRoot) {
        return true
    }

    // check if subtree is equal to main trees left subtree or right
    return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
    // Return true if all nodes of two trees are same
    if p==nil || q==nil {
        return p==q
    }
    return p.Val==q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}