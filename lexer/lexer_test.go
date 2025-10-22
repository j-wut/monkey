
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

  let result = add(five, ten);

  !-/*%5;
  5 < 10 > 5;
  `

  tests := []struct {
    expectedType            token.TokenType
    expectedLiteral         string
    expectedLineNumber      int
    expectedLineCharacter int
  } {
    //let five = 5;
    {token.LET, "let", 1, 2},
    {token.IDENT, "five", 1, 6},
    {token.ASSIGN, "=", 1, 11},
    {token.INT, "5", 1, 13},
    {token.SEMICOLON, ";", 1, 14},
    //let ten = 10;
    {token.LET, "let", 2, 2},
    {token.IDENT, "ten", 2, 6},
    {token.ASSIGN, "=", 2, 10},
    {token.INT, "10", 2, 12},
    {token.SEMICOLON, ";", 2, 14},
    //let add = fn (x, y) {
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
    //  x + y;
    {token.IDENT, "x", 5, 4},
    {token.PLUS, "+", 5, 6},
    {token.IDENT, "y", 5, 8},
    {token.SEMICOLON, ";", 5, 9},
    //}
    {token.RBRACE, "}", 6, 2},
    //let result = add(five, ten);
    {token.LET, "let", 8, 2},
    {token.IDENT, "result", 8, 6},
    {token.ASSIGN, "=", 8, 13},
    {token.IDENT, "add", 8, 15},
    {token.LPAREN, "(", 8, 18},
    {token.IDENT, "five", 8, 19},
    {token.COMMA, ",", 8, 23},
    {token.IDENT, "ten", 8, 25},
    {token.RPAREN, ")", 8, 28},
    {token.SEMICOLON, ";", 8, 29},
    //!-/*5;
    {token.BANG, "!", 10, 2},
    {token.HYPHEN, "-", 10, 3},
    {token.SLASH, "/", 10, 4},
    {token.ASTERISK, "*", 10, 5},
    {token.MOD, "%", 10, 6},
    {token.INT, "5", 10, 7},
    {token.SEMICOLON, ";", 10, 8},
    //5 < 10 > 5;
    {token.INT, "5", 11, 2},
    {token.LT, "<", 11, 4},
    {token.INT, "10", 11, 6},
    {token.GT, ">", 11, 9},
    {token.INT, "5", 11, 11},
    {token.SEMICOLON, ";", 11, 12},

    {token.EOF, "", 12, 2},
  }

  

  l := New(input)
  for i, tt := range tests {
    tok := l.NextToken()

    if tok.Type != tt.expectedType {
      t.Fatalf("tests[%d::%v] - tokenType wrong. expected=%q, got=%q", i, tt, tt.expectedType, tok.Type)
    }
    if tok.Literal != tt.expectedLiteral {
      t.Fatalf("tests[%d::%v] - literal wrong. expected=%s, got=%s", i, tt, tt.expectedLiteral, tok.Literal)
    }
    if tok.LineNumber != tt.expectedLineNumber {
      t.Fatalf("tests[%d::%v] - lineNumber wrong. expected=%d, got=%d", i, tt, tt.expectedLineNumber, tok.LineNumber)
    }
    if tok.LineCharacter != tt.expectedLineCharacter {
      t.Fatalf("tests[%d::%v] - characterNumber wrong. expected=%d, got=%d", i, tt, tt.expectedLineCharacter, tok.LineCharacter)
    }
  }
}
