package main 

import (
  "testing"
)

func TestGenIntSeq(t *testing.T) {
  testseq := []int{2,3,4,5,6,7,8,9,10}
  resseq := GenIntSeq(10)
  for i := 0; i < 8; i++ {
    if testseq[i] != resseq[i] {
      t.Errorf("GenIntSeq(10) postion %d %d; want %d",i, resseq[i],testseq[i])
    }
  } 
}

func TestRemoveMultiples(t *testing.T) {
  testseq := []int{2,3,4,6,8,10}
  reqseq := RemoveMultiples(2,testseq)
  if reqseq[0] != 3 {
    t.Errorf("RemoveMultiples() %d; want %d ",3,reqseq[0])
  }

}

func TestFindPrimes(t *testing.T) {
  testseq := []int{2,3,4,5,6,7,8,9,10}
  primes := []int{2,3,5,7}
  test_primes := FindPrimes(testseq)
  for i := 0; i < 2; i++ {
    if primes[i] != test_primes[i] {
     t.Errorf("FindPrimes() postion %d %d ; want %d",i,test_primes[i],primes[i])
    }
  }

}

