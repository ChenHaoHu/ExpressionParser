package internal

import (
	"log"
)

func OrOperation(a interface{}, b interface{}) bool {
	log.Println(a, "------ || ------", b)

	if a.(bool) || b.(bool) {
		return true
	} else {
		return false
	}
}

func AndOperation(a interface{}, b interface{}) bool {
	log.Println(a, "----- && -------", b)
	if a.(bool) && b.(bool) {
		return true
	} else {
		return false
	}
}

func AddOperation(a interface{}, b interface{}) int {
	aInt := turnToInt(a)
	bInt := turnToInt(b)
	return aInt + bInt

}
func ReduceOperation(a interface{}, b interface{}) int {
	aInt := turnToInt(a)
	bInt := turnToInt(b)
	return aInt - bInt
}
func MultiplicationOperation(a interface{}, b interface{}) int {
	aInt := turnToInt(a)
	bInt := turnToInt(b)
	return aInt * bInt

}
func DivisionOperation(a interface{}, b interface{}) int {
	aInt := turnToInt(a)
	bInt := turnToInt(b)
	return aInt / bInt

}

func DoublequalOperation(a interface{}, b interface{}) bool {
	log.Println(a, "------ == ------", b)

	aInt := turnToInt(a)
	bInt := turnToInt(b)

	if aInt == bInt {
		return true
	} else {
		return false
	}
}

func GreaterOrEqualOperation(a interface{}, b interface{}) bool {
	log.Println(a, "----- >= -------", b)

	aInt := turnToInt(a)
	bInt := turnToInt(b)

	if aInt >= bInt {
		return true
	} else {
		return false
	}
}

func GreaterOperation(a interface{}, b interface{}) bool {
	log.Println(a, "------ > ------", b)

	aInt := turnToInt(a)
	bInt := turnToInt(b)
	if aInt > bInt {
		return true
	} else {
		return false
	}
}

func LesserOrEqualOperation(a interface{}, b interface{}) bool {
	log.Println(a, "----- <= -------", b)

	aInt := turnToInt(a)
	bInt := turnToInt(b)
	if aInt <= bInt {
		return true
	} else {
		return false
	}
}

func LesserOperation(a interface{}, b interface{}) bool {
	log.Println(a, "------ < ------", b)

	aInt := turnToInt(a)
	bInt := turnToInt(b)
	if aInt < bInt {
		return true
	} else {
		return false
	}
}
