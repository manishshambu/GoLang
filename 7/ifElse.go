package main

import "fmt"

func main(){
  num := 99
  if num <= 55 {
    fmt.Println("The number is less than equal to 50")
  } else if num >=55 && num <= 100{
    fmt.Println("The number is between 55 and 100")
  } else {
    fmt.Println("The number is greater than 100")
  }
}
