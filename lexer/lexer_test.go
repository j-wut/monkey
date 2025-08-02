
package lexer

import (
  "testing"
  "github.com/j-wut/monkey/token"
)


func TestNextToken(t *testing.T) {
  input := `
  let five = 5;
  let ten = 10;

  let add = fn (x, y) {
    x + y;
  }
  `

  tests := []struct {
    expectedType            token.TokenType
    expectedLiteral         string
    expectedLineNumber      int
    expectedLineCharacter int
  } {
    {token.LET, "let", 1, 2},
    {token.IDENT, "five", 1, 6},
    {token.ASSIGN, "=", 1, 11},
    {token.INT, "5", 1, 13},
    {token.SEMICOLON, ";", 1, 14},
    {token.LET, "let", 2, 2},
    {token.IDENT, "ten", 2, 6},
    {token.ASSIGN, "=", 2, 10},
    {token.INT, "10", 2, 12},
    {token.SEMICOLON, ";", 2, 14},
    {token.LET, "let", 4, 2},
    {token.IDENT, "add", 4, 6},
    {token.ASSIGN, "=", 4, 10},
    {token.FUNCTION, "fn", 4, 12},
    {token.LPAREN, "(", 4, 15},
    {token.IDENT, "x", 4, 16},
    {token.COMMA, ",", 4, 17},
    {token.IDENT, "y", 4, 19},
    {token.RPAREN, ")", 4, 20},
    {token.LBRACE, "{", 4, 22},
    {token.IDENT, "x", 5, 4},
    {token.PLUS, "+", 5, 6},
    {token.IDENT, "y", 5, 8},
    {token.SEMICOLON, ";", 5, 9},
    {token.RBRACE, "}", 6, 2},
    {token.EOF, "", 7, 2},
  }

  

  l := New(input)
  for i, tt := range tests {
    tok := l.NextToken()

    if tok.Type != tt.expectedType {
      t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
    }
    if tok.Literal != tt.expectedLiteral {
      t.Fatalf("tests[%d] - literal wrong. expected=%s, got=%s", i, tt.expectedLiteral, tok.Literal)
    }
    if tok.LineNumber != tt.expectedLineNumber {
      t.Fatalf("tests[%d] - lineNumber wrong. expected=%d, got=%d", i, tt.expectedLineNumber, tok.LineNumber)
    }
    if tok.LineCharacter != tt.expectedLineCharacter {
      t.Fatalf("tests[%d] - characterNumber wrong. expected=%d, got=%d", i, tt.expectedLineCharacter, tok.LineCharacter)
    }
  }
}
