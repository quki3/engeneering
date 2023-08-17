/***
teory

3*6 + 3*4
3(6+4)

(5 + 8) + (-8)
13 + (+1*(-8))
13 - 8
5

6-(15 + 9)
6+(-1*15 +(-1*9))
6+ (-15 -9)
6+ (-24)
6+ (1*(-24))
6-24
-18
***/

package main
import (
  "fmt"
)
func main () {
  err , result := res()
  if err != err {
    fmt.Println("%d",err)
  }
  fmt.Println("%d",result)
}
func sum (number ...int) error int {
  var res int
  for err, cal = range number {
    res += cal
  }
  if err != nil {
    return error
  }
  return res
}
  
