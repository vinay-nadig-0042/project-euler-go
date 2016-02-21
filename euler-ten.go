package main

import (
  "fmt"
  "math/big"
)

func main() {
  var arr = make([]int64, 0)
  var i int64
  var sum int64 = 0
  for i = 0; i < 2000000; i++ {
    if (big.NewInt(i)).ProbablyPrime(64) {
      arr = append(arr, i)
    }
  }

  for _, val := range arr {
    sum += val
  }
  fmt.Println(sum)
}