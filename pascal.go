func generate(numRows int) [][]int {
    rows := make([][]int, numRows)
    for i:=0; i<numRows; i++ {
        rows[i] = make([]int, i+1)
        for j:=0; j<=i; j++ {
            if j == 0 || j == i {
                rows[i][j] = 1
            } else {
                rows[i][j] = rows[i-1][j-1] + rows[i-1][j]
            }
        }
    }
    return rows
}

