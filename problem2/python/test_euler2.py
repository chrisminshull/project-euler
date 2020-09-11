from euler2 import even_fibonacci
from euler2 import fibonacci

def test_even_fibonacc():
  assert even_fibonacci(100) == 44

def test_fibonacci():
  test_fib = [1,2,3,5,8,13,21,34,55,89]
  fib = fibonacci(100)
  
  assert len(test_fib) == len(fib)
