package _221

func maximalSquare(matrix [][]byte) int {
	xl,yl := len(matrix),len(matrix[0])
	dp := make([][]int,xl)
	max := 0
	for  x := 0;x < xl; x++ {
		dp[x] = append(dp[x],make([]int,yl)...)
		for y := 0;y < yl;y++ {
			if matrix[x][y] == '1' {
				dp[x][y] = 1
				max = 1
			}
		}
	}
	for  x := 1;x < xl; x++ {
		for y := 1;y < yl;y++ {
			if matrix[x][y] == '1' {
				dp[x][y] = mmin(mmin(dp[x][y-1],dp[x-1][y]),dp[x-1][y-1]) +1
				max = mmax(max,dp[x][y])
			}
		}
	}
	return max * max
}

func mmin(x,y int) int {
	if  x < y {
		return x
	}
	return y
}

func mmax(x,y int) int {
	if x > y {
		return x
	}
	return y
}
