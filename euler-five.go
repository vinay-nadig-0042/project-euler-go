package main

import (
  "fmt"
)

func main() {
  curr := 1
  flag := true
  for {
    flag = true
    for i:= 1; i < 20; i++ {
      if curr % i != 0 {
        flag = false
      }
    }
    if flag == true {
      fmt.Println(curr)
      break
    }
    curr += 1
  }
}