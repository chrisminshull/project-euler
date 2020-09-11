package main

import (
   "testing"
)

func TestFibonacci(t *testing.T){
  testfib := []int{1,2,3,5,8,13,21,34,55,89}
  fib := Fibonacci(100)
  if len(fib) != len(testfib) {
    t.Errorf("Fibonacci(100) = %d; want 10 ", len(fib))
  }
}

func TestSumEven(t *testing.T){
  test := []int{2,2,3,4}
  sum := SumEven(test)
  if sum != 8 {
    t.Errorf("SumEven({2,2,3,4}) = %d; want 8", SumEven(test))
  }
}


