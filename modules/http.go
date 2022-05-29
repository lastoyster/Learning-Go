package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
)

func main(){
  response, err := http.Get("http://golang.org/")
  //send and http request
if err != nil {
  fmt.Printf("%s",err)
  //error catching
  os.Exit(1)
  
} else {
  defer response.Body.Close()
  //close the connection
  
  contents,  err := ioutil.ReadAll(response.Body)
  //read content of get request
  if err != nil {
    fmt.printf("%s", err)
    os.Exit(1)
  }
  //when caught errot, print the 
  //content of the string
fmt.Printf("%s\n",string(contents))
    }
}