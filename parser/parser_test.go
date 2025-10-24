package parser

import (
  "testing"

  "github.com/j-wut/monkey/ast"
  "github.com/j-wut/monkey/lexer"
)

func checkParserErrors(t *testing.T, p *Parser) {
  if len(p.Errors) == 0 {
    return
  }

  t.Errorf("parser has %d errors", len(p.Errors))
  for _, msg := range p.Errors {
    t.Errorf("parser error: %q", msg)
  }
  t.FailNow()
}


func TestLetStatements(t *testing.T) {
  input := `
  let x=5;
  let y=10;
  let foobar=838383;
  `

  l := lexer.New(input)
  p := New(l)
  
  program := p.ParseProgram()
  if program == nil {
    t.Fatalf("ParseProgram() returned nil")
  }
  checkParserErrors(t, p)
  if len(program.Statements) != 3 {
    t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements)) 
  }

  tests := []struct {
    expectedIdentifier string
  }{
    {"x"},
    {"y"},
    {"foobar"},
  }

  for i, tt := range tests {
    stmt := program.Statements[i]
    if !testLetStatement(t, stmt, tt.expectedIdentifier) {
      return
    }
  }
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
  if s.TokenLiteral() != "let" {
    t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
    return false
  }

  letStmt, ok := s.(*ast.LetStatement)
  if !ok {
    t.Errorf("s not *ast.LetStatement. got %T", s)
    return false
  }

  if letStmt.Name.Value != name {
    t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
    return false
  }

  if letStmt.Name.TokenLiteral() != name {
    t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s", name, letStmt.Name.TokenLiteral())
    return false
  }

  return true
}


func TestReturnStatements(t *testing.T) {
  input := `
  return 5;
  return 10;
  return 838383;
  `

  l := lexer.New(input)
  p := New(l)
  
  program := p.ParseProgram()
  if program == nil {
    t.Fatalf("ParseProgram() returned nil")
  }
  checkParserErrors(t, p)
  if len(program.Statements) != 3 {
    t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements)) 
  }

  tests := []struct {
    expectedIdentifier string
  }{
    {"5"},
    {"10"},
    {"838383"},
  }

  for i, _ := range tests {
    stmt := program.Statements[i]
    if !testReturnStatement(t, stmt) {
      return
    }
  }
}

func testReturnStatement(t *testing.T, s ast.Statement) bool {
  if s.TokenLiteral() != "return" {
    t.Errorf("s.TokenLiteral not 'return'. got=%q", s.TokenLiteral())
    return false
  }

  _, ok := s.(*ast.ReturnStatement)
  if !ok {
    t.Errorf("s not *ast.ReturnStatement. got %T", s)
    return false
  }

  return true
}


func TestIdentifierExpression(t *testing.T) {
  input := "foobar;"

  l := lexer.New(input)
  p := New(l)
  program := p.ParseProgram()
  checkParserErrors(t, p)

  if len(program.Statements) != 1 {
    t.Fatalf("program has not enough statements. expected=1, got=%d", len(program.Statements))
  }

  stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
  if !ok {
    t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
  }

  ident, ok := stmt.Expression.(*ast.Identifier)
  if !ok {
    t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
  }
  if ident.Value != "foobar" {
    t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
  }
  if ident.TokenLiteral() != "foobar" {
    t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar", ident.TokenLiteral()) 
  }
}
