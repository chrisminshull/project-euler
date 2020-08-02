package main

import (
   "testing"
)

func TestMultiples_3_5(t *testing.T){
  ans := Multiples_3_5(10)
  if ans != 23 {
    t.Errorf("Multiples_3_5(10) = %d; want 23", ans)
  }
}

