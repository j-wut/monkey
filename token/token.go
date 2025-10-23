package token

type TokenType string

type Token struct {
 Type             TokenType
 Literal          string
 FileName         string
 LineNumber       int
 LineCharacter    int
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
  HYPHEN = "-"
  BANG = "!"
  ASTERISK = "*"
  SLASH = "/"
  MOD = "%"

  LT = "<"
  GT = ">"
  EQ = "=="
  NOT_EQ = "!="
  LT_EQ = "<="
  GT_EQ = ">="
  AND = "&&"
  OR = "||"
  
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
  IF = "IF"
  ELSE = "ELSE"
  RETURN = "RETURN"
  TRUE = "TRUE"
  FALSE = "FALSE"
)

var keywords = map[string]TokenType{
  "fn": FUNCTION,
  "let": LET,
  "if": IF,
  "else": ELSE,
  "return": RETURN,
  "true": TRUE,
  "false": FALSE,
}

func LookupIdent(ident string) TokenType {
  if tok, ok := keywords[ident]; ok {
    return tok
  }
  return IDENT
}
