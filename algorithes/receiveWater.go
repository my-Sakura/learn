package main

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	sumWater := 0
	for i := 0; i < len(height); i++ {
		leftMax := maxWater(height, 0, i)
		rightMax := maxWater(height, i, len(height)-1)
		minHeight := min(leftMax, rightMax)
		if height[i] < minHeight {
			sumWater = sumWater + minHeight - height[i]
		} else {
			sumWater += 0
		}
	}

	return sumWater
}

func maxWater(height []int, start, end int) (maxHeight int) {
	for i := start; i <= end; i++ {
		if maxHeight <= height[i] {
			maxHeight = height[i]
		}
	}

	return
}

func min(i, j int) int {
	if i >= j {
		return j
	}

	return i
}

// dynamic programing
// func trap(height []int) int {
//     sumWater := 0
//     length := len(height)
// 	if length == 0 {
// 		return 0
// 	}

//     leftMax := make([]int, length)
//     leftMax[0] = height[0]
//     for i := 1; i < length; i++ {
//         leftMax[i] = max(leftMax[i-1], height[i])
//     }

//     rightMax := make([]int, length)
//     rightMax[length-1] = height[length-1]
//     for i := length - 2; i >=0 ; i-- {
//         rightMax[i] = max(rightMax[i+1], height[i])
//     }

//     for i := 0;i < length; i++ {
//         minHeight := min(leftMax[i], rightMax[i])
//         sumWater = sumWater + minHeight - height[i]
//     }

// 	return sumWater
// }

// func min(i, j int) int {
// 	if i >= j {
// 		return j
// 	}

// 	return i
// }

// func max(i, j int) int {
//     if i >= j {
//         return i
//     }

//     return j
// }
