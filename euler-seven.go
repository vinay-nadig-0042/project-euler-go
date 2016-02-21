package main

import (
  "fmt"
  "math/big"
)

func main() {
  counter := big.NewInt(1)
  num := big.NewInt(1)
  OuterLoop:
  for {
    if num.ProbablyPrime(64) {
      counter.Add(counter, big.NewInt(1))
    }
    if counter.String() == (big.NewInt(10002)).String() {
      fmt.Println(num)
      break OuterLoop
    }
    num.Add(num, big.NewInt(1))
  }
}