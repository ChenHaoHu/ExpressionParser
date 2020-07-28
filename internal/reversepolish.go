package internal

import (
	"bytes"
)

func toExp(str string) []string {
	s := make([]string, 0)
	var t bytes.Buffer
	n := 0 // 用于判断括号是否成对
	for _, r := range str {
		if r == ' ' {
			// 去掉空格
			continue
		}
		if isDigit(r) {
			// 是数字 就写到缓存中
			t.WriteRune(r)
		} else {
			rs := string(r)
			if !isSign(rs) {
				panic("unknown sign: " + rs)
			}
			if t.Len() > 0 {
				// 遇到符号 把缓存中的数字 输出为数
				// 例如 将缓存中的 ['1', '2', '3'] 输出为 "123"
				s = append(s, t.String())
				t.Reset()
			}
			s = append(s, rs)
			if r == '(' {
				n++
			} else if r == ')' {
				n--
			}
		}
	}
	if t.Len() > 0 {
		// 最后一个操作符后面的数字 如果最后一个操作符是 ")" 那么 t.Len() 为0
		s = append(s, t.String())
	}
	if n != 0 {
		panic("the number of '(' is not equal to the number of ')' ")
	}
	return s
}

func isDigit(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func isSign(s string) bool {
	switch s {
	case "+", "-", "*", "/", "(", ")", "<=", ">=", "==", "&&", "||", "<", ">":
		return true
	default:
		return false
	}
}

func toPostfix(exp []string) []string {
	result := make([]string, 0)
	s := NewStack()
	for _, str := range exp {
		if isSign(str) {
			if str == "(" || s.Len() == 0 {
				// "(" 或者 栈为空 直接进栈
				// 括号中的计算 需要单独处理 相当于一个新的上下文
				// 如果栈为空 需要先进栈 和后续操作符比较优先级之后 才能决定计算顺序
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
	case "(", ")":
		return 3
	case "*", "/":
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

func calValue(exp []string) interface{} {
	s := NewStack()
	for _, str := range exp {
		if isSign(str) {
			// 如果是符号 弹出栈顶的两个元素 进行计算
			// 因为栈结构先进后出 所以先弹出b
			b := getTopV(s)
			a := getTopV(s)
			var n interface{}
			switch str {
			case "+":
				n = AddOperation(a, b)
			case "-":
				n = ReduceOperation(a, b)
			case "*":
				n = MultiplicationOperation(a, b)
			case "/":
				n = DivisionOperation(a, b)
			case "==":
				n = DoublequalOperation(a, b)
			case ">=":
				n = GreaterOrEqualOperation(a, b)
			case "<=":
				n = LesserOrEqualOperation(a, b)
			case ">":
				n = GreaterOperation(a, b)
			case "<":
				n = LesserOperation(a, b)
			case "||":
				n = OrOperation(a, b)
			case "&&":
				n = AndOperation(a, b)
			}
			s.Push(n)
		} else {
			s.Push(str)
		}
	}
	return getTopV(s)
}
