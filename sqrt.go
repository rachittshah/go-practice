package main

import (
    "fmt"
    "math"
    "strings"
)

func main() {
    // fixed values of C and H
    const C = 50
    const H = 30
    
    // prompt user to input comma-separated values for D
    fmt.Print("Enter comma-separated values for D: ")
    var input string
    fmt.Scanln(&input)
    
    // split input string into a slice of values
    dValues := strings.Split(input, ",")
    
    // loop through each value of D and calculate Q
    for _, d := range dValues {
        // convert string value of D to a float64
        dFloat, err := strconv.ParseFloat(strings.TrimSpace(d), 64)
        if err != nil {
            fmt.Println(err)
            continue
        }
        
        // calculate Q using formula and round to nearest integer
        q := math.Sqrt((2 * C * dFloat) / H)
        qRounded := math.Round(q)
        
        // print value of Q rounded to nearest integer
        fmt.Printf("%.0f,", qRounded)
    }
}
