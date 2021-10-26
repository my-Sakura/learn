package other

func isValid(s string) bool {
	stack := make([]byte, 0)
	hashMap := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	for i := 0; i < len(s); i++ {
		if _, ok := hashMap[s[i]]; !ok {
			stack = append(stack, s[i])
			continue
		} else {
			if len(stack) == 0 {
				return false
			} else {
				if hashMap[s[i]] != stack[len(stack)-1] {
					return false
				}
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}
