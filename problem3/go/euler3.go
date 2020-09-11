package main

import (
  "fmt"

)

func GenIntSeq(max int) []int {
  int_seq := make([]int,0)
  for i := 2; i <= max; i++ {
    int_seq = append(int_seq, i)
  }
  return int_seq 
}

func RemoveMultiples(prime int, intseq []int) []int {
  reduce_seq := make([]int,0)
  for _, i := range intseq {
    if i % prime != 0 {
       reduce_seq = append(reduce_seq,i)
    }
  }
  return reduce_seq
}

func FindPrimes(intseq []int) []int {
  primes := make([]int, 0)
  for {
    if len(intseq) == 0{
      break
    }
    fmt.Println(intseq[0])
    primes = append(primes,intseq[0])
    intseq = RemoveMultiples(intseq[0],intseq)
  }
  return primes

}
func main() {
  fmt.Println("start")
  max := 100
  integers := GenIntSeq(max)
  primes := FindPrimes(integers)
  fmt.Println(primes)
  

}
