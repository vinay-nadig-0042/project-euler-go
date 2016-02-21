package main

import (
  "fmt"
)

func main() {
  sum_of_squares := 0
  squares_of_sum := 0

  for i := 1; i <= 100; i ++ {
    sum_of_squares += i * i
    squares_of_sum += i
  }

  squares_of_sum *= squares_of_sum

  fmt.Println(squares_of_sum - sum_of_squares)
}