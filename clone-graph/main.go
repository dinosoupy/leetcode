/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */


func cloneGraph(node *Node) *Node {
    // If graph is empty
   if node==nil{
       return node
   }
   visited:=make(map[*Node]*Node)
   dfs(node, visited)
   return visited[node]
}

func dfs(node *Node, visited map[*Node]*Node) {
    // Recursion stopping condition
    if node==nil {
        return
    }
    newNode:=&Node{Val: node.Val}
    visited[node]=newNode
    for _, neighbor:= range node.Neighbors {
        if _, ok:=visited[neighbor]; !ok {
            dfs(neighbor, visited)
        } 
        newNode.Neighbors = append(newNode.Neighbors, visited[neighbor])
    }
}

