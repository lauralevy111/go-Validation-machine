package main

import (
  "fmt"
  "time"
  "math"
)

func main() {
    output:=make(chan float64)

    i:= 10.7

    for i > 0 {
      go func() {output <- i}()

      i--
    }

    // var hello string
    // hello = "hello"
    //
    // fmt.Println(hello)

    greeting(y(), m(), d())


    var n float64
    n = 34
    validateValidaiton(math.Mod(n,32))


}

func validateValidaiton(i float64){
  for j := i; j<=10; j++{
    fmt.Println("Yeah!!!")

  }
  fmt.Println("yeah good move")
}

func ohYes(right bool){
  fmt.Println("she did that!")
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
  fmt.Println("Today is",month,",",day,",",year," and I believe in you!")
}
