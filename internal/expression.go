package internal

import (
	"errors"
	"log"
	"strings"
)

type Engine struct {
	Rule                string
	Tokens              []string
	ReversePolishTokens []string
	Variables           []*Variable
}

var Delims = []rune{'+', '-', '*', '/', '(', ')', '&', '|', '!', '=', '>', '<', '?'}
var CommonToken = []string{"&&", "||", "<=", ">=", "=="}

func NewEngine(rule string) (*Engine, error) {
	engine := &Engine{
		Rule:      rule,
		Variables: make([]*Variable, 10),
	}
	tokens, err := engine.sliceTokens()
	if err != nil {
		return nil, err
	}
	engine.Tokens = tokens
	log.Println(tokens)
	engine.toReversePolish()

	return engine, nil
}

func (engine *Engine) AddVariable(variable *Variable) {
	engine.Variables = append(engine.Variables, variable)
}

func (engine *Engine) sliceTokens() ([]string, error) {
	rule := engine.Rule
	if rule == "" {
		return nil, errors.New("rule can not be nil")
	}
	rule = strings.ReplaceAll(rule, " ", "")

	strs, _ := split(rule, Delims)

	tokens := []string{}

	var previewToken string

	strLen := len(strs)

	for i := 0; i < strLen; i++ {
		currToken := strs[i]
		if isCombinable(previewToken, currToken) {
			//移除最后一个
			tokens = tokens[:len(tokens)-1]
			newToken := previewToken + currToken
			tokens = append(tokens, newToken)
			previewToken = newToken
		} else {

			err, variable := isVariables(currToken)
			if err == nil {

				engine.AddVariable(variable)
			}

			tokens = append(tokens, currToken)
			previewToken = currToken
		}
	}

	return tokens, nil
}

func (engine *Engine) toReversePolish() {
	tokens := engine.Tokens
	postfixExp := toPostfix(tokens)
	// printExp(postfixExp)
	log.Println(postfixExp)
	res := calValue(postfixExp)
	log.Println(res)

}
