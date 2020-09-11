package main

import (
   "fmt"
)


func Fibonacci(max int) []int {

  fib := []int{1,2}
  for {
    a := fib[len(fib)-2]
    b := fib[len(fib)-1]
    temp := a + b 
    if temp >= max {
      break
    }
    fib = append(fib, temp ) 
  }   
  return fib
}

func FibonacciChan(max int, intchan chan int) {

  a := 1
  b := 2
  intchan <- 1
  intchan <- 2 
  for {
    temp := a + b 
    if temp >= max {
      close(intchan)
      break
    }
    a = b 
    b = temp
    intchan <- temp 
  }   
  
}

func SumEvenChan(intchan <-chan int, sumchan chan int){
  sum := 0
  for {
    num, valid :=  <- intchan
    if !valid {
       break
    }
    if num % 2 == 0 {
      sum = sum + num
    }
  }
  sumchan <- sum  
}

func SumEven(nums []int) int {
  sum := 0
  for _ , n := range nums {
    if n % 2 == 0 {
      sum = sum + n
    }
  }
  return sum  
}

func main() {

  intchan := make(chan int, 10)
  sumchan := make(chan int)
  defer close(sumchan)
  //fmt.Println(SumEven(Fibonacci(4000000)))

  go FibonacciChan(4000000, intchan)
  go SumEvenChan(intchan,sumchan)
 
  fmt.Println(<-sumchan)
}
