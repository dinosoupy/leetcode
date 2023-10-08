func pacificAtlantic(heights [][]int) [][]int {
    // m -> rows
    // n -> cols
    m, n := len(heights), len(heights[0]) 

    // two sets storing all coordinates from where water can flow to the respective oceans
    // common coordinates from these will form the res array
    pacific, atlantic := make([][]bool, m), make([][]bool, m)

    // make bool array for each row
    for i := 0; i < m; i++ {
        pacific[i] = make([]bool, n)
        atlantic[i] = make([]bool, n)
    }

    // dfs through every row in first and last col
    for i := 0; i < m; i++ {
        dfs(-1, i, 0, pacific, heights)
        dfs(-1, i, n - 1, atlantic, heights)
    }

    // dfs through every col in first and last row
    for j := 0; j < n; j++ {
        dfs(-1, 0, j, pacific, heights)
        dfs(-1, m - 1, j, atlantic, heights)
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

func dfs(prev, row, col int, visited [][]bool, heights [][]int) {
    // get rows and cols again instead of passing it to the function dfs as arg
    m, n := len(heights), len(heights[0])

    // dfs searched up down left right, some of these may go out of bounds for the island
    // Hence, check row < 0 || row >= m || col < 0 || col >= n 
    // check if cell has been visited, no need to revisit
    // check if water can flow from current cell to previous 
    if row < 0 || row >= m || col < 0 || col >= n || visited[row][col] || prev > heights[row][col]{ return }
    
    // if the statement above executes without returning, water can flow from row, col to visited ocean
    cur := heights[row][col]
    visited[row][col] = true

    // dfs all directions
    dfs(cur, row + 1, col, visited, heights)
    dfs(cur, row - 1, col, visited, heights)
    dfs(cur, row, col + 1, visited, heights)
    dfs(cur, row, col - 1, visited, heights)
}