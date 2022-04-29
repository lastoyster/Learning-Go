package main

import . "fmt"

type List []interface{}

func main() {
    myList := List { 5, "hi", 2.3, 'b', 
                     5i, true, 'c' };
    Println(myList)
    Println("len(myList) =", len(myList))
}
