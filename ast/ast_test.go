package ast

import (
  "testing"

  "github.com/j-wut/monkey/token"
)

func TestString(t *testing.T) {
  program := &Program{
    Statements: []Statement{
      &LetStatement{
        Token: token.Token{Type: token.LET, Literal: "let"},
        Name: &Identifier{
          Token: token.Token{Type: token.IDENT, Literal: "myVar"},
          Value: "myVar",
        },
        Value: &Identifier{
          Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
          Value: "anotherVar",
        },
      },
      &ReturnStatement{
        Token: token.Token{Type: token.RETURN, Literal: "return"},
        ReturnValue: &Identifier{
          Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
          Value: "anotherVar",
        },
      },
      &ExpressionStatement{
        Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
        Expression: &Identifier{
          Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
          Value: "anotherVar",
        },
      },
    },
  }

  if program.String() != "let myVar = anotherVar;return anotherVar;anotherVar;" {
    t.Errorf("program.String() wrong. got=%q", program.String())
  }
}
