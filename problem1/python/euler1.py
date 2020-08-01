#!/usr/local/bin/python3


def sum_multiples_3_5(max=10):
  sum = 0
  for x in range (1,max):
    if x % 3 == 0 or x % 5 == 0:
      sum+=x
  return sum


if __name__ == "__main__":
  print( sum_multiples_3_5(1000))
   
    
