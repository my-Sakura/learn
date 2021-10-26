package main

// import "fmt"

// func main() {
// 	s := []int{6, 7, 8, 10, 4, 6, 99}
// 	res := mergeSort(s)
// 	fmt.Println(res)
// }

// func mergeSort(list []int) []int {
// 	if len(list) <= 1 {
// 		return list
// 	}

// 	m := len(list) >> 1
// 	leftS := mergeSort(list[:m])
// 	rightS := mergeSort(list[m:])

// 	return merge(leftS, rightS)
// }

// func merge(l, r []int) []int {
// 	lIndex, rIndex := 0, 0
// 	res := make([]int, 0)

// 	for lIndex < len(l) && rIndex < len(r) {
// 		if l[lIndex] > r[rIndex] {
// 			res = append(res, r[rIndex])
// 			rIndex++
// 		} else {
// 			res = append(res, l[lIndex])
// 			lIndex++
// 		}
// 	}

// 	if lIndex < len(l) {
// 		res = append(res, l[lIndex:]...)
// 	}

// 	if rIndex < len(r) {
// 		res = append(res, r[rIndex:]...)
// 	}

// 	return res
// }
