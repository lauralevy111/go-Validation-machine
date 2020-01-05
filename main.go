package main

import "fmt"

// this is a comment



func main() {

    var i float64
    i = 24

    var myboolean bool

    myboolean = true;

    for i > 0 && myboolean==true{
      //fmt.Println(i,"I reckon you're right on this one, Laura! ")
      ohYes(myboolean)
      myboolean = !myboolean
      i--
    }
    for i > 0 && myboolean==false{
      //fmt.Println(i,"I reckon you're right on this one, Laura! ")
      //ohYes(myboolean)
      ohWhat()
      myboolean = !myboolean
      i--
    }
    //ohYes(myboolean)

}
func ohWhat(){
  fmt.Println("what do you know???")
}

func ohYes(right bool){
  fmt.Println("I know that's ",right,"!")
}

func add(i, n int)int {
  return i + n
}
