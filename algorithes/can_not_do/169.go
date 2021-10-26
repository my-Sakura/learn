package can_not_do

func majorityElement(nums []int) int {
	var ret, count int

	for _, v := range nums {
		if count == 0 {
			ret = v
		}

		if ret == v {
			count++
		} else {
			count--
		}
	}

	return ret
}
