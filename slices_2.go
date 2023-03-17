package main

import (
    "fmt"
    "strings"
)

func main() {
    var input string
    fmt.Print("Enter comma-separated numbers: ")
    fmt.Scanln(&input)

    // Split the input string into a slice of strings
    strSlice := strings.Split(input, ",")

    // Convert the string slice to an int slice
    intSlice := make([]int, len(strSlice))
    for i, str := range strSlice {
        fmt.Sscan(str, &intSlice[i])
    }

    fmt.Println(intSlice)
}
