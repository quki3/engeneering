package main

import (
  "fmt"
)

func main() {
  divNumeratorByTheDenominator()
}

func divNumeratorByTheDenominator(numerator, denominator int){
  product int := numerator/denominator
  fmt.Println("%d/%d, = %d",numerator,denominator,product)
}
