package main

import "fmt"

// this is a comment



func main() {

    var i float64
    i = 5
    var n float64
    n = i+1
    //boogie woogie

    var myboolean bool

    myboolean = true;


    for i > 0 && myboolean==false{
      //fmt.Println(i,"I reckon you're right on this one, Laura! ")
      ohYes(myboolean)
      myboolean = !myboolean
      i--

    }

    fmt.Println(n)

    validateValidaiton(n)
    fmt.Println(i)
    validateValidaiton(i)

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
    i++
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
