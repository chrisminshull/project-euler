package main

import (
   "fmt"
)

func Multiples_3_5(max int) int {

  sum := 0
  for i := 1; i < max; i++{
    if i%3 == 0 || i%5 ==0 {
      sum+=i
    }
  } 
  return sum
}

func main() {
  fmt.Println(Multiples_3_5(1000))

}
