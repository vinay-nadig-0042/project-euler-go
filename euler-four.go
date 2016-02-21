package main

import (
  "fmt"
  "strconv"
)

func reverse(s string) string {
  chars := []rune(s)
  for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
    chars[i], chars[j] = chars[j], chars[i]
  }
  return string(chars)
}

func main() {
  largest := 1
  res := 0
  straight_str := ""
  rev_str := ""
  for i := 1; i <= 999; i++ {
    for j := 1; j <= 999; j++ {
      res = i * j
      straight_str = strconv.Itoa(res)
      rev_str = reverse(straight_str)
      if straight_str == rev_str && res > largest {
        largest = res
      }
    }
  }
  fmt.Println(largest)
}
