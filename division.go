// Write a program which will find all such numbers which are divisible by 7 but are not a multiple of 5, between 2000 and 3200 (both included). The numbers obtained should be printed in a comma-separated sequence on a single line.

package main

import "fmt"

func main() {
    // loop through numbers from 2000 to 3200 (both inclusive)
    for i := 2000; i <= 3200; i++ {
        // check if the number is divisible by 7 and not a multiple of 5
        if i%7 == 0 && i%5 != 0 {
            // print the number followed by a comma
            fmt.Printf("%d,", i)
        }
    }
}


