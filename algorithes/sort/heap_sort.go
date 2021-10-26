package main

// import "fmt"

// func buildHeap(list []int) {
// 	for i := len(list)>>1 - 1; i >= 0; i-- {
// 		adjustHeap(list, i, len(list))
// 	}
// }

// func adjustHeap(list []int, pos int, l int) {
// 	for {
// 		child := 2*pos + 1
// 		if child > (l - 1) {
// 			break
// 		}

// 		if child+1 < l && list[child+1] > list[child] {
// 			child++
// 		}

// 		if list[pos] < list[child] {
// 			list[pos], list[child] = list[child], list[pos]
// 			pos = child
// 		} else {
// 			break
// 		}
// 	}
// }

// func heapSort(list []int) {
// 	buildHeap(list)

// 	for i := len(list) - 1; i >= 0; i-- {
// 		list[0], list[i] = list[i], list[0]
// 		adjustHeap(list, 0, i)
// 	}
// }

// func main() {
// 	s := []int{3, 1, 2, 5, 4}
// 	heapSort(s)
// 	fmt.Println(s)
// }
