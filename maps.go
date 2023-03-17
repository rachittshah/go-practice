package main

import "fmt"

func main() {
    var n int
    fmt.Print("Enter a number: ")
    fmt.Scan(&n)

    m := make(map[int]int)
    for i := 1; i <= n; i++ {
        m[i] = i * i
    }
    fmt.Println(m)
}
