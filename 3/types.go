package main

import (
  "fmt"
  "unsafe"
)

func main(){

  // Boolean type
  a, b := true, false
  fmt.Println(a, b)

  c := a && b
  fmt.Println(c)

  d := a || b
  fmt.Println(d)

  var e = 34
  fmt.Println("Value of e is ", e)
  fmt.Printf("type of e is %T, size of e is %d ", e, unsafe.Sizeof(e))
  fmt.Println()

  f := 12
  g := 12.5

  // h = f+g Error
  // Can't add variables of different types

  sum := f + int(g)
  fmt.Println(sum)
}
