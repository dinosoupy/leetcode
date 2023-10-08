func canFinish(numCourses int, prerequisites [][]int) bool {
    // construct graph where vertices are course numbers 
    // edges are prerequisites stored in adjacency list

    adjList:=make(map[int][]int)
    for _, course:= range prerequisites {
        subject:=course[0]
        preq:=course[1]
        if _, ok:= adjList[subject]; !ok {
            adjList[subject] = []int{preq}
        } else {
            adjList[subject] = append(adjList[subject], preq)
        }
    }

    // visited stores state of the node when dfs-ing
    // visited[node] =    1 -> visited
    // visited[node] =   -1 -> partially visited
    visited := make(map[int]int)
    for node:= range adjList{
        if visited[node] == 1{
            continue
        }
        if !dfs(node, adjList, visited){
            return false
        }
    }
    return true
}

func dfs(node int, adjList map[int][]int, visited map[int]int) bool {
    
    // non neighbors -> no prerequisites
    neighbourArr, ok := adjList[node]
    if !ok{
        return true
    }
    
    // encountered again while dfs-ing -> cycle present
    if visited[node] == -1{
        return false
    }
        
    // completely visited
    if visited[node] == 1{
        return true
    }
    
    // mark as partially visited while dfs-ing through neighbors
    visited[node] = -1
    
    // dfs through neighbors
    for _, neighbour := range neighbourArr{
        if !dfs(neighbour, adjList, visited){
            return false
        }
    }
    visited[node] = 1
    return true
}
