package main

import (
  "fmt"
  "math"

)

func Sqrt(x float64) float64 {
    //hint 1: ciclo usar que itere 10 vecs
    //hint 2: usar una variable que se calclule en cada iteracion empezar en el valor medio
    z := float64(x/2)
    cnt := 0
    for  delta := z; math.Abs(delta) > 1e-6; {
        zi := z
        z -= (z*z-x)/(2*z)
        delta = z-zi
        cnt++
     }
     fmt.Println(cnt)
     return z
}

func main() {
  fmt.Println(Sqrt(10))
  fmt.Println(math.Sqrt(10))

}
