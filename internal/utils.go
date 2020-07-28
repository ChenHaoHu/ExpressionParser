package internal

import (
	"errors"
	"strings"
)

/**
	exp:num{$age}

	 type { $ name }


**/
func isVariables(str string) (error, *Variable) {

	strs := strings.Split(str, "{$")
	if len(strs) == 2 && str[len(str)-1] == byte('}') {
		strs[1] = strs[1][0 : len(strs[1])-1]
		err, variable := parseVariables(strs)
		if err != nil {
			return err, nil
		}
		return nil, variable
	}

	return errors.New(str + "is not a variable"), nil
}

func parseVariables(str []string) (error, *Variable) {

	if len(str) != 2 {
		return errors.New("error"), nil
	}

	variable := &Variable{
		name:    str[0],
		vartype: VarType(str[1]),
	}

	return nil, variable

}

func isCombinable(leftToken string, rightToken string) bool {
	if leftToken == "" {
		return false
	}

	token := leftToken + rightToken

	return isContain(CommonToken, token)
}

func isContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

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
	strs = append(strs, s[point:])

	return strs, nil
}
