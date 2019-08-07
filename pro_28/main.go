package main

import (
  "fmt"
)

func main() {

  fmt.Println("hello wolrd;")
  str1 := "mississippi"
  str2 := "pi"
  res := strStr(str1, str2)
  fmt.Println(res)
}


func strStr(haystack string, needle string) int {

  if (needle == "" || len(haystack) == 1 || haystack == needle) {
   return 0
  }

  if (needle == "") {
    return 0
  }
  for i := 0; i < len(haystack) - len(needle) + 1; i ++ {
    if (haystack[i : i + len(needle)] == needle) {
      return i
    }
  }
  return -1
}
