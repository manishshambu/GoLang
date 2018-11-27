package main

import "fmt"

func main(){
  option := 9

  switch option{

  case 1:
    fmt.Println("a")
  case 2:
    fmt.Println("b")
  case 3:
    fmt.Println("c")
  case 4:
    fmt.Println("d")
  // Not allowed
  // case 4:
  //   fmt.Println("duplicate")
  case 5:
    fmt.Println("e")
  case 6,7,8,9:
    fmt.Println("fghij")
    // Transfer control to the immediate case instead of exiting the switch
    fallthrough
  default:
    fmt.Println("Default")
  }

}
