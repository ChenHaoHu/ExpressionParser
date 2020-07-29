package ep

func split(s string, sep []rune) ([]string, error) {

	strs := []string{}
	point := 0
	for index, runeValue := range s {
		for i := 0; i < len(sep); i++ {
			if runeValue == sep[i] {
				if point < index {

					strs = append(strs, s[point:index])

				}
				strs = append(strs, string(sep[i]))
				point = index + 1
				break
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

	return isContain(CommonToken, token)
}
