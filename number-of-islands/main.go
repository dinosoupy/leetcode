func numIslands(grid [][]byte) int {
    var dfs func(r, c int) bool
    dfs = func(r, c int) bool {
        if grid[r][c] != '1' {
            return false
        }
        grid[r][c] = '2'
        if r-1 >= 0 {
            dfs(r-1, c)
        }
        if r+1 < len(grid) {
            dfs(r+1, c)
        }
        if c-1 >= 0 {
            dfs(r, c-1)
        }
        if c+1 < len(grid[r]) {
            dfs(r, c+1)
        }
        return true
    }
    
    n := 0
    for r := range grid {
        for c := range grid[r] {
            if dfs(r, c) {
                n++
            }    
        }
    }
    return n
}