package main

// func partition(list []int, low, high int) int {
// 	pivot := list[low]
// 	for low < high {
// 		for low < high && pivot <= list[high] {
// 			high--
// 		}
// 		list[low] = list[high]

// 		for low < high && pivot >= list[low] {
// 			low++
// 		}
// 		list[high] = list[low]
// 	}
// 	list[low] = pivot

// 	return low
// }

// func quickSort(list []int, low, high int) {
// 	if high > low {
// 		pivot := partition(list, low, high)
// 		quickSort(list, low, pivot - 1)
// 		quickSort(list, pivot + 1, high)
// 	}
// }

// func Sort(list []int) {
// 	quickSort(list, 0, len(list)-1)
// }

// func main() {
// 	s := []int{5, 1, 1, 2, 0, 0}
// 	Sort(s)
// 	fmt.Println(s)
// }
