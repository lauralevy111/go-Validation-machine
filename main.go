package main

import "fmt"
import "math"

// this is a comment



func main() {

    var i float64
    i = 49

    var myboolean bool

    myboolean = true;


    fmt.Println(i,"I reckon you're right on this one, Laura! ", myboolean, ", ", math.Sqrt(i))
    ohYes(myboolean)

}

func ohYes(var right bool){
  fmt.Println("I know that's ",right,"!")
}

func add(i, n int)int {
  return i + n
}
