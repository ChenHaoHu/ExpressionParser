package ep

import (
	"errors"
	"strings"
)

var SupportVarType []string

func init() {
	SupportVarType = []string{"num", "str", "ip"}
}

type Variable struct {
	Name  string
	Vtype string
}

func getVariablesValue(v *Variable, context map[string]string) (interface{}, string, error) {
	name := v.Name
	vtype := v.Vtype

	vtype = strings.ToLower(vtype)
	name = strings.ToLower(name)
	vv, ok := context[name]
	if !ok {
		return nil, "", errors.New("上下文中缺少变量")
	}

	return vv, vtype, nil

}

/**
	exp:num{$age}

	type { $ name }

**/
func isVariables(str string) (error, *Variable, bool) {

	//因为放弃显示标注类型 这里直接判断是否 $ 开头
	if str[0] == '$' {
		strs := []string{}
		strs = append(strs, "common")
		strs = append(strs, str[1:])
		err, variable := parseVariables(strs)
		if err != nil {
			return err, nil, false
		}
		return nil, variable, true
	}

	strs := strings.Split(str, "{$")
	if len(strs) == 2 && str[len(str)-1] == byte('}') {
		strs[1] = strs[1][0 : len(strs[1])-1]
		err, variable := parseVariables(strs)
		if err != nil {
			return err, nil, false
		}
		return nil, variable, true
	}

	return errors.New(str + "is not a variable"), nil, false
}

func parseVariables(str []string) (error, *Variable) {

	if len(str) != 2 {
		return errors.New("error"), nil
	}

	variable := &Variable{
		Name:  str[1],
		Vtype: str[0],
	}

	return nil, variable

}
