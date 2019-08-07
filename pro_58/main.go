package main

import (
    "fmt"
    "strings"
)


func main() {
    // fmt.Println("a ")

    s := "b a c "

    fmt.Println(lengthOfLastWord(s))

}

func lengthOfLastWord(s string) int {


    if s == ""  || s == " "{
        return 0
    }

    if len(s) == 1 {
        return 1
    }

    ss := strings.Split(s, " ")

    fmt.Println(ss)

    fmt.Println(len(ss))

    for i := 0; i < len(ss); i ++ {
        if (ss[i] == "") {
            fmt.Println("kong !!!")
        }
        if ss[i] == " " {
            fmt.Println("ge !!!")
        }
    }

    if ss[len(ss) - 1] != "" {
         return len(ss[len(ss) - 1])
    }

    return len(ss[len(ss) - 2])

}
