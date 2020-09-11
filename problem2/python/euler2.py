#!/usr/local/bin/python3

# Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
#
# 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
#
# By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

def even_fibonacci(max=100):
  sum = 0
  fib = fibonacci(max)
  for num in fib:
    if num % 2 == 0:
      sum += num
  return sum  


def fibonacci(max=100):
  fib = [1,2]
  a , b = 1 , 2 
  while a + b < max:
    temp = a + b
    fib.append(temp)
    a = b
    b = temp
  return fib
     
if __name__ == "__main__":
  print (even_fibonacci(max=4000000))    