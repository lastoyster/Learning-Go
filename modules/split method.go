package main 

import (
  "fmt"
  "strings"
)

var variable string
var strarray []string

func main(){
  variable = "Lorem Ipsum Dolor sit amet"

  fmt.Println(variable)

  strarray = strings.Split(variable, "")

  //do  something as often as the array contains elements
  for i := 0; i< len(strarray); i++ {

    //print the array element defined by i
    fmt.Println(strarray[i])
  }
}