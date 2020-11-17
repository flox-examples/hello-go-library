package library

import "fmt"

import "github.com/flox-examples/go-library/language"

const Language = language.Language

func HelloString(s string) string {
  return "Hello world, " + s + "!"
}

func DoWork() {
  fmt.Println("Doing some work")
}
