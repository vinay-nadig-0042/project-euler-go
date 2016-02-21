package main

import (
  "fmt"
  "sync"
)

func main() {
  sum := 0
  var wg sync.WaitGroup
  wg.Add(999999)
  for i := 1; i < 1000000; i++ {
    go func(i int) {
      defer wg.Done()
      if i % 3 == 0 || i % 5 == 0 {
        sum += i
      }
    }(i)
  }
  wg.Wait()
  fmt.Println("Sum: ", sum)
}