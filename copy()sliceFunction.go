package main

import "fmt"
/* copy() -> function copies element from one slice to another one , The copy function follows the following signature:

func copy( destination, source [] T) int

This function takes two argumemts, the source slice that we want to copy and the destination slice to which we want to copy the elements it always return an integer type of value that is total number of copied element the return value is always the minimum of the length of source slice and destination slice. */

func main() {
	slice1 := [] int{10,20,30,40,50}
	slice2 := [] int{60,70,80,90,100,110,120,130}
	n := copy(slice2, slice1)
	fmt.Println("Elemebts copied -",n)
	fmt.Println("Slice 1- ",slice1)
	fmt.Println("Slice 2- ",slice2)
}
