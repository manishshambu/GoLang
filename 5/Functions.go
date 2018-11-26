package main

import "fmt"

func main(){
  fmt.Println("Functions usage in goLang")

  var a, b int = 5, 8
  var c = multiply(a, b)
  fmt.Println("Value of c is", c)

  // Multiple return Values
  var add, multiply = addMul(a, b)
  fmt.Println("Value of add is", add, "Value of multiply is", multiply)

  // Named return Values
  var addition, multiplication = namedReturn(a, b)
  fmt.Println("Value of add is", addition, "Value of multiply is", multiplication)

}


func multiply(a int, b int) int{
  var result = a * b
  return result
}

func addMul(a int, b int) (int, int){
  return (a+b), (a*b)
}

func namedReturn(a int, b int) (addition int, multiplication int){
  addition = a + b
  multiplication = a * b

  return
}
