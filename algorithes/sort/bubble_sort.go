package main

// import "fmt"

l - 1 cycle && l - 1 exchange
func bubbleSort(list []int) {
	for i := 1; i < len(list); i++ {
		for j := 0; j < len(list)-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

// func main() {
// 	s := []int{2, 1, 3, 6, 5, 4, 7, 5}
// 	bubbleSort(s)
// 	fmt.Println(s)
// }
