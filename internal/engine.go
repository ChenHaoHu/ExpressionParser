package ep

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

	reversePolishTokens := toReversePolish(tokens)

	engine.ReversePolishTokens = reversePolishTokens

	log.Println("原表达式：", tokens)
	log.Println("逆波兰表达式：", reversePolishTokens)

	return engine, nil
}

func (engine *Engine) Calculate(context map[string]string) bool {

	log.Println("上下文变量有:", context)

	reversePolishTokens := engine.ReversePolishTokens
	res := calValue(reversePolishTokens, context)
	log.Println(res)
	return res.(bool)
}

func (engine *Engine) addVariable(variable *Variable) {
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

			_, variable, flag := isVariables(currToken)
			if flag == true {

				engine.addVariable(variable)
			}

			tokens = append(tokens, currToken)
			previewToken = currToken
		}
	}

	return tokens, nil
}
