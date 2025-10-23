package main

import (
  "fmt"
  "os"
  "os/user"
  "github.com/j-wut/monkey/repl"
)

func main() {
  user, err := user.Current()
  if err != nil {
    panic(err)
  }

  fmt.Printf("Hello %s! This is Monkey\n", user.Username)
  repl.Start(os.Stdin, os.Stdout)
}
