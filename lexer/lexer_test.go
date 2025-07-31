
package lexer

import (
  "testing"
  "monkey/token"
)

func TestNextToken(token *testing.T) {
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
}
