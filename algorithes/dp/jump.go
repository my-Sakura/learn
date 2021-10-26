package dp

// O(N)
func jump(nums []int) int {
	count, l := 0, len(nums)-1

	for l > 0 {
		for i := 0; i < l; i++ {
			if i+nums[i] >= l {
				count++
				l = i
			}
		}
	}

	return count
}

// O(N)
// func jump(nums []int) int {
//     end := 0
//     maxPosition := 0
//     step := 0

//     for i := 0; i < len(nums) - 1; i++ {
//         maxPosition = max(maxPosition, i + nums[i])
//         if i == end {
//             end = maxPosition
//             step++
//         }
//     }

//     return step
// }

// func max(a, b int) int {
//     if a > b {
//         return a
//     }

//     return b
// }

func canJump(nums []int) bool {
	l := len(nums) - 1

	for l > 0 {
		for i := 0; i < l; i++ {
			if i+nums[i] >= l {
				l = i
				break
			}
			if i == l-1 && i+nums[i] < l {
				return false
			}
		}
	}

	return true
}

// O(N)
// func canJump(nums []int) bool {
//     if len(nums) <  2 {
//         return true
//     }
//     maxPosition := nums[0]
//      for i := 0; i < len(nums); i++ {
//          if i > maxPosition {
//              return false
//          }
//          maxPosition = max(maxPosition, i + nums[i])
//      }

//      return true
// }

// func max(i, j int) int {
//     if i > j {
//         return i
//     }

//     return j
// }
