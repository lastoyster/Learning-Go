package main

import (
    "fmt"
    "reflect"
)

/* two slices are said to be equal if they both of same type contain same elements. Golang provide a function DeepEqual from a package reflect. 

slices can not be compared using ==
An Exception...        slice == nil
                        
                        - David Ordás*/

func main() {
    slice1 := []int {1,2,3,4,5}
    slice2 := []int {1,2,3,4,5}
    slice3 := []int {9,5,5,5,2}
    fmt.Println(reflect.DeepEqual(slice1,slice2))
    fmt.Println(reflect.DeepEqual(slice1,slice3))
}