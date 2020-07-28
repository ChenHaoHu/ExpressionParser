package internal

import (
	"errors"
	"fmt"
	"strconv"
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

func turnToInt(v interface{}) int {
	switch v.(type) {
	case int: // push进去的计算结果为int
		return v.(int)
	case string: // exp中的数据为string
		if i, err := strconv.Atoi(v.(string)); err != nil {
			panic(err)
		} else {
			return i
		}
	}
	panic(fmt.Sprintf("unknown value type: %T", v))
}
