package main

import (
  "fmt"
)

func fib(curr int, prev int, sum int) int {
  if curr > 4000000 {
    return sum
  } else {
    if curr % 2 == 0 {
      sum = sum + curr
    }
    return fib(curr + prev, curr, sum)
  }
  return 0
}

func main() {
  sum := fib(2, 1, 0)
  fmt.Println(sum)
}