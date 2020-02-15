package main

import (
  "fmt"
  "math"
  "time"
)

func main() {

    var i float64
    i = 28
    //var n float64
    //n = i+4

    //var b float64
    //b = add(i,n)
    //fmt.Println(b)

    //n=dozenIt(n,i)
    //fmt.Println(n)

    var myboolean bool
    myboolean = true;

    for i > 0 && myboolean==false{
      //fmt.Println(i,"I reckon you're right on this one, Laura! ")
      ohYes(myboolean)
      myboolean = !myboolean
      i--

    }

    //fmt.Println(n)

    //validateValidaiton(n)
    //fmt.Println(i)
    //validateValidaiton(i)

    for i > 0 && myboolean==true{
      //fmt.Println(i,"I reckon you're right on this one, Laura! ")
      //ohYes(myboolean)

      //myboolean = !myboolean


      if math.Mod(i,3)==0{
        var utctime string
        utctime = whatTimeIsItUTC()
        fmt.Println(utctime)
      }

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

func whatTimeIsItUTC()string{

  t := (time.Now().UTC())//.UnixNano())

  var outtie string

  outtie = t.String()
  return outtie
  //fmt.Println(t.Format("2006-01-02 15:04:05"))
}
