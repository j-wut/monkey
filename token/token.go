package token

type TokenType string

type Token struct {
 Type             TokenType
 Literal          string
 FileName         string
 LineNumber       int
 CharacterNumber  int
}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"
  
  // Identifiers + Literals
  IDENT = "IDENT"
  INT = "INT"

  // Operators
  ASSIGN = "="
  PLUS = "+"
  
  // Delimiters
  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  // Keywords
  FUNCTION = "FUNCTION"
  LET = "LET"

)
