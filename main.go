package main

import (
  "fmt"
  "time"
)

func main() {

    var i float64
    i = 2
    var n float64
    n = i+4

    var b float64
    b = add(i,n)
    fmt.Println(b)

    n=dozenIt(n,i)
    fmt.Println(n)

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
  for j := i; j<=10; j++{
    fmt.Println("Yeah!!!")

  }
  fmt.Println("yeah good move")
}

func ohYes(right bool){
  fmt.Println("xshe did that!")
}

func add(i, n float64)float64 {
  return i + n
}

func dozenIt(a, b float64) float64{
  return (a/b)*12
}

func whatTimeIsItUTC(name string){

  time.Now().UTC().UnixNano()
}
