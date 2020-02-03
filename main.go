package main

import "fmt"

// this is a comment



func main() {

    var i int64
    i = 50
    var n int64
    n = i+13
    //boogie woogie

    var myboolean bool

    myboolean = false;


    for i > 0 && myboolean==false{
      //fmt.Println(i,"I reckon you're right on this one, Laura! ")
      ohYes(myboolean)
      myboolean = !myboolean
      i--

    }

    fmt.Println(n)

    validateValidaiton(n)

    for i > 0 && myboolean==true{
      //fmt.Println(i,"I reckon you're right on this one, Laura! ")
      ohYes(myboolean)

      myboolean = !myboolean

      i--
    }

}

func validateValidaiton(i float64){
  for i <10{
    fmt.Println("Yeah!!!")
  }
  fmt.Println("yeah good move")
}

func ohYes(right bool){
  fmt.Println("oh yes huntie! she did that!")
}

func add(i, n int)int {
  return i + n
}

func dozenIt(a, b int) int{
  return (a+b)*12
}
