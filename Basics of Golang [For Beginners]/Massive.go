package main

import "fmt"

func main() {

    var a [2]string
    a[0] = "Hello"
    a[1] = "Oleg"
    fmt.Println(a[0], a[1])
    fmt.Println(a)

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)

    x := [5]float64{
      98,
      93,
      77,
      82,
      83,
    }
    var total float64 = 0
      for _, value := range x {
        total += value
      }
    fmt.Println(total / float64(len(x))) // arithmetical mean (len converted to float64)
}
