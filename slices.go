// Take the string ‘hello world’ and slice it in two.

package main

import "fmt"

func main() {
    mystring := "hello world"
    hello := mystring[0:5]
    world := mystring[6:11]
    fmt.Println(hello)
    fmt.Println(world)
}
