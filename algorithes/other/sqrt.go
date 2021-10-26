package 

func mySqrt(x int) int {
	l, r := 0, x
	var ans int
 
	for l <= r {
		mid := (l + r) >> 1
		if mid * mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
 
	return ans
 }

// func mySqrt(x int) int {
// 	if x == 0 {
// 		return 0
// 	}

// 	ans := int(math.Exp(0.5 * math.Log(float64(x)))
// 	if (ans + 1) * (ans + 1) <= x {
// 		return ans + 1
// 	}

// 	return ans + 1
// }