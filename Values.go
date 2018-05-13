package main

import "fmt"

var t = "package" // -YES
// t = "package" - NO
// t := "package" - NO
// var t := "package" - NO

func main() {
  m := "two" //Inside functions only, not in the package-level variables
  //s = "one" // We can't to initialize variables with "=", without "var"
  var s = "eight"
  var x int //other initialization type
  n := "three" //We can initialize variable with :=
  n = "four" // "=" we use, when want to  assign a new value to a variable

  fmt.Println(t)  //Out: "package"
  fmt.Println(m)   //  Out: "two"
//  fmt.Println(s)  // Out: "undefined: s"
  fmt.Println(s)  //   Out: "eight"
  fmt.Println(x)  //   Out: "0"
  fmt.Println(n)  //   Out: "four"
}
