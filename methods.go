// Create a method that sums two numbers

package main

import "fmt"

func main() {
   var a float64 = 3
   var b float64 = 9
   var ret = sum(a, b)
   fmt.Printf( "Value is : %.2f\n", ret )
}

func sum(num1, num2 float64) float64 {
   return num1+num2
}
