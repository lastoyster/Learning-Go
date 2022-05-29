package main

import "fmt"

func main(){
  var s string
  fmt.Print("please insert a string and press enter")
  fmt.Scan(&s)
  fmt.Printf("read string \"%v\" from stdin\n", s)
}