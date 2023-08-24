func isValidSudoku(board [][]byte) bool {
    var row [9][9]bool
    var col [9][9]bool
    var box [9][9][9]bool

    for i, _:= range board {
        for j, _:= range board[i] {
            cell:=board[i][j]

            if cell=='.'{
               continue
            }

            digit:=cell-'0'-1

            if row[i][digit] || col[j][digit] || box[i/3][j/3][digit] {
                return false
            } else {
                row[i][digit]=true
                col[j][digit]=true
                box[i/3][j/3][digit]=true
            }
        }
    }
    return true
}