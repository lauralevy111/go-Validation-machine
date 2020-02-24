package main

import (
  "fmt"
  "math"
  "time"
)

func main() {
    output:=make(chan float64)

    var i float64
    i = 10
    // var n float64
    // n = i+4

    var myboolean bool
    myboolean = true;

    //var utctime string
    //utctime = whatTimeIsItUTC()
    //fmt.Println(utctime)

    for i > 0 && myboolean==true{

      if math.Mod(i,3)==0{
        // n=dozenIt(i,n)
        // fmt.Println(n)
      }

      i--
      go func() {output <- i}()
    }

    var hello string
    hello = "hello"//greeting(y(), m(), d())

    greeting(y(), m(), d())

    fmt.Println(hello)

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

func whatTimeIsItUTC() string{

  t := (time.Now().UTC())//.UnixNano())

  var outtie string

  outtie = t.String()
  return outtie
}

func y() int{
  t := (time.Now().UTC())
  var yearObj int

  yearObj = t.Year()

  return yearObj
}

func m() string{
  t := (time.Now().UTC())
  var monthObj time.Month
  var monthString string

  monthObj = t.Month()
  monthString = monthObj.String()

  return monthString
}

func d() int{
  t := (time.Now().UTC())
  var dayObj int

  dayObj = t.Day()

  return dayObj
}

func greeting( year int, month string, day int) {
  fmt.Println("Hello! Today is",month,",",day,",",year," and I believe in you!")
}
