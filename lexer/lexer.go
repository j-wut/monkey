package lexer

import (
  "github.com/j-wut/monkey/token"
)

type Lexer struct {
  input           string
  position        int
  readPosition    int
  character       byte
  fileName        string
  lineNumber      int
  lineCharacter int
}

func New(input string) *Lexer {
  l := &Lexer{
    input: input,
    lineCharacter: -1,
    fileName: "",
  }
  l.readChar()
  return l
}

func (l *Lexer) readChar() {
  if l.readPosition >= len(l.input) {
    l.character = 0
    if l.readPosition > len(l.input) {
      return
    }
  } else {
      l.character = l.input[l.readPosition]
  }
  l.position = l.readPosition
  l.readPosition += 1
  l.lineCharacter += 1
}

func (l *Lexer) skipWhitespace() {
  for l.character == ' ' || l.character == '\t' || l.character == '\n' {
    lastChar := l.character
    l.readChar()
    if lastChar == '\n' {
      l.lineCharacter = 0
      l.lineNumber += 1
    }
  }
}

func (l *Lexer) readNum() string {
  position := l.position
  for isDigit(l.character) {
    l.readChar()
  }
  return l.input[position:l.position]
}

func isDigit(ch byte) bool {
  return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
  position := l.position
  for isLetter(l.character) {
    l.readChar()
  }
  return l.input[position:l.position]
}

func isLetter(ch byte) bool {
  return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func (l *Lexer) NextToken() token.Token {
  var tok token.Token

  var tokenType token.TokenType

  l.skipWhitespace()

  switch l.character {
  // Deliminters
  case '(':
    tokenType = token.LPAREN
  case ')':
    tokenType = token.RPAREN
  case ',':
    tokenType = token.COMMA
  case ';':
    tokenType = token.SEMICOLON
  case '{':
    tokenType = token.LBRACE
  case '}':
    tokenType = token.RBRACE
  // Operators
  case '=':
    tokenType = token.ASSIGN
  case '+':
    tokenType = token.PLUS
  case '-':
    tokenType = token.HYPHEN
  case '*':
    tokenType = token.ASTERISK
  case '/':
    tokenType = token.SLASH
  case '!':
    tokenType = token.BANG
  case '%':
    tokenType = token.MOD
  case '<':
    tokenType = token.LT
  case '>':
    tokenType = token.GT

  // BASE + Ident + Literal + Keywords
  case 0:
    return newToken(token.EOF, "", l.fileName, l.lineNumber, l.lineCharacter)
  default:
    if isLetter(l.character) {
      identLineCharacter := l.lineCharacter
      identLiteral := l.readIdentifier()
      return newToken(token.LookupIdent(identLiteral), identLiteral, l.fileName, l.lineNumber, identLineCharacter)
    } else if isDigit(l.character) {
      numLineCharacter := l.lineCharacter
      numLiteral := l.readNum()
      return newToken(token.INT, numLiteral, l.fileName, l.lineNumber, numLineCharacter)
    } else {
      return newToken(token.ILLEGAL, string(l.character), l.fileName, l.lineNumber, l.lineCharacter)
    }
  }

  tok = newToken(tokenType, string(l.character), l.fileName, l.lineNumber, l.lineCharacter)
  l.readChar()
  return tok
}

func newToken(tokenType token.TokenType, ch string, fileName string, lineNumber int, lineCharacter int) token.Token {
  return token.Token{
    Type: tokenType,
    Literal: ch,
    FileName: fileName,
    LineNumber: lineNumber,
    LineCharacter: lineCharacter,
  }
}
