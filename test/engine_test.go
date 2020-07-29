package ep

import (
	e "ep"
	"testing"
)

func TestGetTokens(t *testing.T) {
	rule := "$name == Mary && ( $age > 20 && $age < 100 ) ||$number == 1234567890 && $item @ [aa,bb,cc]"
	_, err :=e.NewEpEngine(rule)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}
}

func TestCalExpression(t *testing.T) {

	rule := "$name == Mary && ( $age > 20 && $age < 100 ) ||$number == 1234567890 && $item @ [aa,bb,cc]"

	engine, _ := e.NewEpEngine(rule)

	m := map[string]string{"age": "30", "name": "Mary", "number": "1234567890", "item": "dd"}

	res := engine.Check(m)

	if res == true {
		t.Fatalf("Calculate error")

	}
}
