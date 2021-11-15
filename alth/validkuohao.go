package alth

func isValid(s string) bool {
	stack := make([]byte, 0)
	for i:=0;i<len(s);i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		}else if len(stack) != 0 && match(stack[len(stack)-1]) == s[i] {
			stack = stack[:len(stack)-1]
		}else {
			return false
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func match(ch byte) byte {
	if ch == '(' {
		return ')'
	}else if ch == '{' {
		return '}'
	}else {
		return ']'
	}
}