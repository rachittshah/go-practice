// Make a program that divides x by 2 if it’s greater than 0

package main

import "fmt"

func main() {
    x := 3.0
    if x > 2 {
        x = x / 2
    }
    fmt.Println(x)
}
