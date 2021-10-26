package other

func generate(numRows int) [][]int {
	ret := make([][]int, 0)
	for i := 0; i < numRows; i++ {
		mid := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				mid[j] = 1
			} else {
				mid[j] = ret[i-1][j-1] + ret[i-1][j]
			}
		}
		ret = append(ret, mid)
	}

	return ret
}
