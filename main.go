package main

import "fmt"

// this is a comment



func main() {

    var i int
    i++
    i++
    var myboolean bool

    myboolean = true;


    fmt.Println(i,"I reckon you're right on this one, Laura!",myboolean)
    ohYes()
    fmt.Println(add(i,i))
}

func ohYes(){
  fmt.Println("I know that's right!")
}

func add(i, n int)int {
  return i + n
}
