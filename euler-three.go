package main

import (
  "fmt"
)

func prime_factors(n int) []int {
  z := 2
  res := make([]int, 0)
  for z * z <= n  {
    if n % z == 0 {
       res = append(res, z)
       n = n / z
    } else {
       z = z + 1
    }
  }
  if n > 1 {
    res = append(res, n)
  }
  return res
} 

func main() {
  n := 600851475143
  factors := prime_factors(n)
  fmt.Println("%#v", factors)
}
