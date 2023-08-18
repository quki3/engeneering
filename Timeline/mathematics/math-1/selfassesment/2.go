package main

import (
  "fmt"
)

func main() {
  a := [...]int{11,3,-4}
  rational(a)
}

func rational(a int) string {
  fmt.Println("%d/1 \n",a)
}
