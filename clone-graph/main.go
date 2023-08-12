/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    
}

// dfs solution

func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    
    record := make(map[*Node]*Node)
    
    dfs(node, record)
    
    return record[node]
}

func dfs(node *Node, record map[*Node]*Node) {
    if node == nil {
        return 
    }
    
    newNode := new(Node)
    newNode.Val = node.Val
    
    record[node] = newNode
    
    for _, child := range node.Neighbors {
        if _, exist := record[child]; !exist {
            dfs(child, record)
        }
        
        newNode.Neighbors = append(newNode.Neighbors, record[child])
    }
}



// Cache solution

func cloneGraph(node *Node) *Node {
    cache := map[*Node]*Node{}
    return cloneGraph2(node, cache)
}

func cloneGraph2(node *Node, cache map[*Node]*Node) *Node {
    if node == nil {
        return nil
    }
    if n, ok := cache[node]; ok {
        return n
    }
    clone := &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
    cache[node] = clone
    for i := 0; i < len(node.Neighbors); i++ {
        n := cloneGraph2(node.Neighbors[i], cache)
        if n != nil {
            clone.Neighbors = append(clone.Neighbors, n)
        }
    }
    return clone

}

