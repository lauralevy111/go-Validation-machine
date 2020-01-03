package main

import "fmt"

// this is a comment



func main() {

    var i float64
    i = 3

    var myboolean bool

    myboolean = true;

    for i > 0 {
      fmt.Println(i,"I reckon you're right on this one, Laura! ")
      i --
    }
    ohYes(myboolean)

}

func ohYes(right bool){
  fmt.Println("I know that's ",right,"!")
}

func add(i, n int)int {
  return i + n
}
