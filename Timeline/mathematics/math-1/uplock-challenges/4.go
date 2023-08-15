/* TEORY

square root of x = [2]root{x}
periodic = ~

[2]root{25} = 25^1/2 = [2]root{ 5 * 5 ] = [2]root{5} * [2]root{5} = 5
5 = N(natural number)

33/9 = (27/9)+(6/9) = 3 + 6/9 = 3 + { x = 6/9 } = 3 + { 0,x = 60/9 } = 3 + { 0,x = (54/9)+(6/9) } = 3 + {0,6x = 6/9 } = 3+0,6~
33/9 = 3,6~ (Q rational number)(periodic)

[2]root{11} = 

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
