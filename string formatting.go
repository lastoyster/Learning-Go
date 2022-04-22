package main
import  ."fmt"
  
  type List []interface{}

func main() {
    myList := List {5, "hello","me", 5.5, 'c',
            5i, true, 'a'};
        Println(myList)
        Println("len(myList) =",len(myList))
}
  
  
