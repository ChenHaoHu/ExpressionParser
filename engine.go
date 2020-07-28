package main

import (
	expression "expression/internal"
)

func main() {

	rule := "num{$age} >= 20 && num{$age} < 100"

	expression.NewEngine(rule)

	rule = "str{$name} == 'Mary' && ( num{$age} > 20 && num{$age} < 100 ) || num{$number} == 1234567890"

	expression.NewEngine(rule)

}
