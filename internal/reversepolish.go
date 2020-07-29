package ep

import "log"

func isDigit(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func isSign(s string) bool {
	switch s {
	case "+", "-", "*", "/", "(", ")", "<=", ">=", "==", "&&", "||", "<", ">", "@":
		return true
	default:
		return false
	}
}

func toReversePolish(exp []string) []string {
	result := make([]string, 0)
	s := NewStack()
	for _, str := range exp {
		if isSign(str) {
			if str == "(" || s.Len() == 0 {

				s.Push(str)
			} else {
				if str == ")" {
					// 若为 ")" 依次弹出栈顶元素并输出 直到遇到 "("
					for s.Len() > 0 {
						if s.Peek().(string) == "(" {
							s.Pop()
							break
						}
						result = appendStr(result, s.Pop().(string))
					}
				} else {
					// 判断其与栈顶符号的优先级
					// 如果栈顶是 "(" 说明是新的上下文 不能相互比较优先级
					for s.Len() > 0 && s.Peek().(string) != "(" && signCompare(str, s.Peek().(string)) <= 0 {
						// 当前符号的优先级 不大于栈顶元素 弹出栈顶元素并输出
						// 优先级高的操作 需要先计算
						// 优先级相同 因为栈中的操作是先放进去的 也需要先计算
						result = appendStr(result, s.Pop().(string))
					}
					// 当前符号入栈
					s.Push(str)
				}
			}
		} else {
			// 若是数字就输出
			result = appendStr(result, str)
		}
	}
	for s.Len() > 0 {
		result = appendStr(result, s.Pop().(string))
	}
	return result
}

func appendStr(slice []string, str string) []string {
	if str == "(" || str == ")" {
		// 后缀表达式 不包含括号
		return slice
	}
	return append(slice, str)
}

func signCompare(a, b string) int {
	return getSignValue(a) - getSignValue(b)
}

func getSignValue(a string) int {
	switch a {
	case "@":
		return 5
	case "(", ")":
		return 4
	case "*", "/":
		return 3
	case "+", "-":
		return 2
	case ">", "<", "<=", ">=", "==":
		return 1
	default:
		return 0
	}
}

func getTopV(s *Stack) interface{} {
	v := s.Pop()
	return v
}

func calValue(exp []string, context map[string]string) interface{} {
	s := NewStack()
	for _, str := range exp {
		if isSign(str) {
			// 如果是符号 弹出栈顶的两个元素 进行计算
			// 因为栈结构先进后出 所以先弹出b
			b := getTopV(s)
			a := getTopV(s)
			var n interface{}
			n = OperationHandler(a, b, context, str)
			log.Println(a, str, b, "=", n)
			s.Push(n)
		} else {
			s.Push(str)
		}
	}
	return getTopV(s)
}
