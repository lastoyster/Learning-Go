package main

import (
   . "fmt"
)

func main() {
    integers := [5] int { 10, 20, 30, 40, 50 }
    nactural := [5] int { 10, 20, 30, 40, 50 }
    Println("The Length of array :",len(integers))
    
    if integers == nactural {
        Println("Both Arrays are same")
    }
    
    item := iter(integers)
    Println(item())
    Println(item())
    Println(item())
}

func iter(array [5]int) func() (int) {
    count := -1
    return func() (int) {
        count++
        return array[count]
    }
}
