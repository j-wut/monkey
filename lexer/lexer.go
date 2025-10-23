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

func (l *Lexer) peekChar() byte {
  return l.input[l.readPosition]
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
  for isDigit(l.peekChar()) {
    l.readChar()
  }
  return l.input[position:l.position + 1]
}

func isDigit(ch byte) bool {
  return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
  position := l.position
  for isLetter(l.peekChar()) {
    l.readChar()
  }
  return l.input[position:l.position + 1]
}

func isLetter(ch byte) bool {
  return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func (l *Lexer) NextToken() token.Token {
  var tok token.Token

  var tokenType token.TokenType
  
  var tokenStart int
  var tokenLiteral string
  multiChar := false
  

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
  case '+':
    tokenType = token.PLUS
  case '-':
    tokenType = token.HYPHEN
  case '*':
    tokenType = token.ASTERISK
  case '/':
    tokenType = token.SLASH
  case '%':
    tokenType = token.MOD
  case '=':
    if l.peekChar() == '=' {
      tokenStart = l.lineCharacter
      ch := l.character
      l.readChar()
      tokenLiteral = string(ch) + string(l.character)
      multiChar = true
      tokenType = token.EQ
    } else {
      tokenType = token.ASSIGN
    }
  case '!':
    if l.peekChar() == '=' {
      tokenStart = l.lineCharacter
      ch := l.character
      l.readChar()
      tokenLiteral = string(ch) + string(l.character)
      multiChar = true
      tokenType = token.NOT_EQ
    } else {
      tokenType = token.BANG
    }
  case '<':
    if l.peekChar() == '=' {
      tokenStart = l.lineCharacter
      ch := l.character
      l.readChar()
      tokenLiteral = string(ch) + string(l.character)
      multiChar = true
      tokenType = token.LT_EQ
    } else {
      tokenType = token.LT
    }
  case '>':
    if l.peekChar() == '=' {
      tokenStart = l.lineCharacter
      ch := l.character
      l.readChar()
      tokenLiteral = string(ch) + string(l.character)
      multiChar = true
      tokenType = token.GT_EQ
    } else {
      tokenType = token.GT
    }
  case '&':
    if l.peekChar() == '&' {
      tokenStart = l.lineCharacter
      ch := l.character
      l.readChar()
      tokenLiteral = string(ch) + string(l.character)
      multiChar = true
      tokenType = token.AND
    } else {
      tokenType = token.ILLEGAL
    }
  case '|':
    if l.peekChar() == '|' {
      tokenStart = l.lineCharacter
      ch := l.character
      l.readChar()
      tokenLiteral = string(ch) + string(l.character)
      multiChar = true
      tokenType = token.OR
    } else {
      tokenType = token.ILLEGAL
    }


  // BASE + Ident + Literal + Keywords
  case 0:
    return newToken(token.EOF, "", l.fileName, l.lineNumber, l.lineCharacter)
  default:
    if isLetter(l.character) {
      tokenStart = l.lineCharacter
      tokenLiteral = l.readIdentifier()
      multiChar = true
      tokenType = token.LookupIdent(tokenLiteral)
    } else if isDigit(l.character) {
      tokenStart = l.lineCharacter
      tokenLiteral = l.readNum()
      multiChar = true
      tokenType = token.INT
    } else {
      tokenType = token.ILLEGAL
    }
  }
  if multiChar {
    tok = newToken(tokenType, tokenLiteral, l.fileName, l.lineNumber, tokenStart)
  } else {
    tok = newToken(tokenType, string(l.character), l.fileName, l.lineNumber, l.lineCharacter)
  }
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
