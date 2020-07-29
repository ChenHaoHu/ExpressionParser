package ep

import (
	"fmt"
	"strconv"
)

func isContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
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
