package main
import "fmt"

func main() {
    
    defer func(name string){
        fmt.Println("Bye,",name)
    } ("User")
    
    greet := func(name string){
        fmt.Println("Hello", name)
    }
    greet("User")
}
