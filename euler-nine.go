package main

import (
  "fmt"
)

func main() {
  OuterLoop:
  for i := 0; i <= 1000; i++ {
    for j := 0; j <= 1000; j++ {
      for k := 0; k <= 1000; k++ {
        if i * i + j * j == k * k && i + j + k == 1000 && i < j && j < k {
          fmt.Println(i * j * k)
          break OuterLoop
        }
      }
    }
  }
}