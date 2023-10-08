func pacificAtlantic(heights [][]int) [][]int {
    m, n := len(heights), len(heights[0]) 
    pacific, atlantic := make([][]bool, m), make([][]bool, m)

    var dfs func(prev, row, col int, visited [][]bool)
    dfs = func(prev, row, col int, visited [][]bool) {
        if row < 0 || row >= m || col < 0 || col >= n || visited[row][col] || prev > heights[row][col]{ return }
        
        cur := heights[row][col]
        visited[row][col] = true

        dfs(cur, row + 1, col, visited)
        dfs(cur, row - 1, col, visited)
        dfs(cur, row, col + 1, visited)
        dfs(cur, row, col - 1, visited)
    }

    for i := 0; i < m; i++ {
        pacific[i] = make([]bool, n)
        atlantic[i] = make([]bool, n)
    }

    for i := 0; i < m; i++ {
        dfs(-1, i, 0, pacific)
        dfs(-1, i, n - 1, atlantic)
    }

    for j := 0; j < n; j++ {
        dfs(-1, 0, j, pacific)
        dfs(-1, m - 1, j, atlantic)
    }

    res:=[][]int{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if pacific[i][j] && atlantic[i][j] {
                res = append(res, []int{i, j})
            }
        }
    }

    return res
}