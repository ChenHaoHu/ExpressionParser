package ep

import (
	"log"
	"strings"
)

// operation.go 文件主要描述每一个运算符的具体操作
// 运算时，传入运算符两边的 运算数， 进入方法时，均以interface{}类型传入
// 不同的运算符处理的数据类型不一样，比如 加减乘除处理的是数值类型，&&与||处理的主要是布尔类型
//所以需要我们在程序中，根据不同的运算符去转换，但是在处理之前，我们需要校验当前值是否是一个需要用户上下文传入的变量
//如果以上出现错误，会报相对的异常

//符号运算预处理
//检查是否是变量 如果是变量 转成变量对应的格式 然后返回
func checkV(v interface{}, context map[string]string) (interface{}, string) {

	switch v.(type) {
	case string:
		err, variable, flag := isVariables(v.(string))
		if flag == false {
			return v, ""
		}
		v, t, err := getVariablesValue(variable, context)
		if err != nil {
			panic(err)
		}

		log.Println(variable.Name, "---->", v)

		return v, t

	default:
		return v, ""
	}

}

func OperationHandler(v1 interface{}, v2 interface{}, context map[string]string, opt string) interface{} {

	log.Println(v1, opt, v2)

	//预处理两个变量
	v1, _ = checkV(v1, context)
	v2, _ = checkV(v2, context)

	log.Println(v1, opt, v2)

	switch opt {
	case "+":
		//加号运算符 两端必须为数值类型 暂且只支持 int 类型
		return AddOperation(turnToInt(v1), turnToInt(v2))
	case "-":
		return ReduceOperation(turnToInt(v1), turnToInt(v2))
	case "*":
		return MultiplicationOperation(turnToInt(v1), turnToInt(v2))
	case "/":
		return DivisionOperation(turnToInt(v1), turnToInt(v2))
	case ">":
		return GreaterOperation(turnToInt(v1), turnToInt(v2))
	case "<":
		return LesserOperation(turnToInt(v1), turnToInt(v2))
	case ">=":
		return GreaterOrEqualOperation(turnToInt(v1), turnToInt(v2))
	case "<=":
		return LesserOrEqualOperation(turnToInt(v1), turnToInt(v2))
	case "==":
		return DoublequalOperation(v1, v2)
	case "&&":
		return AndOperation(v1.(bool), v2.(bool))
	case "||":
		return OrOperation(v1.(bool), v2.(bool))
	case "@":
		return InOperation(v1.(string), v2.(string))
	default:
		panic("不支持的操作符号:" + opt)
	}

}

func InOperation(a string, b string) bool {
	b = b[1 : len(b)-1]
	items := strings.Split(b, ",")

	for _, item := range items {
		if strings.EqualFold(item, a) {
			return true
		}
	}
	return false

}

func OrOperation(a bool, b bool) bool {

	if a || b {
		return true
	} else {
		return false
	}
}

func AndOperation(a bool, b bool) bool {

	if a && b {
		return true
	} else {
		return false
	}
}

func AddOperation(a int, b int) int {

	return a + b

}
func ReduceOperation(a int, b int) int {

	return a - b
}
func MultiplicationOperation(a int, b int) int {

	return a * b

}
func DivisionOperation(a int, b int) int {

	return a / b

}

func DoublequalOperation(a interface{}, b interface{}) bool {

	//因为相等判断 比较复杂 故写在这里
	//判断a和b的类型 然后判断

	if a == b {
		return true
	} else {
		return false
	}
}

func GreaterOrEqualOperation(a int, b int) bool {

	if a >= b {
		return true
	} else {
		return false
	}
}

func GreaterOperation(a int, b int) bool {

	if a > b {
		return true
	} else {
		return false
	}
}

func LesserOrEqualOperation(a int, b int) bool {

	if a <= b {
		return true
	} else {
		return false
	}
}

func LesserOperation(a int, b int) bool {

	if a < b {
		return true
	} else {
		return false
	}
}
