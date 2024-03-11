package ast

import (
	"github.com/Germanchrystan/Interpreter-In-Go/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program {
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{ Type: token.LET, Literal: "let"},
				Name: &Idenytifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String8) wrong. got=%q", program.String())
	}
}