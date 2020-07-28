package internal

type VarType string

var SupportVarType []VarType

func init() {
	SupportVarType = []VarType{"num", "str", "ip"}
}

type Variable struct {
	name    string
	vartype VarType
}
