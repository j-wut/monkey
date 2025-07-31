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
  characterNumber int
}

func New(input string) *Lexer {
  l := &Lexer{
    input: input,
    characterNumber: -1,
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
    for {
      l.character = l.input[l.readPosition]

      if l.character == '\n' {
        l.lineNumber += 1
        l.characterNumber = -1
      } else {
        break
      }
    }

  }
  l.position = l.readPosition
  l.readPosition += 1
  l.characterNumber += 1
}

func (l *Lexer) NextToken() token.Token {
  var tok token.Token

  var tokenType token.TokenType

  switch l.character {
  case '=':
    tokenType = token.ASSIGN
  case ';':
    tokenType = token.SEMICOLON
  case '(':
    tokenType = token.LPAREN
  case ')':
    tokenType = token.RPAREN
  case ',':
    tokenType = token.COMMA
  case '+':
    tokenType = token.PLUS
  case '{':
    tokenType = token.LBRACE
  case '}':
    tokenType = token.RBRACE
  case 0:
    return newToken(token.EOF, "", l.fileName, l.lineNumber, l.characterNumber)
  }

  tok = newToken(tokenType, string(l.character), l.fileName, l.lineNumber, l.characterNumber)
  l.readChar()
  return tok
}

func newToken(tokenType token.TokenType, ch string, fileName string, lineNumber int, characterNumber int) token.Token {
  return token.Token{
    Type: tokenType,
    Literal: ch,
    FileName: fileName,
    LineNumber: lineNumber,
    CharacterNumber: characterNumber,
  }
}
