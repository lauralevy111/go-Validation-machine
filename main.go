package main

import "fmt"

// this is a comment



func main() {

    var i float64
    i = 38
    //boogie woogie

    var myboolean bool

    myboolean = false;


    for i > 0 && myboolean==false{
      //fmt.Println(i,"I reckon you're right on this one, Laura! ")
      //ohYes(myboolean)
      ohWhat()
      myboolean = !myboolean
      i--
    }
    for i > 0 && myboolean==true{
      //fmt.Println(i,"I reckon you're right on this one, Laura! ")
      ohYes(myboolean)
      myboolean = !myboolean
      i--
    }

}
func ohWhat(){
  fmt.Println("this is my MOMENT")
}

func ohYes(right bool){
  fmt.Println("I know that's ",right,"!")
}

func add(i, n int)int {
  return i + n
}

func dozenIt(a, b int) int{
  return (a+b)*12
}
