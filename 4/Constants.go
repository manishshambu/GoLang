package main

import ("fmt"
        "math")

func main(){
  const a = 50

  // Does not work as we cannot change constants
  // a = 49

  fmt.Println("Value of a is ", a)

  // Allowed
  var b = math.Sqrt(4)
  fmt.Println("Value of b is ", b)

  // Not Allowed
  // To be allocated in compile time
  // const c = math.Sqrt(4)

  // Boolean constants
  // Numeric constants

  var d = 5.9/8
  fmt.Printf("Value of d is %f and type is %T", d, d)
  fmt.Println()
}
