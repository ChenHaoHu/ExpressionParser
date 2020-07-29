package main

import (
	expression "expression/internal"
)

func main() {

	// rule := "num{$age} >= 20 && num{$age} < 100"
	rule := "100 >= 20 && 30 < 100 || 1200 == 300  &&(50/10)>5"

	expression.NewEngine(rule)

	rule = "100 >= 20 && 30 < 100 || 1200 == 300  &&((50/10)>5)"

	expression.NewEngine(rule)

	// rule = "str{$name} == 'Mary' && ( num{$age} > 20 && num{$age} < 100 ) || num{$number} == 1234567890"

	// expression.NewEngine(rule)

}
