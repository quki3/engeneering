/* TEORY

square root of x = [2]root{x}

[2]root{25} = 25^1/2 = [2]root{ 5 * 5 ] = [2]root{5} * [2]root{5} = 5
5 = N(natural number)
*/
package main

import (
  "fmt"
  "math"
)

func main () {
  const (
    a = 2
    b = 23
  )
  c := math.Sqrt(a+b)
  fmt.Print("%.1f",c)
}
