/*** Newton second law of motion is F=m.a*/
package main

import (
  "fmt"
)
func main() {
  const (
    m = 15 // 15kg
    a = 1  // m/seg
  )
  f , err := NewtonLaw(m,a)
  if err !=nil {
    fmt.Print("you must write int")
  }
  fmt.Printf("%.1f",f)
}

func NewtonLaw (m , a int) (int , error) {
  if reflect.TypeOf(m) != int || feflect.TypeOf(a) != int {
    return  error
  }else {
  return m * a
  }
}
  
