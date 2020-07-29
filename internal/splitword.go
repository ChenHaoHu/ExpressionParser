package ep


func split(s string, sep []rune) ([]string, error) {

	strs := []string{}
	point := 0
	flag := false

	l := []rune{'`'}


	for index, runeValue := range s {
		//为了适配正则 防止部分正则表达 影响 词法切割
		//规定：在闭合的 ` 符号中 不处理符号
		if runeValue == l[0]  {
			flag = !flag
		}
		if flag == false {
			for i := 0;  i < len(sep); i++ {
				if runeValue == sep[i]  {
					if point < index {
						strs = append(strs, s[point:index])
					}
					strs = append(strs, string(sep[i]))
					point = index + 1
					break
				}
			}
		}


	}

	if len(s) > point {
		strs = append(strs, s[point:])
	}

	return strs, nil
}

func isCombinable(leftToken string, rightToken string) bool {
	if leftToken == "" {
		return false
	}

	token := leftToken + rightToken

	return isContain(CommonDoubleToken, token)
}
