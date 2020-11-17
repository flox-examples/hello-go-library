package example_library

import "fmt"

func HelloString(s string) string {
  return "Hello world, " + s + "!"
}

func DoWork() {
  fmt.Println("Doing some work")
}
