package main

import "fmt"

func main(){
  for i:=0; i <= 10; i++{

    if i == 2{
      continue
    }
    fmt.Printf(" %d",i)

    if i >= 5{
      break
    }
  }
  fmt.Println()

  // While loop
  j := 0
  for j <= 10{
    fmt.Printf(" %d", j)
    j = j + 1
  }
  fmt.Println()

  // Infinite loop
  for{
    fmt.Println("Infinite")
  }

}
