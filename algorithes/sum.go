package main

func sum(array []int) int {
	if len(array) == 0 {
		return 0
	}
	if len(array) == 1 {
		return array[0]
	}

	return array[0] + sum(array[1:])
}
