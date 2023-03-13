package main
import (
    "fmt"
)
func main() {
    // define array
    a := []int{1, 2, 3, 4, 5, 6}

    // loop over array
    for i := 0; i < len(a); i = i +1 {
        fmt.Println("character :", a[i])
    }
}
