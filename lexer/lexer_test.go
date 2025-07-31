
package lexer

import (
  "testing"
  "github.com/j-wut/monkey/token"
)

func TestNextToken(t *testing.T) {
  input := `=+(){},;`

  tests := []struct {
    expectedType            token.TokenType
    expectedLiteral         string
    expectedLineNumber      int
    expectedCharacterNumber int
  } {
    {token.ASSIGN, "=", 0, 0},
    {token.PLUS, "+", 0, 1},
    {token.LPAREN, "(", 0, 2},
    {token.RPAREN, ")", 0, 3},
    {token.LBRACE, "{", 0, 4},
    {token.RBRACE, "}", 0, 5},
    {token.COMMA, ",", 0, 6},
    {token.SEMICOLON, ";", 0, 7},
    {token.EOF, "", 0, 8},
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
    if tok.CharacterNumber != tt.expectedCharacterNumber {
      t.Fatalf("tests[%d] - characterNumber wrong. expected=%d, got=%d", i, tt.expectedCharacterNumber, tok.CharacterNumber)
    }
  }
}
