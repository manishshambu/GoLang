package main

import "fmt"

func main() {
  var age, year int = 100, 4
  fmt.Println("My age is", age, year)

  age = 54
  fmt.Println("My age is",age)

  age = 29
  fmt.Println("My age is", age)

  var number = 10
  fmt.Println("Number is", number)

  var (
    name = "Mainsh"
    city = "Boulder"
    graduation int =2018

  )
  fmt.Println(name, "is from", city, "graduating in ", graduation )

}
